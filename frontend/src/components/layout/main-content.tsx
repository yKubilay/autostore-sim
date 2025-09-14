import { Plus, Minus, Maximize2 } from "lucide-react"
import { Button } from "@/components/ui/button"

export function MainContent() {
  return (
    <div className="flex flex-col h-full">
      <div className="p-4 border-b border-border flex items-center justify-between">
        <h1 className="text-xl font-semibold">Warehouse 1</h1>
        <div className="flex items-center gap-2">
          <div className="flex gap-1">
            <Button variant="outline" size="sm">
              2D
            </Button>
            <Button variant="default" size="sm">
              3D
            </Button>
          </div>
        </div>
      </div>

      {/* Main Content Area */}
      <div className="flex-1 p-4">
        <div className="w-full h-full bg-gray-100 rounded-lg flex items-center justify-center relative">
          <div className="text-center text-muted-foreground">
            <div className="text-4xl mb-2">🏭</div>
            <p>Warehouse Visualization</p>
            <p className="text-sm">Main content will go here</p>
          </div>

          {/* Control buttons */}
          <div className="absolute bottom-4 right-4 flex flex-col gap-2">
            <Button size="sm" variant="outline">
              <Plus className="w-4 h-4" />
            </Button>
            <Button size="sm" variant="outline">
              <Minus className="w-4 h-4" />
            </Button>
            <Button size="sm" variant="outline">
              <Maximize2 className="w-4 h-4" />
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}
