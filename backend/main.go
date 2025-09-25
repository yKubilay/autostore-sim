package main

import (
	"autostore-sim/backend/handlers"
	"autostore-sim/backend/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting AutoStore Warehouse Simulation")

	// Create thread-safe warehouse
	safeWarehouse := models.GetDefaultSafeWarehouse()

	// Create robots using pointers for goroutines
	robots := []*models.Robot{
		&models.Robot{ID: 1, X: 0, Y: 0, Z: 0, Status: "idle"},
		&models.Robot{ID: 2, X: 1, Y: 1, Z: 0, Status: "idle"},
		&models.Robot{ID: 3, X: 2, Y: 2, Z: 0, Status: "idle"},
	}

	// Set initial robot positions in warehouse grid
	for _, robot := range robots {
		safeWarehouse.SetRobotPosition(robot.X, robot.Y, robot.Z, true)
	}

	// Create done channel for graceful shutdown
	done := make(chan bool)

	// Start robot goroutines
	fmt.Println("Starting robot goroutines:")
	for _, robot := range robots {
		robot.StartRobot(safeWarehouse, done)
	}

	// Display initial state
	fmt.Println("Initial robot positions:")
	for _, robot := range robots {
		robot.DisplayInfo()
	}

	// Test the new goroutine system by sending move commands
	fmt.Println("\nTesting robot movement with channels:")
	time.Sleep(1 * time.Second) // Let robots initialize

	// Send move commands to robots through their channels
	robots[0].Commands <- models.RobotCommand{Type: "move", X: 1, Y: 0, Z: 0}
	robots[1].Commands <- models.RobotCommand{Type: "move", X: 2, Y: 1, Z: 0}
	robots[2].Commands <- models.RobotCommand{Type: "move", X: 3, Y: 2, Z: 0}

	fmt.Println("Move commands sent! Check console for robot responses.")
	time.Sleep(2 * time.Second) // Give robots time to move

	fmt.Println("\nRobot positions after movement:")
	for _, robot := range robots {
		fmt.Printf("Robot %d at (%d, %d, %d) - Status: %s\n",
			robot.ID, robot.X, robot.Y, robot.Z, robot.Status)
	}

	fmt.Println("Warehouse is running!")
	fmt.Println("API available at http://localhost:8080")

	// Start web server in a separate goroutine
	go startWebServer()

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
