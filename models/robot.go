package models

import (
	"fmt"
)

// Robot represents an AutoStore robot
type Robot struct {
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Z      int    `json:"z"`
	Status string `json:"status"`
}

// DisplayInfo prints robot information to console
func (r Robot) DisplayInfo() {
	fmt.Printf("Robot %d at position (%d, %d) - Status: %s\n", r.ID, r.X, r.Y, r.Status)
}

// Old MoveTo method for compatibility temporarily
func (r *Robot) MoveTo(newX, newY int) {
	r.X = newX
	r.Y = newY
	r.Status = "moving"
	fmt.Printf("Robot %d moved to (%d, %d)\n", r.ID, r.X, r.Y)
}

// New MoveTo method with warehouse bounds checking
func (r *Robot) MoveToWithBounds(newX, newY, newZ int, warehouse Warehouse) bool {
	if !warehouse.IsValidPosition(newX, newY, newZ) {
		fmt.Printf("Robot %d: Invalid position (%d, %d, %d)\n", r.ID, newX, newY, newZ)
		return false
	}

	r.X = newX
	r.Y = newY
	r.Z = newZ
	r.Status = "moving"
	fmt.Printf("Robot %d moved to (%d, %d, %d)\n", r.ID, r.X, r.Y, r.Z)
	return true
}
