package handlers

import (
	"autostore-sim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WahouseData holds all warehouse state data
type WarehouseData struct {
	Robots       []models.Robot       `json:"robots"`
	Orders       []models.Order       `json:"orders"`
	Workstations []models.Workstation `json:"workstations"`
}

// Global warehouse data (will be in a db later on)
var warehouse WarehouseData

// SetWarehouseData initliazises the warehouse data
func SetWarehouseData(robots []models.Robot, orders []models.Order, workstations []models.Workstation) {
	warehouse.Robots = robots
	warehouse.Orders = orders
	warehouse.Workstations = workstations
}

// GetRobots returns all robots
func GetRobots(c *gin.Context) {
	c.JSON(http.StatusOK, warehouse.Robots)
}

// GetOrders returns all orders
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, warehouse.Orders)
}

// GetWorkstations returns all workstations
func GetWorkstations(c *gin.Context) {
	c.JSON(http.StatusOK, warehouse.Workstations)
}

// GetWarehouseStatus returns complete warehouse state
func GetWarehouseStatus(c *gin.Context) {
	c.JSON(http.StatusOK, warehouse)
}
