import './App.css'
import { Button } from './components/ui/button'
import { ArrowRight } from 'lucide-react'
import confetti from 'canvas-confetti'


function App() {
  const handleClick = () => {
    confetti({
      particleCount: 100,
      spread: 70,
      origin: { y: 0.6 }
    })
  }

  return (
    <div className="min-h-screen flex items-center justify-center px-4">
      <div className="max-w-3xl text-center space-y-8">
        <h1 className="text-6xl font-bold tracking-tight">
          Making it easier to learn anything
        </h1>

        <Button size="lg" className="text-lg" onClick={handleClick}>
          Start Learning <ArrowRight className="ml-2 h-5 w-5" />
        </Button>
      </div>
    </div>
  )
}

export default App
