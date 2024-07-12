package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

//done := make(chan struct{})：创建一个空结构体类型的通道。空结构体类型（struct{}）在 Go 中是零大小的，可以用来进行信号传递而不需要实际传递数据。

//Goroutine：在新的 Goroutine 中，开始一个无限循环从 WebSocket 连接读取消息。如果读取操作出错（例如，连接关闭），defer close(done) 会关闭 done 通道，通知主 Goroutine。

//select 语句中的 case <-done:：当 done 通道被关闭时，这个 case 会被触发，主 Goroutine 会退出循环，从而终止程序。

// interrupt 信号处理：当接收到中断信号（如 Ctrl+C）时，主 Goroutine 会尝试优雅地关闭 WebSocket 连接，然后等待 done 通道的信号，确保读取 Goroutine 能够正常退出。

func main() {
	interrupt := make(chan os.Signal, 1)
	// 构建url
	wsUrl := url.URL{Scheme: "ws", Host: "localhost:10500", Path: "/ws"}
	log.Printf("connecting to %s", wsUrl.String())

	// 连接websocket地址
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), nil)
	if err != nil {
		log.Println("dail:", err)
	}

	defer conn.Close()

	done := make(chan struct{})

	// 用于产生周期性事件。创建一个 Ticker 会返回一个通道，这个通道会在指定的时间间隔上发送当前时间的值。
	// time.NewTicker(time.Second) 创建了一个新的 Ticker，它会每秒钟触发一次。
	ticker := time.NewTicker(time.Second)

	go func() {
		defer close(done)
		for {
			var msg map[string]interface{}
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("Received: %v", msg)
		}
	}()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			// 执行此分支时，会从 ticker.C 通道中接收当前时间值
			msg := map[string]interface{}{
				"time":    t.String(),
				"message": "Hello WebSocket",
			}
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println("write error")
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
