import { Auth0Provider } from '@auth0/auth0-react';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import { Header } from './Components/Header';
import { AuthContextProvider } from './Contexts/AuthContextProvider';
import './index.css';
import { ArticleRead } from './Pages/ArticleRead';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    path: 'article/:id',
    element: <ArticleRead />,
  },
]);
ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <AuthContextProvider>
      <Auth0Provider
        domain={import.meta.env.AUTH0_DOMAIN}
        clientId={import.meta.env.AUTH0_CLIENT_ID}
      >
        <Header />
        <RouterProvider router={router} />
      </Auth0Provider>
    </AuthContextProvider>
  </React.StrictMode>
);
