# AutoStore Warehouse Simulation - System Architecture

## Overview
This document describes the architecture of the AutoStore warehouse simulation system, built with Go backend and React frontend.

## System Architecture

### API - External Interface

### SERVICES - Business Logic

### MODELS - Data Structures

🚀 **main.go**
**Entry Point & Orchestration**
• Initializes all services
• Starts robot goroutines
• Manages system lifecycle

**Warehouse**
**Thread-safe 3D grid**
• 8x8x5 storage cells
• Collision detection
• Inventory tracking

**Robot**
**Autonomous worker units**
• Channel-based commands
• Realistic AutoStore timing
• Status management

**Order**
**Customer requests**
• Priority levels
• Status tracking
• Customer & product info

**Product**
**Auto parts inventory**
• Categories & specifications
• Storage positions
• Realistic quantities

**Workstation**
**Delivery endpoints**
• Status tracking
• Order completion

**ProductService**
**Product Management**
• JSON catalog loading
• Warehouse placement
• Inventory operations

**OrderService**
**Order Lifecycle**
• Random order generation
• Robot assignment
• Status updates

**WarehouseService**
**Robot Operations**
• Order assignment
• Robot movement
• Delivery coordination

**handlers/api.go**
**REST Endpoints**
• GET /robots, /orders
• POST /orders
• System status

**Gin Web Server**
**HTTP Router**
• Port :8080
• JSON responses

**products.json**
**Product Catalog**
• Auto parts database
• Categories & specs

## Component Responsibilities

### 🏗️ Architecture Layers

**Entry Point**
main.go: System orchestration, service initialization, goroutine management, lifecycle control

**Models (Data Layer)**
- Warehouse: Thread-safe 3D grid with collision detection and inventory tracking
- Robot: Autonomous units with channel-based commands and realistic AutoStore timing
- Order: Customer requests with priority levels and comprehensive status tracking
- Product: Auto parts inventory with categories, specifications, and storage positions
- Workstation: Delivery endpoints for order completion and port management

**Services (Business Logic)**
- ProductService: Product catalog management, JSON loading, warehouse placement, inventory operations
- OrderService: Complete order lifecycle, robot assignment, status tracking, automated generation
- WarehouseService: Robot operation coordination, movement control, delivery management

**API (External Interface)**
- Handlers: REST endpoints for robots, orders, and system status with full CRUD operations
- Gin Server: HTTP router serving JSON responses on port 8080 with middleware support

### 🔄 Key Data Flows
- **Order Processing**: API → OrderService → Robot Assignment → Warehouse Operations
- **Product Management**: JSON Catalog → ProductService → Warehouse Placement
- **Robot Operations**: OrderService Commands → Robot Channels → Warehouse Updates
- **Status Monitoring**: All Components → API Responses → Frontend Display

### ⚡ Real-time Features
- **Goroutine-based Robots**: Each robot runs independently with channel communication
- **Thread-safe Operations**: Concurrent access to warehouse grid with mutex protection
- **Realistic Timing**: Based on actual AutoStore specifications (3.1 m/s horizontal, 1.6 m/s lift)
- **Automatic Order Generation**: Continuous order creation for demonstration purposes

### 🎯 Design Principles
- **Separation of Concerns**: Clear boundaries between data, business logic, and API layers
- **Concurrency Safety**: Thread-safe operations for multi-robot simulation
- **Realistic Simulation**: Based on actual AutoStore warehouse specifications
- **Extensible Architecture**: Easy to add new features and integrate with frontend