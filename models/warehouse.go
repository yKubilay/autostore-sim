package models

// Warehouse represents the physical warehouse structure
type Warehouse struct {
	Width  int `json:"width"`  // X dimension
	Height int `json:"height"` // Y dimension
	Levels int `json:"levels"` // Z dimension (for 3D)
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
