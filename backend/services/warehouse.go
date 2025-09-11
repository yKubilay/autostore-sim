package services

import (
	"autostore-sim/models"
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
func ExecuteOrders(robots []models.Robot, orders []models.Order) {
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

// ProcessPendringorders automatically assigns and executes pending orders
func ProcessPendingOrders(robots []models.Robot, orders []models.Order) {
	//Find pending orders
	for i, order := range orders {
		if order.Status == "pending" {
			// Try to assign this order
			AssignOrder(robots, orders, i)

			// If it was assigned, execute it
			if orders[i].Status == "assigned" {
				ExecuteOrders(robots, orders)
			}
		}
	}
}
