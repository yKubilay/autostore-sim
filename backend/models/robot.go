package models

import (
	"fmt"
	"time"
)

// Robot represents an AutoStore robot
type Robot struct {
	ID              int               `json:"id"`
	X               int               `json:"x"`
	Y               int               `json:"y"`
	Z               int               `json:"z"`
	Status          string            `json:"status"`
	Commands        chan RobotCommand `json:"-"`
	Updates         chan RobotUpdate  `json:"-"`
	BroadcastUpdate func(RobotUpdate) `json:"-"` // Callback for broadcasting updates
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
func (r *Robot) StartRobot(sw *SafeWarehouse, done chan bool) {
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

				r.processCommand(cmd, sw)

			case <-done:
				fmt.Printf("Robot %d shutting down\n", r.ID)
				return // Exiting the goroutine
			}
		}
	}()
}

// processCommand handles actual command execution for robots
func (r *Robot) processCommand(cmd RobotCommand, sw *SafeWarehouse) {
	switch cmd.Type {
	case "move":
		// Use SafeWarehouse's thread-safe movement
		if sw.CanRobotMoveTo(cmd.X, cmd.Y, cmd.Z) {

			travelTime := r.calculateTravelTime(cmd.X, cmd.Y, cmd.Z)

			r.Status = "moving"
			fmt.Printf("Robot %d moving to (%d, %d, %d), estimated time: %v\n",
				r.ID, cmd.X, cmd.Y, cmd.Z, travelTime)

			// Simulate travel time
			time.Sleep(travelTime)

			// Update position after travel
			r.X = cmd.X
			r.Y = cmd.Y
			r.Z = cmd.Z
			r.Status = "idle"

			fmt.Printf("Robot %d arrived at (%d, %d, %d)\n", r.ID, r.X, r.Y, r.Z)

			// Broadcast update via WebSocket
			if r.BroadcastUpdate != nil {
				r.BroadcastUpdate(RobotUpdate{
					RobotID: r.ID,
					X:       r.X,
					Y:       r.Y,
					Z:       r.Z,
					Status:  r.Status,
				})
			}
		} else {
			r.Status = "error"
			fmt.Printf("Robot %d: Move failed\n", r.ID)
			fmt.Printf("Robot %d cannot move to (%d, %d, %d) - position occupied or out of bounds\n",
				r.ID, cmd.X, cmd.Y, cmd.Z)
			fmt.Printf("Robot %d: Move failed - invalid position (%d, %d, %d)\n", r.ID, cmd.X, cmd.Y, cmd.Z)
		}
	case "pick":
		// First move to pick location if not already there
		if r.X != cmd.X || r.Y != cmd.Y || r.Z != cmd.Z {
			travelTime := r.calculateTravelTime(cmd.X, cmd.Y, cmd.Z)
			r.Status = "moving"
			fmt.Printf("Robot %d moving to pick location (%d, %d, %d) - ETA: %.1fs\n",
				r.ID, cmd.X, cmd.Y, cmd.Z, travelTime.Seconds())
			time.Sleep(travelTime)
			r.X = cmd.X
			r.Y = cmd.Y
			r.Z = cmd.Z
		}

		r.Status = "picking"
		fmt.Printf("Robot %d picking up item at (%d, %d, %d)\n", r.ID, cmd.X, cmd.Y, cmd.Z)
		// Realistic pick time (lowering bin, grabbing, lifting)
		time.Sleep(2 * time.Second)
		r.Status = "carrying"
		fmt.Printf("Robot %d picked up item for order %d\n", r.ID, cmd.OrderID)

		// Broadcast update via WebSocket
		if r.BroadcastUpdate != nil {
			r.BroadcastUpdate(RobotUpdate{
				RobotID: r.ID,
				X:       r.X,
				Y:       r.Y,
				Z:       r.Z,
				Status:  r.Status,
				OrderID: cmd.OrderID,
			})
		}
	case "drop":
		r.Status = "dropping"
		fmt.Printf("Robot %d dropping item at (%d, %d, %d)\n", r.ID, cmd.X, cmd.Y, cmd.Z)
		// Realistic drop time (lowering, placing, lifting)
		time.Sleep(1500 * time.Millisecond)
		r.Status = "idle"
		fmt.Printf("Robot %d completed delivery for order %d\n", r.ID, cmd.OrderID)

		// Broadcast update via WebSocket
		if r.BroadcastUpdate != nil {
			r.BroadcastUpdate(RobotUpdate{
				RobotID: r.ID,
				X:       r.X,
				Y:       r.Y,
				Z:       r.Z,
				Status:  r.Status,
				OrderID: cmd.OrderID,
			})
		}
	}
}

// calculateTravelTime calculates realistic travel time based on distance
func (r *Robot) calculateTravelTime(targetX, targetY, targetZ int) time.Duration {

	// Real AutoStore physical constants
	const (
		GRID_WIDTH_METERS      = 0.705 // 705mm wide direction
		GRID_DEPTH_METERS      = 0.480 // 480mm narrow direction
		BIN_HEIGHT_METERS      = 0.330 // 330mm bins
		ROBOT_HORIZONTAL_SPEED = 3.1   // m/s (real spec)
		ROBOT_LIFT_SPEED       = 1.6   // m/s (real spec)
		ROBOT_ACCELERATION     = 0.8   // m/s²
	)

	// Calculate Manhattan distance
	deltaX := abs(r.X - targetX)
	deltaY := abs(r.Y - targetY)
	deltaZ := abs(r.Z - targetZ)

	// Calculate actual distances in meters
	horizontalDistance := float64(deltaX)*GRID_WIDTH_METERS + float64(deltaY)*GRID_DEPTH_METERS
	verticalDistance := float64(deltaZ) * BIN_HEIGHT_METERS

	// Calculate travel times based on real AutoStore speeds
	horizontalTime := horizontalDistance / ROBOT_HORIZONTAL_SPEED
	verticalTime := verticalDistance / ROBOT_LIFT_SPEED
	totalTime := horizontalTime + verticalTime

	// Add small base time for acceleration/deceleration
	if totalTime > 0 {
		accelTime := ROBOT_HORIZONTAL_SPEED / ROBOT_ACCELERATION // Time to reach max speed
		totalTime += accelTime * 0.5                             // Account for accel/decel
	}

	return time.Duration(totalTime * float64(time.Second))
}

// abs returns absolute value of integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
