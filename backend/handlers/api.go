package handlers

import (
	"autostore-sim/backend/services"
	"autostore-sim/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WarehouseData holds all warehouse state data
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

// GetWarehouseData returns current warehouse state for processing
func GetWarehouseData() ([]models.Robot, []models.Order, []models.Workstation) {
	return warehouse.Robots, warehouse.Orders, warehouse.Workstations
}

func ProcessWarehouseOrders() {
	services.ProcessPendingOrders(warehouse.Robots, warehouse.Orders)
}

// CreateOrderRequest represents the JSON structure for creating orders
type CreateOrderRequest struct {
	ItemX     int `json:"item_x" binding:"required"`
	ItemY     int `json:"item_y" binding:"required"`
	DeliveryX int `json:"delivery_x" binding:"required"`
	DeliveryY int `json:"delivery_y" binding:"required"`
}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate new order ID
	newID := len(warehouse.Orders) + 1

	// Create new order
	newOrder := models.Order{
		ID:            newID,
		ItemX:         req.ItemX,
		ItemY:         req.ItemY,
		DeliveryX:     req.DeliveryX,
		DeliveryY:     req.DeliveryY,
		Status:        "pending",
		AssignedRobot: 0,
	}

	// Add to warehouse
	warehouse.Orders = append(warehouse.Orders, newOrder)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   newOrder,
	})
}
