import axios from "axios"
import { ArticleDetail } from "../types/article";
import { useAuth0Token } from "../../Utils/auth";

export const articleService = {
  /**
   * 新規記事投稿
   * @param userId
   * @param title
   * @param content 
   * @param token
   */
  new: async (userId: number,title: string, content: string, token: string) => {
    const params = new URLSearchParams();
    params.append("userId", `${userId}`)
    params.append("title", title)
    params.append("content", content)
    await axios.post(`${import.meta.env.VITE_BACKEND}/article`, params, {
       headers: {
        "authorization": `Bearer ${token}`,
       }
      })
  },

/**
 * 特定の記事取得
 * @param id
 */
  one: (id: string) => {
    return axios.get<ArticleDetail>(`${import.meta.env.VITE_BACKEND}/article/${id}`)
  },

}