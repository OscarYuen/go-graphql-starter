package model

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
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *NotificationHub) BroadcastMessage(message []byte) {
	h.broadcast <- message
}

