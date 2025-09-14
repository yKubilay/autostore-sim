"use client"

import { ReactNode } from "react"
import { OrdersColumn } from "@/components/orders-column"
import { MainContent } from "@/components/layout/main-content"
import { Sidebar } from "@/components/ui/sidebar"

interface AppLayoutProps {
  children?: ReactNode
}

export function AppLayout({ children }: AppLayoutProps) {
  return (
    <div className="flex h-screen bg-white">
      <Sidebar />

      <div className="flex flex-1 gap-4 pl-1 pr-4 pb-4 pt-20 ml-12">
        {children ? (
          children
        ) : (
          <>
            <div className="w-80 bg-gray-100 rounded-lg border border-border shadow-md">
              <OrdersColumn />
            </div>

            <div className="flex-1 bg-white rounded-lg border border-border shadow-md">
              <MainContent />
            </div>
          </>
        )}
      </div>
    </div>
  )
}