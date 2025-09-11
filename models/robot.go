package models

import "fmt"

// Robot represents an AutoStore robot
type Robot struct {
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Status string `json:"status"`
}

// DisplayInfo prints robot information to console
func (r Robot) DisplayInfo() {
	fmt.Printf("Robot %d at position (%d, %d) - Status: %s\n", r.ID, r.X, r.Y, r.Status)
}

// MoveTo updates robot position and status
func (r *Robot) MoveTo(newX, newY int) {
	r.X = newX
	r.Y = newY
	r.Status = "moving"
	fmt.Printf("Robot %d moved to (%d, %d)\n", r.ID, r.X, r.Y)
}
