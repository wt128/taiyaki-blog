import { useContext, useState } from 'react';
import './App.css';
import { AuthInfoContext, LoggedInContext } from './Contexts/AuthContextProvider';
import { ArticleList } from './Pages/ArticleList';

function App() {
  const isLoggedIn = useContext(LoggedInContext)
  const [authInfo, setAuthInfo] = useContext(AuthInfoContext)
  return (
    <>
      <div className="App">
        <div>
          {isLoggedIn ? `ID: ${authInfo?.userId}` : "not login"}
          <button
            onClick={()=> setAuthInfo({userId: 1, sessionId: "abcdefg"})}>
              login
            </button>
            <button
              onClick={()=> setAuthInfo({userId: null, sessionId: ""})}/>
        </div>
        <ArticleList />
      </div>
    </>
  );
}

export default App;
