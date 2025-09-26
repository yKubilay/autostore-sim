package models

// Product represents an auto part stored in the warehouse
type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`        // "Air Filter Honda Civic"
	SKU         string   `json:"sku"`         // "AF-HC-2023"
	Category    Category `json:"category"`    // Engine, Brakes, etc.
	Brand       string   `json:"brand"`       // "Bosch", "ACDelco", etc.
	VehicleYear int      `json:"vehicle_year"` // 2020, 2019, etc.
	VehicleMake string   `json:"vehicle_make"` // "Honda", "Toyota", etc.
	Price       float64  `json:"price"`       // 29.99
	Weight      float64  `json:"weight_kg"`   // 0.5 kg
	Position    Position `json:"position"`    // Where it's stored in warehouse
}

// Category represents product categories for auto parts
type Category string

const (
	CategoryEngine     Category = "engine"
	CategoryBrakes     Category = "brakes"
	CategoryElectrical Category = "electrical"
	CategoryFilters    Category = "filters"
	CategoryLighting   Category = "lighting"
	CategoryMaintenance Category = "maintenance"
)

// Position represents a 3D coordinate in the warehouse grid
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// ProductCatalog holds our inventory of products
type ProductCatalog struct {
	Products map[int]*Product `json:"products"` // Map of product ID -> Product
	NextID   int              `json:"next_id"`  // For generating new product IDs
}

// NewProductCatalog creates a new product catalog
func NewProductCatalog() *ProductCatalog {
	return &ProductCatalog{
		Products: make(map[int]*Product),
		NextID:   1,
	}
}

// AddProduct adds a product to the catalog
func (pc *ProductCatalog) AddProduct(name, sku, brand, vehicleMake string,
	vehicleYear int, category Category, price, weight float64) *Product {

	product := &Product{
		ID:          pc.NextID,
		Name:        name,
		SKU:         sku,
		Category:    category,
		Brand:       brand,
		VehicleYear: vehicleYear,
		VehicleMake: vehicleMake,
		Price:       price,
		Weight:      weight,
		Position:    Position{X: -1, Y: -1, Z: -1}, // Unassigned position
	}

	pc.Products[pc.NextID] = product
	pc.NextID++

	return product
}

// GetProduct retrieves a product by ID
func (pc *ProductCatalog) GetProduct(id int) *Product {
	return pc.Products[id]
}

// GetProductsByCategory returns all products in a category
func (pc *ProductCatalog) GetProductsByCategory(category Category) []*Product {
	var products []*Product
	for _, product := range pc.Products {
		if product.Category == category {
			products = append(products, product)
		}
	}
	return products
}