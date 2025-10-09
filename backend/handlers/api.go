package handlers

import (
	"autostore-sim/backend/models"
	"autostore-sim/backend/services"
	ws "autostore-sim/backend/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	OrderService   *services.OrderService
	ProductService *services.ProductService
	Warehouse      *models.SafeWarehouse
	Robots         []*models.Robot
	Workstations   []models.Workstation
	WebSocketHub   *ws.Hub
}

var server Server

// InitializeServer sets up all services for API handlers
func InitializeServer(os *services.OrderService, ps *services.ProductService,
	wh *models.SafeWarehouse, rbs []*models.Robot, wss []models.Workstation, hub *ws.Hub) {
	server = Server{
		OrderService:   os,
		ProductService: ps,
		Warehouse:      wh,
		Robots:         rbs,
		Workstations:   wss,
		WebSocketHub:   hub,
	}
}

// GetRobots returns all robots
func GetRobots(c *gin.Context) {
	c.JSON(http.StatusOK, server.Robots)
}

// GetOrders returns all orders
func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, server.OrderService.GetActiveOrders())
}

// GetWorkstations returns all workstations
func GetWorkstations(c *gin.Context) {
	c.JSON(http.StatusOK, server.Workstations)
}

// GetWarehouseStatus returns complete warehouse state
func GetWarehouseStatus(c *gin.Context) {
	c.JSON(http.StatusOK, server)
}

// GetWarehouseData returns current warehouse state for processing
func GetWarehouseData() ([]*models.Robot, []models.Order, []models.Workstation) {
	return server.Robots, server.OrderService.GetActiveOrders(), server.Workstations
}

func ProcessWarehouseOrders() {
	server.OrderService.ProcessPendingOrders(server.Robots)
}

// CreateOrderRequest represents the JSON structure for creating orders
type CreateOrderRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
	ProductID    int    `json:"product_id" binding:"required"`
	RequestedQty int    `json:"requested_qty" binding:"required,min=1"`
	Priority     string `json:"priority"`
}

// CreateOrder creates a new order
func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default to normal priority if not specified
	priority := models.PriorityNormal
	if req.Priority != "" {
		priority = models.Priority(req.Priority)
	}

	// Create order via OrderService
	order := server.OrderService.CreateOrder(req.CustomerName, req.ProductID, req.RequestedQty, priority)
	if order == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// HandleWebSocket upgrades HTTP connection to WebSocket
func HandleWebSocket(c *gin.Context) {
	ws.ServeWs(server.WebSocketHub, c.Writer, c.Request)
}

// BroadcastRobotUpdate sends robot state updates to all connected WebSocket clients
func BroadcastRobotUpdate(update models.RobotUpdate) {
	if server.WebSocketHub != nil {
		server.WebSocketHub.BroadcastRobotUpdate(update)
	}
}
