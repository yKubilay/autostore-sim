package models

import "time"

// Order represents a customer order for auto parts
type Order struct {
	ID            int         `json:"id"`
	CustomerName  string      `json:"customer_name"`          // "Smith Auto Repair"
	ProductID     int         `json:"product_id"`             // ID of requested product
	RequestedQty  int         `json:"requested_qty"`          // How many items needed
	Status        OrderStatus `json:"status"`                 // pending, assigned, picking, etc.
	Priority      Priority    `json:"priority"`               // normal, urgent, express
	AssignedRobot int         `json:"assigned_robot"`         // Which robot is handling this
	CreatedAt     time.Time   `json:"created_at"`             // When order was placed
	CompletedAt   *time.Time  `json:"completed_at,omitempty"` // When order finished
	DeliveryPort  Position    `json:"delivery_port"`          // Which port to deliver to
}

// OrderStatus represents the current state of an order
type OrderStatus string

const (
	OrderPending    OrderStatus = "pending"    // Waiting for robot assignment
	OrderAssigned   OrderStatus = "assigned"   // Robot assigned but not started
	OrderPicking    OrderStatus = "picking"    // Robot is picking the product
	OrderDelivering OrderStatus = "delivering" // Robot moving to delivery port
	OrderCompleted  OrderStatus = "completed"  // Order fulfilled
	OrderFailed     OrderStatus = "failed"     // Could not fulfill (no stock, etc.)
)

// Priority represents order urgency
type Priority string

const (
	PriorityNormal  Priority = "normal"  // Regular delivery
	PriorityUrgent  Priority = "urgent"  // Rush job - higher priority
	PriorityExpress Priority = "express" // Emergency - highest priority
)

// OrderQueue manages pending orders
type OrderQueue struct {
	Orders []Order `json:"orders"`
	NextID int     `json:"next_id"`
}

// NewOrderQueue creates a new order queue
func NewOrderQueue() *OrderQueue {
	return &OrderQueue{
		Orders: make([]Order, 0),
		NextID: 1,
	}
}

// AddOrder adds a new order to the queue
func (oq *OrderQueue) AddOrder(customerName string, productID int, requestedQty int, priority Priority) *Order {
	order := Order{
		ID:            oq.NextID,
		CustomerName:  customerName,
		ProductID:     productID,
		RequestedQty:  requestedQty,
		Status:        OrderPending,
		Priority:      priority,
		AssignedRobot: 0,
		CreatedAt:     time.Now(),
		DeliveryPort:  Position{X: -1, Y: -1, Z: -1}, // Will be assigned later
	}

	oq.Orders = append(oq.Orders, order)
	oq.NextID++

	return &order
}

// GetPendingOrders returns all orders waiting for robot assignment
func (oq *OrderQueue) GetPendingOrders() []Order {
	var pending []Order
	for _, order := range oq.Orders {
		if order.Status == OrderPending {
			pending = append(pending, order)
		}
	}
	return pending
}

// GetOrderByID finds an order by its ID
func (oq *OrderQueue) GetOrderByID(id int) *Order {
	for i := range oq.Orders {
		if oq.Orders[i].ID == id {
			return &oq.Orders[i]
		}
	}
	return nil
}
