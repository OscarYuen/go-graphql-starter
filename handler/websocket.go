package handler

import (
	"../config"
	"../model"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(hub *model.NotificationHub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthorized := r.Context().Value("is_authorized").(bool); !isAuthorized {
			http.Error(w, config.CredentialsError, http.StatusUnauthorized)
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		userId := r.Context().Value("user_id").(*int64)
		client := model.NewNotificationClient(userId, hub, conn, make(chan []byte, 256))
		client.RegisterClient()

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.WritePump()
		go client.ReadPump()
	})
}
