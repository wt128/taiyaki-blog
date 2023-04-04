import {
  Box,
  createStyles,
  makeStyles,
  Theme,
  Snackbar,
  CircularProgress,
} from '@material-ui/core';
import { ArticleCard } from './ArticleCard';
import axios from 'axios';
import { useEffect, useState } from 'react';

export interface Article {
  id: number;
  title: string;
  content: string;
  explain: string;
  userId: number;
  
}
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      margin: '1.5rem',
      flexWrap: 'wrap',
      gap: '10px 1%',
      width: 'calc(100% - 40px)',
      justifyContent: 'space-between',
    },
    loading: {
      margin: '0 auto'
    }
  })
);

const useFetchArticleList = () => {
  const [data, setData] = useState<Article[]>([]);
  const [isFetching, setIsFetching] = useState(false);
  useEffect(() => {
    setIsFetching(true);
    axios
      .get<Article[]>(`${import.meta.env.VITE_BACKEND}/article`)
      .then((res) => {
        setData(res.data)
        console.log(data)
      })
      .catch(() => console.log(`error`))
      .finally(() => setIsFetching(false));
  }, []);
  return { lists: data, isFetching };
};
export const ArticleList = () => {
  const classes = useStyles();
  const { lists, isFetching } = useFetchArticleList();
  console.log(isFetching)
  // const lists: Article[] = [
  //   {
  //     id: 1,
  //     title: "aaaaaa",
  //     content: "aaaaaaaa"
  //   },
  //   {
  //     id: 1,
  //     title: "aaaaaa",
  //     content: "aaaaaaaa"
  //   },
  //   {
  //     id: 1,
  //     title: "aaaaaa",
  //     content: "aaaaaaaa"
  //   },
  //   {
  //     id: 1,
  //     title: "aaaaaa",
  //     content: "aaaaaaaa"
  //   },
  // ]
  const listItems = lists ? (
    lists.map((item) => (
      <ArticleCard title={item.title} content={item.content} id={item.id} />
    ))
  ) : (
    <>no items yet</>
  );

  return (
    <div className={classes.root}>
      { isFetching  ? (
        <div className={classes.loading}>
          <CircularProgress />
        </div>
      ) : (
        listItems
      )}
    </div>
  );
};
