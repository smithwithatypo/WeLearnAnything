import { useState } from 'react'
import './App.css'
import { Button } from './components/ui/button'
import { Card } from './components/ui/card'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Card>
        <Button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </Button>
      </Card>
    </>
  )
}

export default App
