import { Auth0Provider, useAuth0 } from '@auth0/auth0-react';
import './App.css';
import { Read as ArticleRead } from './Pages/Article/Read';
import { Edit as ArticleEdit } from './Pages/Article/Edit';
import { List as ArticleList } from './Pages/Article/List';
import { Header } from './Components/Header';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import { useEffect } from 'react';

// const router = createBrowserRouter([
//   {
//     path: '/',
//     element: <ArticleList />,
//   },
//   {
//     path: 'article/:id',
//     element: auth0.isAuthenticated ? <ArticleRead /> : <ArticleList />,
//   },
//   {
//     path: 'article/new',
//     element: auth0.isAuthenticated ?  <ArticleEdit /> : <ArticleList />,
//   },
// ]);

const route = (isAuthenticated: Boolean) => {
  return createBrowserRouter([
    {
      path: '/',
      element: <ArticleList />,
    },
    {
      path: 'article/:id',
      element: isAuthenticated ? <ArticleRead /> : <ArticleList />,
    },
    {
      path: 'article/new',
      element: isAuthenticated ?  <ArticleEdit /> : <ArticleList />,
    },
  ])
}
function App() {
  const auth0 = useAuth0()
  const isAuthenticated = auth0.isAuthenticated
  const router = route(isAuthenticated)
  return (
    <>
      <div className="App">
        <Auth0Provider
          domain={import.meta.env.VITE_AUTH0_DOMAIN}
          clientId={import.meta.env.VITE_AUTH0_CLIENT_ID}
          authorizationParams={{
            redirect_uri: window.location.origin,
            audience: import.meta.env.VITE_BACKEND,
            scope: "*"
          }}
        >
            <Header />
          <RouterProvider router={router} />
        </Auth0Provider>
      </div>
    </>
  );
}

export default App;
