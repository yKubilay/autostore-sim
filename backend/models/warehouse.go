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
	Width  int               `json:"width"`
	Height int               `json:"height"`
	Levels int               `json:"levels"`
	Grid   [][][]StorageCell `json:"-"`
	Mutex  sync.RWMutex      `json:"-"`
}

// NewSafeWarehouse creates new thread-safe with the initialized grid
func NewSafeWarehouse(width, height, levels int) *SafeWarehouse {
	// Initilizating 3D grid
	grid := make([][][]StorageCell, width)
	for x := range grid {
		grid[x] = make([][]StorageCell, height)
		for y := range grid[x] {
			grid[x][y] = make([]StorageCell, levels)
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
	return NewSafeWarehouse(8, 8, 5)
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

// For robot collision detection - check if another robot is at position
func (sw *SafeWarehouse) HasRobotAt(x, y, z int) bool {
	// We'll track robot positions separately - for now, return false
	return false
}

// For checking if cell has inventory (for picking operations)
func (sw *SafeWarehouse) HasInventory(x, y, z int) bool {
	if !sw.IsValidPosition(x, y, z) {
		return false
	}

	sw.Mutex.RLock()
	hasStock := !sw.Grid[x][y][z].IsEmpty()
	sw.Mutex.RUnlock()

	return hasStock
}

// Robot movement - just check bounds (simplified for now)
func (sw *SafeWarehouse) CanRobotMoveTo(x, y, z int) bool {
	return sw.IsValidPosition(x, y, z)
}

// IsValidPosition checks if coordinates are within SafeWarehouse bounds
func (sw *SafeWarehouse) IsValidPosition(x, y, z int) bool {
	return x >= 0 && x < sw.Width &&
		y >= 0 && y < sw.Height &&
		z >= 0 && z < sw.Levels
}
