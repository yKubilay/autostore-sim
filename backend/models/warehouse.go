package models

import (
	"sync"
)

// Warehouse represents the physical warehouse structure
type Warehouse struct {
	Width  int `json:"width"`  // X dimension
	Height int `json:"height"` // Y dimension
	Levels int `json:"levels"` // Z dimension (for 3D)
}

// Safewarehouse creates thread-safe warehouse with collision detection
type SafeWarehouse struct {
	Width  int          `json:"width"`
	Height int          `json:"height"`
	Levels int          `json:"levels"`
	Grid   [][][]bool   `json:"-"`
	Mutex  sync.RWMutex `json:"-"`
}

// NewSafeWarehouse creates new thread-safe with the initialized grid
func NewSafeWarehouse(width, height, levels int) *SafeWarehouse {
	// Initilizating 3D grid
	grid := make([][][]bool, width)
	for x := range grid {
		grid[x] = make([][]bool, height)
		for y := range grid[x] {
			grid[x][y] = make([]bool, levels)
		}
	}

	return &SafeWarehouse{
		Width:  width,
		Height: height,
		Levels: levels,
		Grid:   grid,
	}
}

// GetDefaultSafeWarehouse returns a standard 5x5x3 thread-safe warehouse
func GetDefaultSafeWarehouse() *SafeWarehouse {
	return NewSafeWarehouse(5, 5, 3)
}

// IsValidPosition checks if coordinates are within warehouse bounds
func (w Warehouse) IsValidPosition(x, y, z int) bool {
	return x >= 0 && x < w.Width &&
		y >= 0 && y < w.Height &&
		z >= 0 && z < w.Levels
}

// GetDefaultWarehouse returns a standard 5x5x3 AutoStore warehouse
func GetDefaultWarehouse() Warehouse {
	return Warehouse{
		Width:  5,
		Height: 5,
		Levels: 3,
	}
}

// IsPositionOccupied checks if a position is occupied using read lock
func (sw *SafeWarehouse) IsPositionOccupied(x, y, z int) bool {
	if !sw.IsValidPosition(x, y, z) {
		return true // Out of bounds = "occupied"
	}

	sw.Mutex.RLock() // Read lock, multiple robots can check simultaneously
	occupied := sw.Grid[x][y][z]
	sw.Mutex.RUnlock()

	return occupied
}

// SetRobotPosition safely sets and unsets a robot position using write lock
func (sw *SafeWarehouse) SetRobotPosition(x, y, z int, occupied bool) bool {
	if !sw.IsValidPosition(x, y, z) {
		return false
	}

	sw.Mutex.Lock() // Write lock, only one robot can modify at a time
	sw.Grid[x][y][z] = occupied
	sw.Mutex.Unlock()

	return true
}

// TryMoveRoboty safely moves a robot from one position to another
func (sw *SafeWarehouse) TryMoveRobot(fromX, fromY, fromZ, toX, toY, toZ int) bool {
	if !sw.IsValidPosition(toX, toY, toZ) {
		return false // Cannot move to invalid position
	}

	sw.Mutex.Lock()         // Write lock for entire operation
	defer sw.Mutex.Unlock() // Automatically unlock when function exits

	// Check if destintation is already occupied
	if sw.Grid[toX][toY][toZ] {
		return false // Destination occupied, cannot move
	}

	// Move robot, free old position and occupy new position
	if sw.IsValidPosition(fromX, fromY, fromZ) {
		sw.Grid[fromX][fromY][fromZ] = false // free old position
	}
	sw.Grid[toX][toY][toZ] = true // Occupy new position

	return true // move successful
}

// IsValidPosition checks if coordinates are within SafeWarehouse bounds
func (sw *SafeWarehouse) IsValidPosition(x, y, z int) bool {
	return x >= 0 && x < sw.Width &&
		y >= 0 && y < sw.Height &&
		z >= 0 && z < sw.Levels
}
