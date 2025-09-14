import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Activity, Package, Users, TrendingUp } from "lucide-react"

export function DashboardLayout() {
  return (
    <div className="flex-1 space-y-4 p-6">
      <div className="flex items-center justify-between space-y-2">
        <h2 className="text-3xl font-bold tracking-tight">AutoStore Dashboard</h2>
      </div>
      
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Active Orders</CardTitle>
            <Package className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">1,234</div>
            <p className="text-xs text-muted-foreground">
              +20.1% from last hour
            </p>
          </CardContent>
        </Card>
        {/* Add more cards as needed */}
      </div>
    </div>
  )
}