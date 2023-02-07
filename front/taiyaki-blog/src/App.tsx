import { useState } from 'react'
import reactLogo from './assets/react.svg'
import {Header} from "./Components/Header"
import './App.css'
import { ArticleList } from './Pages/ArticleList'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <Header />
    <div className="App">
      <ArticleList />
    </div>
    </>
  )
}

export default App
