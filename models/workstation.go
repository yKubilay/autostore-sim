package models

// Workstation represents a port where robots deliver bins
type Workstation struct {
	ID     int    `json:"id"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Status string `json:"status"`
}
