package services

import (
	"autostore-sim/backend/models"
	"fmt"
)

// AssignOrder assigns an order to an available robot
func AssignOrder(robots []models.Robot, orders []models.Order, orderIndex int) {
	// Find first available robot
	for i, robot := range robots {
		if robot.Status == "idle" {
			// Assign order to robot
			orders[orderIndex].Status = "assigned"
			orders[orderIndex].AssignedRobot = robot.ID
			robots[i].Status = "assigned"

			fmt.Printf("Assigned order %d to Robot %d\n", orders[orderIndex].ID, robot.ID)
			return
		}
	}
	fmt.Println("No available robots!")
}

// ExecuteOrders moves assigned robots to their pickup locations
func ExecuteOrders(robots []models.Robot, orders []models.Order, warehouse models.Warehouse) {
	for i, robot := range robots {
		if robot.Status == "assigned" {
			// Find which order the robot is assigned to
			for j, order := range orders {
				if order.AssignedRobot == robot.ID && order.Status == "assigned" {
					// Move robot to pickup location
					if robots[i].MoveToWithBounds(order.ItemX, order.ItemY, 0, warehouse) {
						robots[i].Status = "picking"
						orders[j].Status = "in_progress"
						fmt.Printf("Robot %d is now picking up item for Order %d\n", robot.ID, order.ID)
					} else {
						fmt.Printf("Robot %d cannot reach item at (%d, %d) - out of bounds\n", robot.ID, order.ItemX, order.ItemY)
						// Reset robot and order status when movemvent fails
						robots[i].Status = "idle"
						orders[j].Status = "failed"
					}
					break
				}
			}
		}
	}
}

// ProcessPendringorders automatically assigns and executes pending orders
func ProcessPendingOrders(robots []models.Robot, orders []models.Order, workstations []models.Workstation) {
	warehouse := models.GetDefaultWarehouse() // Create 5x5x3 warehouse
	//Find pending orders
	for i, order := range orders {
		if order.Status == "pending" {
			// Try to assign this order
			AssignOrder(robots, orders, i)
			// If it was assigned, execute it
			if orders[i].Status == "assigned" {
				ExecuteOrders(robots, orders, warehouse)
			}
		}
	}
	// Deliver picked orders to Workstation
	DeliverOrders(robots, orders, workstations, warehouse)

}

// DeliverOrders moves robots with picked items to delivery workstations
func DeliverOrders(robots []models.Robot, orders []models.Order, workstations []models.Workstation, warehouse models.Warehouse) {
	for i, robot := range robots {
		if robot.Status == "picking" {
			for j, order := range orders {
				if order.AssignedRobot == robot.ID && order.Status == "in_progress" {
					for k, workstation := range workstations {
						if workstation.Status == "available" {
							// Use bounded movement to workstation
							if robots[i].MoveToWithBounds(workstation.X, workstation.Y, 0, warehouse) {
								robots[i].Status = "delivering"
								orders[j].Status = "delivered"
								workstations[k].Status = "busy"
								fmt.Printf("Robot %d delivering Order %d to Workstation %d\n",
									robot.ID, order.ID, workstation.ID)
							} else {
								fmt.Printf("Robot %d cannot reach workstation at (%d, %d) - out of bounds\n",
									robot.ID, workstation.X, workstation.Y)
							}
							return
						}
					}
					fmt.Printf("No available workstations for Robot %d\n", robot.ID)
				}
			}
		}
	}
}
