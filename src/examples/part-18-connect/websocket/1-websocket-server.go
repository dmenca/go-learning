package main

import (
	"log"
	"net/http"

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
	for {
		// 读取消息
		// map[string]interface{} 用于存储 JSON 对象，其中键是字符串，值可以是任何类型。
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Fatal("read json error", err)
		}
		err = ws.WriteJSON(msg)
		if err != nil {
			log.Fatal("write json error", err)
		}
	}

}

func main() {
	// 注册/ws路由
	http.HandleFunc("/ws", handleConnections)

	err := http.ListenAndServe(":10500", nil)
	if err != nil {
		log.Println("Error starting HTTP server:", err)
	} else {
		log.Println("Http server listening on :8080")
	}

}
