import React, { useEffect, useState } from 'react';

/**
 * 簡易的な認証情報の型のサンプル
 * @type {{userId: number | null, sessionId: string}}
 */
type AuthInfo = {
  userId: number | null;
  sessionId: string;
};

// ログイン状態のContext
export const LoggedInContext = React.createContext<number | null>(null);

// 認証情報と認証情報セットのContext
export const AuthInfoContext = React.createContext<
  [AuthInfo, React.Dispatch<React.SetStateAction<AuthInfo>>]
>([{ userId: null, sessionId: "" }, () => {}]);

/**
 * デフォルトのAuthInfoを取得
 * ローカルストレージから取得できた場合はその値をパース
 * 取得できない場合は空の情報を返す
 * @returns
 */
const getDefaultAuthInfo = (): AuthInfo => {
  const defaultAuthInfo = document.cookie.match(
    new RegExp('sessionId\=([^\;]*)\;*')
  );
  if (defaultAuthInfo) {
    return { userId: 1, sessionId: defaultAuthInfo[1] };
  } else {
    return { userId: null, sessionId: "" };
  }
};

const setAuthInfoToCookie = (authInfo: AuthInfo) => {
  //const autoInfoStringfy = JSON.stringify(authInfo);
  document.cookie = `sessionId=${authInfo.sessionId}`;
};
export const AuthContextProvider: React.FC<{}> = (props) => {
  const [loggedIn, setLoggedIn] = useState<boolean>(false);
  const [authInfo, setAuthInfo] = useState<AuthInfo>(getDefaultAuthInfo());
  useEffect(() => {
    if (authInfo?.userId) {
      setAuthInfoToCookie(authInfo)
      setLoggedIn(true);
    } else {
      setLoggedIn(false);
    }
  }, [authInfo]);

  return (
    <LoggedInContext.Provider value={loggedIn}>
      <AuthInfoContext.Provider value={[authInfo, setAuthInfo]}>
        {props.children}
      </AuthInfoContext.Provider>
    </LoggedInContext.Provider>
  );
};
