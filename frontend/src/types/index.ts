// AutoStore specific types
export interface Robot {
  id: string
  status: 'idle' | 'picking' | 'charging' | 'maintenance'
  batteryLevel: number
  currentTask?: string
  location: { x: number; y: number; z: number }
}

export interface Bin {
  id: string
  location: { x: number; y: number; z: number }
  contents: InventoryItem[]
  capacity: number
  currentWeight: number
}

export interface InventoryItem {
  id: string
  sku: string
  name: string
  quantity: number
  location: string
  lastUpdated: string
}

export interface Order {
  id: string
  status: 'pending' | 'picking' | 'completed' | 'cancelled'
  items: OrderItem[]
  priority: 'low' | 'medium' | 'high'
  createdAt: string
  estimatedCompletion?: string
}

export interface OrderItem {
  id: string
  sku: string
  name: string
  quantity: number
  pickedQuantity: number
  binLocation?: string
}

export interface SystemMetrics {
  totalOrders: number
  completedOrders: number
  activeRobots: number
  avgOrderTime: number
  systemEfficiency: number
}