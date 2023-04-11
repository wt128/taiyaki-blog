import { useAuth0 } from "@auth0/auth0-react";
import { useState } from "react";

export const useAuth0Token = () => {
  const { isAuthenticated, getAccessTokenSilently } = useAuth0();
  const [accessToken, setAccessToken] = useState("");
  const fetchToken = async () => {
    // JWTを取得して状態に保存する
    const accessToken = await getAccessTokenSilently()
    setAccessToken(accessToken);
  };
  // ログイン済みの場合のみJWTを取得する
  if (isAuthenticated) {
    fetchToken();
  }

  return accessToken;
};