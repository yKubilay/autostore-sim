package main

import (
	"autostore-sim/backend/handlers"
	"autostore-sim/backend/services"
	"autostore-sim/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting AutoStore Warehouse Simulation")

	// Create multiple robots using slice
	robots := []models.Robot{
		{ID: 1, X: 0, Y: 0, Status: "idle"},
		{ID: 2, X: 1, Y: 1, Status: "idle"},
		{ID: 3, X: 2, Y: 2, Status: "idle"},
	}

	// Create some orders
	orders := []models.Order{
		{ID: 1, ItemX: 4, ItemY: 2, DeliveryX: 0, DeliveryY: 0, Status: "pending", AssignedRobot: 0},
		{ID: 2, ItemX: 3, ItemY: 5, DeliveryX: 1, DeliveryY: 1, Status: "pending", AssignedRobot: 0},
	}

	// Create workstations (ports)
	workstations := []models.Workstation{
		{ID: 1, X: 0, Y: 4, Status: "available"}, // Edge of warehouse
		{ID: 2, X: 4, Y: 4, Status: "available"}, // Another edge
	}

	// Initialize warehouse data for API
	handlers.SetWarehouseData(robots, orders, workstations)

	// Start web server in a separate goroutine
	go startWebServer()

	// Display initial state
	fmt.Println("Initial robot positions:")
	for _, robot := range robots {
		robot.DisplayInfo()
	}

	fmt.Println("\nPending orders:")
	for _, order := range orders {
		order.DisplayInfo()
	}

	// Assign orders to robots
	fmt.Println("\nAssigning orders:")
	services.AssignOrder(robots, orders, 0)
	services.AssignOrder(robots, orders, 1)

	// Display updated state
	fmt.Println("\nAfter assignment:")
	for _, robot := range robots {
		robot.DisplayInfo()
	}

	fmt.Println("\nUpdated orders:")
	for _, order := range orders {
		order.DisplayInfo()
	}

	// Execute the orders
	fmt.Println("\nExecuting orders:")
	services.ExecuteOrders(robots, orders)

	fmt.Println("\nRobots after moving to pickup:")
	for _, robot := range robots {
		robot.DisplayInfo()
	}

	fmt.Println("\nOrders in progress:")
	for _, order := range orders {
		order.DisplayInfo()
	}

	fmt.Println("Warehouse is running!")
	fmt.Println("API available at http://localhost:8080")

	// Start automatic order processor in a separate goroutine
	go startOrderProcessor()

	// Keep main running
	select {}
}

// startOrderProcessor runs in a goroutine and processes pending orders periodically
func startOrderProcessor() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Process orders directly on warehouse data
		handlers.ProcessWarehouseOrders()
	}
}

func startWebServer() {
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		api.GET("/robots", handlers.GetRobots)
		api.GET("/orders", handlers.GetOrders)
		api.GET("/workstations", handlers.GetWorkstations)
		api.GET("/status", handlers.GetWarehouseStatus)

		// POST endpoint to create orders
		api.POST("/orders", handlers.CreateOrder)
	}
	fmt.Println("Web server starting on :8080")
	r.Run(":8080")
}
