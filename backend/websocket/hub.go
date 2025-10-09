package websocket

import (
	"autostore-sim/backend/models"
	"encoding/json"
	"log"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to them
type Hub struct {
	// Registered clients
	clients map[*Client]bool

	// Inbound messages from clients
	broadcast chan []byte

	// Register requests from clients
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client

	// Mutex for thread-safe access
	mu sync.RWMutex
}

// NewHub creates a new Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the hub's main loop
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("Client registered, total clients: %d", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mu.Unlock()
			log.Printf("Client unregistered, total clients: %d", len(h.clients))

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// BroadcastWarehouseState sends current warehouse state to all connected clients
func (h *Hub) BroadcastWarehouseState(robots []*models.Robot, orders []models.Order) {
	state := map[string]interface{}{
		"type":   "warehouse_update",
		"robots": robots,
		"orders": orders,
	}

	data, err := json.Marshal(state)
	if err != nil {
		log.Printf("Error marshaling warehouse state: %v", err)
		return
	}

	h.broadcast <- data
}

// BroadcastRobotUpdate sends a single robot update to all connected clients
func (h *Hub) BroadcastRobotUpdate(update models.RobotUpdate) {
	message := map[string]interface{}{
		"type":   "robot_update",
		"update": update,
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling robot update: %v", err)
		return
	}

	h.broadcast <- data
}
