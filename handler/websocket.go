package handler

import (
	"../model"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
	"net/http"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(ctx context.Context,hub *model.NotificationHub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)

			return
		}
		client := model.NewNotificationClient(hub,  conn,  make(chan []byte, 256))
		client.RegisterClient()

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.WritePump()
		go client.ReadPump()
	})
}
