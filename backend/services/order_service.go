package services

import (
	"autostore-sim/backend/models"
	"fmt"
	"math/rand"
)

// OrderService handles order processing and robot assignment
type OrderService struct {
	orderQueue     *models.OrderQueue
	productService *ProductService
	warehouse      *models.SafeWarehouse
}

// NewOrderService creates a new order service
func NewOrderService(productService *ProductService, warehouse *models.SafeWarehouse) *OrderService {
	return &OrderService{
		orderQueue:     models.NewOrderQueue(),
		productService: productService,
		warehouse:      warehouse,
	}
}

// GenerateRandomOrder creates a realistic customer order
func (os *OrderService) GenerateRandomOrder() *models.Order {
	// Random customer names (auto repair shops)
	customers := []string{
		"Smith Auto Repair", "QuickFix Motors", "Downtown Garage",
		"Highway Service Center", "Metro Auto Parts", "City Car Care",
		"Precision Automotive", "Express Auto Repair",
	}

	// Get random product from catalog
	products := os.productService.GetAllProducts()
	if len(products) == 0 {
		return nil
	}

	randomProduct := products[rand.Intn(len(products))]
	randomCustomer := customers[rand.Intn(len(customers))]

	// Random quantity (1-5 items for realistic orders)
	requestedQty := rand.Intn(5) + 1

	// Random priority (80% normal, 15% urgent, 5% express)
	var priority models.Priority
	priorityRoll := rand.Intn(100)
	switch {
	case priorityRoll < 80:
		priority = models.PriorityNormal
	case priorityRoll < 95:
		priority = models.PriorityUrgent
	default:
		priority = models.PriorityExpress
	}

	return os.orderQueue.AddOrder(randomCustomer, randomProduct.ID, requestedQty, priority)
}

// AssignAvailablePort assigns a delivery port for the order
func (os *OrderService) AssignAvailablePort() models.Position {
	// Use north edge ports (y=0) - randomly pick one
	portX := rand.Intn(os.warehouse.Width) // 0-7 for 8x8 warehouse
	return models.Position{X: portX, Y: 0, Z: 0}
}

// ProcessPendingOrders assigns robots to pending orders
func (os *OrderService) ProcessPendingOrders(robots []*models.Robot) {
	pendingOrders := os.orderQueue.GetPendingOrders()

	for _, order := range pendingOrders {
		// Find available robot
		availableRobot := os.findAvailableRobot(robots)
		if availableRobot == nil {
			continue // No robots available
		}

		// Find product in warehouse
		productLocation := os.findProductInWarehouse(order.ProductID, order.RequestedQty)
		if productLocation == nil {
			// Mark order as failed - no stock
			os.updateOrderStatus(order.ID, models.OrderFailed)
			fmt.Printf("Order %d failed - insufficient stock for product %d\n", order.ID, order.ProductID)
			continue
		}

		// Get actual order pointer from queue (not the loop copy)
		actualOrder := os.orderQueue.GetOrderByID(order.ID)
		if actualOrder == nil {
			continue
		}

		// Assign robot and update order
		os.assignRobotToOrder(availableRobot, actualOrder, *productLocation)
	}
}

// findAvailableRobot returns first idle robot
func (os *OrderService) findAvailableRobot(robots []*models.Robot) *models.Robot {
	for _, robot := range robots {
		if robot.Status == "idle" {
			return robot
		}
	}
	return nil
}

// findProductInWarehouse locates product with sufficient quantity
func (os *OrderService) findProductInWarehouse(productID int, requiredQty int) *models.Position {
	// Search through warehouse grid for this product
	for x := 0; x < os.warehouse.Width; x++ {
		for y := 0; y < os.warehouse.Height; y++ {
			for z := 0; z < os.warehouse.Levels; z++ {
				os.warehouse.Mutex.RLock()
				cell := os.warehouse.Grid[x][y][z]
				os.warehouse.Mutex.RUnlock()

				if cell.CanFulfill(productID, requiredQty) {
					return &models.Position{X: x, Y: y, Z: z}
				}
			}
		}
	}
	return nil // Product not found or insufficient quantity
}

// assignRobotToOrder sends pick command to robot
func (os *OrderService) assignRobotToOrder(robot *models.Robot, order *models.Order, productLocation models.Position) {
	// Assign delivery port
	order.DeliveryPort = os.AssignAvailablePort()
	order.AssignedRobot = robot.ID
	os.updateOrderStatus(order.ID, models.OrderAssigned)

	// Send pick command to robot
	pickCommand := models.RobotCommand{
		Type:    "pick",
		X:       productLocation.X,
		Y:       productLocation.Y,
		Z:       productLocation.Z,
		OrderID: order.ID,
	}

	robot.Commands <- pickCommand
	fmt.Printf("Assigned Order %d to Robot %d - pick from (%d,%d,%d)\n",
		order.ID, robot.ID, productLocation.X, productLocation.Y, productLocation.Z)
}

// updateOrderStatus updates order status
func (os *OrderService) updateOrderStatus(orderID int, status models.OrderStatus) {
	order := os.orderQueue.GetOrderByID(orderID)
	if order != nil {
		order.Status = status
	}
}

// GetActiveOrders returns all non-completed orders
func (os *OrderService) GetActiveOrders() []models.Order {
	var active []models.Order
	for _, order := range os.orderQueue.Orders {
		if order.Status != models.OrderCompleted && order.Status != models.OrderFailed {
			active = append(active, order)
		}
	}
	return active
}

// CreateOrder creates a new order and adds it to the queue
func (os *OrderService) CreateOrder(customerName string, productID int, requestedQty int, priority models.Priority) *models.Order {
	// Validate product exists
	product := os.productService.GetProductByID(productID)
	if product == nil {
		return nil
	}

	// Add order through the OrderQueue (which handles ID generation and initialization)
	return os.orderQueue.AddOrder(customerName, productID, requestedQty, priority)
}
