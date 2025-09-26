package services

import (
	"autostore-sim/backend/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// ProductService handles product-related operations
type ProductService struct {
	catalog *models.ProductCatalog
}

// ProductData represents the JSON structure from the data file
type ProductData struct {
	Products []models.Product `json:"products"`
}

// NewProductService creates new product service
func NewProductService() *ProductService {
	return &ProductService{
		catalog: models.NewProductCatalog(),
	}
}

// LoadProductsFromFile simulates loading products from an external API
// In real system, this would be: LoadProductsFromAPI()
func (ps *ProductService) LoadProductsFromFile(filename string) error {
	// Get path for the JSON file
	dataPath := filepath.Join("data", filename)

	// Read the JSON file
	jsonData, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read product data: %w", err)
	}

	var data ProductData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("failed to parse product JSON: %w", err)
	}

	// Add products to catalog
	for _, product := range data.Products {
		// Create a copy to avoid pointer issues
		productCopy := product
		ps.catalog.Products[product.ID] = &productCopy

		//Update NextID to be higher than any existing ID
		if product.ID >= ps.catalog.NextID {
			ps.catalog.NextID = product.ID + 1
		}
	}

	fmt.Printf("Loaded %d products from %s\n", len(data.Products), filename)
	return nil
}

// GetAllProducts return all products in the catalog
func (ps *ProductService) GetAllProducts() []*models.Product {
	var products []*models.Product
	for _, product := range ps.catalog.Products {
		products = append(products, product)
	}
	return products
}

// GetProductByID retrieves a product by ID
func (ps *ProductService) GetProductByID(id int) *models.Product {
	return ps.catalog.GetProduct(id)
}

// GetProductsByCategory returns products in a specific category
func (ps *ProductService) GetProductsByCategory(category models.Category) []*models.Product {
	return ps.catalog.GetProductsByCategory(category)
}

// PlaceProductsInWarehouse randomly assigning products to warehouse positions
func (ps *ProductService) PlaceProductsInWarehouse(warehouse *models.SafeWarehouse) error {
	products := ps.GetAllProducts()

	// Get all available positions excluding edge positions for ports
	availablePositions := ps.getStoragePositions(warehouse)

	// Shuffle positions for random placement
	rand.Shuffle(len(availablePositions), func(i, j int) {
		availablePositions[i], availablePositions[j] = availablePositions[j], availablePositions[i]
	})

	// Place products in random positions
	for i, product := range products {
		if i >= len(availablePositions) {
			fmt.Printf("Warning: More products (%d) than available storage positions (%d)\n",
				len(products), len(availablePositions))
			break
		}

		position := availablePositions[i]
		product.Position = position
		fmt.Printf("Placed %s at position (%d, %d, %d)\n",
			product.Name, position.X, position.Y, position.Z)
	}

	return nil
}

// getStoragePositions returns all positions excluding edge ports
func (ps *ProductService) getStoragePositions(warehouse *models.SafeWarehouse) []models.Position {
	var positions []models.Position

	for x := 0; x < warehouse.Width; x++ {
		for y := 0; y < warehouse.Height; y++ {
			for z := 0; z < warehouse.Levels; z++ {
				// Skip edge positions (ports) and only skip y=0 which is north edge
				if y == 0 {
					continue // This is a port position
				}

				positions = append(positions, models.Position{X: x, Y: y, Z: z})
			}
		}
	}

	return positions
}

// GetProductCount returns total number of products
func (ps *ProductService) GetProductCount() int {
	return len(ps.catalog.Products)
}
