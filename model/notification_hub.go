package model

import (
	"encoding/json"
)

type NotificationHub struct {
	// Registered clients.
	clients map[*NotificationClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *NotificationClient

	// Unregister requests from clients.
	unregister chan *NotificationClient
}

func NewNotificationHub() *NotificationHub {
	return &NotificationHub{
		broadcast:  make(chan []byte),
		register:   make(chan *NotificationClient),
		unregister: make(chan *NotificationClient),
		clients:    make(map[*NotificationClient]bool),
	}
}

func (h *NotificationHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case notificationStr := <-h.broadcast:
			notification := &Notification{}
			if err := json.Unmarshal(notificationStr, notification); err == nil {
				for client := range h.clients {
					if *client.userId == notification.To {
						select {
						case client.send <- notificationStr:
						default:
							close(client.send)
							delete(h.clients, client)
						}
					}
				}
			}
		}
	}
}

func (h *NotificationHub) BroadcastMessage(notification *Notification) {
	notificationStr, _ := json.Marshal(notification)
	h.broadcast <- notificationStr
}
