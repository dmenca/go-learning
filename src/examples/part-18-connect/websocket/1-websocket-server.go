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

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if err := ws.Close(); err != nil {
			log.Panicln("close ws error", err)
		}
	}()
	// 每个连接独立处理
	for {
		// 读取消息
		// map[string]interface{} 用于存储 JSON 对象，其中键是字符串，值可以是任何类型。
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("read json error", err)
			break
		}
		err = ws.WriteJSON(msg)
		if err != nil {
			log.Printf("write json error", err)
			break
		}
	}

}

func main() {
	// 注册/ws路由
	interrupt := make(chan os.Signal, 1)
	// 使用 signal.Notify 来监听操作系统信号 这样可以将操作系统的中断信号（如 Ctrl+C）通知到 interrupt 通道。
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/ws", handleConnections)
	go func() {
		err := http.ListenAndServe(":10500", nil)
		if err != nil {
			log.Fatal("Error starting HTTP server:", err)
		} else {
			log.Println("Http server listening on :8080")
		}
	}()

	log.Println("server go on")

	select {
	case <-interrupt:
		log.Println("server receive interrupt siginal")
	}
}
