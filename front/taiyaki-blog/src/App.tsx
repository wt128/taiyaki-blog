import { Auth0Provider } from '@auth0/auth0-react';
import './App.css';
import { Read as ArticleRead } from './Pages/Article/Read';
import { Edit as ArticleEdit } from './Pages/Article/Edit';
import { List as ArticleList } from './Pages/Article/List';
import { Header } from './Components/Header';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

const router = createBrowserRouter([
  {
    path: '/',
    element: <ArticleList />,
  },
  {
    path: 'article/:id',
    element: <ArticleRead />,
  },
  {
    path: 'article/new',
    element: <ArticleEdit />,
  },
]);
function App() {
  return (
    <>
      <div className="App">
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
      </div>
    </>
  );
}

export default App;
