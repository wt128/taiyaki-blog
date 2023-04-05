import axios from "axios"

/**
 * 新規記事投稿
 * @param userId
 * @param title
 * @param content 
 */
export const articleService = {
  new: async (userId: number,title: string, content: string) => {
   const params = new URLSearchParams();
    params.append("userId", `${userId}`)
    params.append("title", title)
    params.append("content", content)
    await axios.post(`${import.meta.env.VITE_BACKEND}/article`, params)
  }
}