package models

import "fmt"

// Order represents a warehouse order
type Order struct {
	ID            int    `json:"id"`
	ItemX         int    `json:"item_x"`
	ItemY         int    `json:"item_y"`
	DeliveryX     int    `json:"delivery_x"`
	DeliveryY     int    `json:"delivery_y"`
	Status        string `json:"status"`
	AssignedRobot int    `json:"assigned_robot"`
}

// DisplayInfo prints order information to console
func (o Order) DisplayInfo() {
	fmt.Printf("Order %d: Pick from (%d, %d) deliver to (%d, %d) - Status: %s\n",
		o.ID, o.ItemX, o.ItemY, o.DeliveryX, o.DeliveryY, o.Status)
}
