package main

import "fmt"

// Robot struct, custom data type
type Robot struct {
	ID     int
	X      int
	Y      int
	Status string
}

// Order struct, represents warehouse orders
type Order struct {
	ID            int
	ItemX         int // Where to pick up the item
	ItemY         int
	DeliveryX     int // Where to pick up an item
	DeliveryY     int
	Status        string
	AssignedRobot int // Which robot is handling this order
}

// Method belonging to Robot
func (r Robot) DisplayInfo() {
	fmt.Printf("Robot %d at position (%d, %d) - Status: %s\n", r.ID, r.X, r.Y, r.Status)
}

// Method, move robot to new position
func (r *Robot) MoveTo(newX, newY int) {
	r.X = newX
	r.Y = newY
	r.Status = "moving"
	fmt.Printf("Robot %d moved to (%d, %d)\n", r.ID, r.X, r.Y)
}

// Method for Order, display, order info
func (o Order) DisplayInfo() {
	fmt.Printf("Order %d: Pick from (%d, %d) deliver to (%d, %d) - Status: %s\n",
		o.ID, o.ItemX, o.ItemY, o.DeliveryX, o.DeliveryY, o.Status)
}

func main() {
	fmt.Println("Starting AutoStore Warehouse Simulation")

	// Create multiple robots using slice
	robots := []Robot{
		{ID: 1, X: 0, Y: 0, Status: "idle"},
		{ID: 2, X: 1, Y: 1, Status: "idle"},
		{ID: 3, X: 2, Y: 2, Status: "idle"},
	}

	// Create some orders
	orders := []Order{
		{ID: 1, ItemX: 4, ItemY: 2, DeliveryX: 0, DeliveryY: 0, Status: "pending", AssignedRobot: 0},
		{ID: 2, ItemX: 3, ItemY: 5, DeliveryX: 1, DeliveryY: 1, Status: "pending", AssignedRobot: 0},
	}

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
	assignorder(robots, orders, 0)
	assignorder(robots, orders, 1)

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
	executeOrders(robots, orders)

	fmt.Println("\nRobots after moving to pickup:")
	for _, robot := range robots {
		robot.DisplayInfo()
	}

	fmt.Println("\nOrders in progress:")
	for _, order := range orders {
		order.DisplayInfo()
	}

	startWarehouse()
	fmt.Println("Warehouse is running!")
}

func startWarehouse() {
	fmt.Println("Initializing robots...")
	fmt.Println("Setting up warehouse grid...")
	fmt.Println("Ready for orders!")
}

// Function to assign an order to available robot
func assignorder(robots []Robot, orders []Order, orderIndex int) {
	// Find first available robot
	for i, robot := range robots {
		if robot.Status == "idle" {
			// Assign order to robot
			orders[orderIndex].Status = "assigned"
			orders[orderIndex].AssignedRobot = robot.ID
			robots[i].Status = "assigned"

			fmt.Printf("ASsigned order %d to Robot %d\n", orders[orderIndex].ID, robot.ID)
			return
		}
	}
	fmt.Println("No available robots!")

}

// Function to move assigned robots to their pickup location
func executeOrders(robots []Robot, orders []Order) {
	for i, robot := range robots {
		if robot.Status == "assigned" {
			// Find which order the robot is assigned to
			for j, order := range orders {
				if order.AssignedRobot == robot.ID && order.Status == "assigned" {
					// Move robot to pickup location
					robots[i].MoveTo(order.ItemX, order.ItemY)
					robots[i].Status = "picking"
					orders[j].Status = "in_progress"

					fmt.Printf("Robot %d is now picking up item for Order %d\n", robot.ID, order.ID)
					break
				}
			}
		}
	}
}
