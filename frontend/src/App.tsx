import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { AppLayout } from '@/components/layout/app-layout'

function App() {
  return (
    <Router>
      <AppLayout>
        {/* Remove the Routes to always show default layout */}
      </AppLayout>
    </Router>
  )
}

export default App