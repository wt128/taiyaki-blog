import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { articleService } from '../../Features/api/article';
import { Skeleton } from '@mui/material';
import { Box, Typography, makeStyles, createStyles, Chip } from '@material-ui/core';
import { ReactMarkdown } from 'react-markdown/lib/react-markdown';
import rehypeHighLight from 'rehype-highlight';
import remarkGfm from 'remark-gfm';
import { Theme } from '@emotion/react';
import CodeBlock from '../../Components/CodeBlock';
import { Author } from '../../Components/Author';

export interface ArticleDetail {
  id: number;
  content: string;
  title: string;
  author: string;
  createdAt: string;
}

const useStyles = makeStyles((_: Theme) =>
  createStyles({
    md: {
      lineHeight: '1em',
    },
  })
);
const useFetchArticle = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [data, setData] = useState<ArticleDetail>({
    id: 0,
    content: '',
    title: '',
    author: '',
    createdAt: ''
  });
  const [isFetching, setIsFetching] = useState(false);
  useEffect(() => {
    if (id !== undefined) {
      if (!id.match(/\d+/)) {
        navigate('/');
        return;
      }
      articleService.one(id).then((data) => setData(data.data));
    }
  }, [id, navigate]);
  return {
    article: data,
    isFetching,
  };
};
export const Read = () => {
  const { article, isFetching } = useFetchArticle();
  const classes = useStyles();
  if (isFetching) {
    return <Skeleton variant="rectangular" width={210} height={60} />;
  }
  return (
    <Box sx={{ marginTop: '100' }}>
      <Box width={100} mt={10}>
        <Typography variant="h1" align="center">
          {article.title}
        </Typography>
      </Box>
        <Author 
          author={article.author}
          createdAt={article.createdAt}
        />
        <Chip />
      <Typography align="center">
        <ReactMarkdown
          className={classes.md}
          components={{
            code: CodeBlock,
          }}
          rehypePlugins={[rehypeHighLight]}
          remarkPlugins={[remarkGfm]}
        >
          {article.content}
        </ReactMarkdown>
      </Typography>
    </Box>
  );
};
