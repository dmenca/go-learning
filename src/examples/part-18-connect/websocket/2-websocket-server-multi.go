package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var server = &Server{
	clients:    make(map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	broadcast:  make(chan Message),
}

// 定义客户端 客户端包含订阅的频道以及真实连接
type Client struct {
	conn     *websocket.Conn
	channels map[string]bool
}

// 定义服务端
type Server struct {
	// 包含连接的客户端
	clients map[*Client]bool
	// 订阅客户端信号量
	register chan *Client
	// 取消订阅客户端信号量
	unregister chan *Client
	// 发送消息到订阅了指定频道的客户端信号量
	broadcast chan Message
}

type Message struct {
	// 频道消息
	Channel string `json:"channel"`
	// 消息内容
	Data map[string]interface{} `json:"data"`
}

// 运行方法
func (s *Server) run() {
	for {
		select {
		case client := <-s.register:
			log.Println("registry client")
			s.clients[client] = true
		case client := <-s.unregister:
			log.Println("unregistry client")
			delete(s.clients, client)
			closeClientConnection(client)
		case message := <-s.broadcast:
			log.Println("broadcast message:", message)
			channel := message.Channel
			for client, _ := range s.clients {
				b, exists := client.channels[channel]
				if exists && b {
					if err := client.conn.WriteJSON(message); err != nil {
						s.unregister <- client
						log.Println("write message error:", err)
					}
				}
			}
		}
	}
}

func closeClientConnection(client *Client) {
	client.conn.Close()
}

func handleMessages(client *Client) {
	// 定义注销订阅
	defer func() {
		server.unregister <- client
		closeClientConnection(client)
	}()
	// 定义接收消息和订阅
	for {
		var message Message
		err := client.conn.ReadJSON(&message)
		if err != nil {
			log.Println("read message error.", err)
			break
		}
		log.Println("success received:", message)
		switch message.Channel {
		case "subscribe":
			client.channels[message.Data["channel"].(string)] = true
		case "unsubscribe":
			delete(client.channels, message.Data["channel"].(string))
		default:
			log.Println("handleMessages :", message)
			server.broadcast <- message
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: ws, channels: make(map[string]bool)}
	log.Println("handle connection:", client)
	// 注册客户端
	server.register <- client
	go handleMessages(client)

}

func main() {
	// 注册/ws路由
	interrupt := make(chan os.Signal, 1)
	// 使用 signal.Notify 来监听操作系统信号 这样可以将操作系统的中断信号（如 Ctrl+C）通知到 interrupt 通道。
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/ws", handleConnections)
	go func() {
		err := http.ListenAndServe(":10501", nil)
		if err != nil {
			log.Fatal("Error starting HTTP server:", err)
		} else {
			log.Println("Http server listening on :8080")
		}
	}()

	// 开启server监听
	go server.run()

	log.Println("server go on")

	select {
	case <-interrupt:
		log.Println("server receive interrupt siginal")
	}
}
