import { useAuth0 } from "@auth0/auth0-react";
import { useEffect, useState } from "react";

export const useAuth0Token = () => {
  const [accessToken, setAccessToken] = useState("");
  const { isAuthenticated, user, getAccessTokenSilently } = useAuth0();
  console.log(user)
  const fetchToken = async () => {
    // JWTを取得して状態に保存する
    const accessToken = await getAccessTokenSilently()
    setAccessToken(accessToken);
  };
  useEffect(() => {
    if (isAuthenticated || !accessToken) {
      fetchToken();
    }
  }, [isAuthenticated, user?.sub])
  // ログイン済みの場合のみJWTを取得する
  return accessToken;
};