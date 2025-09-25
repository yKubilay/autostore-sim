package models

import (
	"fmt"
	"time"
)

// Robot represents an AutoStore robot
type Robot struct {
	ID       int               `json:"id"`
	X        int               `json:"x"`
	Y        int               `json:"y"`
	Z        int               `json:"z"`
	Status   string            `json:"status"`
	Commands chan RobotCommand `json:"-"`
	Updates  chan RobotUpdate  `json:"-"`
}

// RobotCommand represents a command sent to robot
type RobotCommand struct {
	Type    string `json:"type"` // "move", "pick", "drop"
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Z       int    `json:"z"`
	OrderID int    `json:"order_id"`
}

// RobotUpdate represents status updates from robots
type RobotUpdate struct {
	RobotID int    `json:"robot_id"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Z       int    `json:"z"`
	Status  string `json:"status"`
	OrderID int    `json:"order_id,omitempty"`
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

// StartRobot to launch the robot as gouroutine with channels for communication
func (r *Robot) StartRobot(warehouse *Warehouse, done chan bool) {
	// Initialize channels
	r.Commands = make(chan RobotCommand, 10)
	r.Updates = make(chan RobotUpdate, 10)

	fmt.Printf("Robot %d started as goroutine at position (%d, %d, %d)\n", r.ID, r.X, r.Y, r.Z)

	// Launch the worker goroutine
	go func() {
		for {
			select {
			// Listen for commands
			case cmd := <-r.Commands:
				fmt.Printf("Robot %d received command: %s to (%d, %d, %d)\n",
					r.ID, cmd.Type, cmd.X, cmd.Y, cmd.Z)

				r.processCommand(cmd, warehouse)

			case <-done:
				fmt.Printf("Robot %d shutting down\n", r.ID)
				return // Exiting the goroutine
			}
		}
	}()
}

// processCommand handles actual command execution for robots
func (r *Robot) processCommand(cmd RobotCommand, warehouse *Warehouse) {
	switch cmd.Type {
	case "move":

		// Use existing bounds checking method for robots
		if r.MoveToWithBounds(cmd.X, cmd.Y, cmd.Z, *warehouse) {
			r.Status = "idle"

			// Sending updates back
			r.Updates <- RobotUpdate{
				RobotID: r.ID,
				X:       r.X,
				Y:       r.Y,
				Z:       r.Z,
				Status:  r.Status,
				OrderID: cmd.OrderID,
			}
		} else {
			r.Status = "error"
			fmt.Printf("Robot %d: Move failed\n", r.ID)
		}
	case "pick":
		r.Status = "picking"
		fmt.Printf("Robot %d picking up item at (%d, %d, %d)\n", r.ID, cmd.X, cmd.Y, cmd.Z)
		// Simulating pick time
		time.Sleep(500 * time.Millisecond)
		r.Status = "carrying"
	case "drop":
		r.Status = "dropping"
		fmt.Printf("Robot %d dropping item at (%d, %d, %d)\n", r.ID, cmd.X, cmd.Y, cmd.Z)
		time.Sleep(300 * time.Millisecond)
		r.Status = "idle"
	}
}
