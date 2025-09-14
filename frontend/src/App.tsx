import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { AppLayout } from '@/components/layout/app-layout'
import DashboardPage from '@/pages/dashboard/page'

function App() {
  return (
    <Router>
      <AppLayout>
        <Routes>
          <Route path="/" element={<DashboardPage />} />
          <Route path="/dashboard" element={<DashboardPage />} />
        </Routes>
      </AppLayout>
    </Router>
  )
}

export default App