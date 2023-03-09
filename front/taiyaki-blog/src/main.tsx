import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import App from './App'
import { Header } from './Components/Header'
import './index.css'
import { ArticleRead } from './Pages/ArticleRead'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />
  },
  {
    path: "article/:id",
    element: <ArticleRead />
  }
])
ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Header />
    <RouterProvider router={router} />
  </React.StrictMode>,
)
