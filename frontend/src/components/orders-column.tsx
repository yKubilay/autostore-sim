export function OrdersColumn() {
  return (
    <div className="flex flex-col h-full">
      {/* Header */}
      <div className="p-4 border-b border-border">
        <h2 className="font-semibold text-gray-900 mb-2">Reported operations</h2>

        {/* Location Selector */}
        <div className="flex items-center gap-2 p-2 bg-white rounded-lg mb-4 shadow-sm">
          <div className="w-3 h-3 bg-green-500 rounded-full"></div>
          <span className="text-sm font-medium text-gray-900">Production facility</span>
          <span className="text-sm text-gray-600">Floor 1</span>
        </div>

        {/* Tabs */}
        <div className="flex gap-1">
          <button
            className="px-3 py-1.5 rounded-md text-sm font-medium"
            style={{
              backgroundColor: "#1f2937",
              color: "#ffffff",
              border: "none",
            }}
          >
            Missions
          </button>
          <button className="px-3 py-1.5 text-gray-600 hover:text-gray-900 text-sm">Robots</button>
          <button className="px-3 py-1.5 text-gray-600 hover:text-gray-900 text-sm">Logs</button>
        </div>
      </div>

      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {[1, 2, 3, 4].map((index) => (
          <div key={index} className="bg-white rounded-lg p-4 border border-border shadow-sm h-24">
            {/* Empty card - user will add content later */}
          </div>
        ))}
      </div>
    </div>
  )
}
