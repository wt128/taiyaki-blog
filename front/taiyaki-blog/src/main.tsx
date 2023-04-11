import { Auth0Provider } from '@auth0/auth0-react';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import { Header } from './Components/Header';
import './index.css';
import { Read as ArticleRead } from './Pages/Article/Read';
import { Edit as ArticleEdit } from './Pages/Article/Edit';

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
  },
  {
    path: 'article/:id',
    element: <ArticleRead />,
  },
  {
    path: 'article/edit',
    element: 
    
      <ArticleEdit />,
  },
]);
ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
      <Auth0Provider
        domain={import.meta.env.VITE_AUTH0_DOMAIN}
        clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
        authorizationParams={{
          redirect_uri: window.location.origin,
        }}
      >
        <Header />
        <RouterProvider router={router} />
      </Auth0Provider>
  </React.StrictMode>
);
