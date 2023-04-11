import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { articleService } from '../../Features/api/article';
import { Skeleton } from '@mui/material';
import { Box, Typography } from '@material-ui/core';
import { ReactMarkdown } from 'react-markdown/lib/react-markdown';

export interface ArticleDetail {
  id: number;
  content: string;
  title: string;
  author: string;
}

const useFetchArticle = () => {
  const { id } = useParams();
  const [data, setData] = useState<ArticleDetail>({
    id: 0,
    content: "",
    title: "",
    author: ""
  });
  const [isFetching, setIsFetching] = useState(false);
  useEffect(() => {
    if (id !== undefined) {
      articleService.one(id).then((data) => setData(data.data));
    }
  }, []);
  return {
    article: data,
    isFetching,
  };
};
export const Read = () => {
  const { article, isFetching } = useFetchArticle();
  if (isFetching) {
    return <Skeleton variant="rectangular" width={210} height={60} />;
  }
  return (
    <Box sx={{ width: '100%', maxWidth: 500 }}>
      <Typography variant="h1">{article.title}</Typography>
      <Typography variant="body1">
        <ReactMarkdown>{article.content}</ReactMarkdown>
      </Typography>
    </Box>
  );
};
