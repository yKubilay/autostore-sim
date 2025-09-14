"use client"

import { Home, Clock, GitBranch, Settings, Bell, Users, FileText, Bookmark, Search } from "lucide-react"
import { Input } from "@/components/ui/input"

const navigationItems = [
  { icon: Home, label: "Home" },
  { icon: Clock, label: "Recent" },
  { icon: GitBranch, label: "Operations" },
  { icon: Users, label: "Teams" },
  { icon: Bell, label: "Notifications" },
  { icon: FileText, label: "Reports" },
  { icon: Bookmark, label: "Bookmarks" },
  { icon: Settings, label: "Settings" },
]

export function Sidebar() {
  return (
    <div className="flex flex-col bg-white fixed left-0 top-0 h-full z-20 w-12">
      {/* Header with logo and search */}
      <div className="h-16 bg-white flex items-center px-4 fixed top-0 left-0 right-0 z-30">
        {/* Logo on the left */}
        <div className="flex items-center gap-3">
          <div className="w-8 h-8 bg-black rounded-lg flex items-center justify-center">
            <div className="w-4 h-4 bg-white rounded-sm"></div>
          </div>
          <span className="font-semibold text-lg">Autostore Simulator</span>
        </div>

        {/* Search bar positioned where orders column ends */}
        <div className="flex items-center justify-start">
          <div className="relative w-96 ml-40">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground w-4 h-4" />
            <Input placeholder="Search for customer orders, jobs, vehicles and assets" className="pl-10" />
          </div>
        </div>
      </div>

      {/* Navigation - always collapsed */}
      <div className="flex flex-col py-4 mt-16">
        <nav className="flex flex-col gap-4 px-0">
          {navigationItems.map((item, index) => (
            <button
              key={index}
              className="group relative w-12 h-10 flex items-center justify-center text-muted-foreground hover:text-foreground hover:bg-accent rounded-md transition-colors mx-auto"
              title={item.label}
            >
              <item.icon className="w-5 h-5" />
              
              {/* Tooltip */}
              <div className="absolute left-full ml-2 px-2 py-1 bg-gray-900 text-white text-xs rounded opacity-0 group-hover:opacity-100 transition-opacity pointer-events-none whitespace-nowrap z-50">
                {item.label}
              </div>
            </button>
          ))}
        </nav>
      </div>
    </div>
  )
}