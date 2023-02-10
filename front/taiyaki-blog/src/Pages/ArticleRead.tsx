import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

export interface ArticleDetail {
  id: number,
  body: string,
  author: string,
}
const initArticleDetail = {
  id: 0,
  body: "",
  author: "",
}

const hoge = {
  id: 0,
  body: "",
  author: "",
}
export const ArticleRead = () => {
  const {id} = useParams()
  const [article, setArticle] = useState<ArticleDetail>(initArticleDetail);
  useEffect(() => {
   /* const data = async () => {
      const res = await fetch('http://localhost/');
      return res;
    }; */
    setArticle(hoge);
  });
  return <h1>{id}{article.author}</h1>;
};
