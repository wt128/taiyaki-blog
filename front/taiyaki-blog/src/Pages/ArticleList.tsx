import { Box, createStyles, makeStyles, Theme } from "@material-ui/core";
import { ArticleCard } from "./ArticleCard";

export interface Article {
  id: number;
  title: string;
  content: string;
}
const useStyles = makeStyles((theme: Theme) => 
  createStyles({
    root: {
      display: "flex",
      margin: "1.5rem",
      flexWrap: "wrap",
      gap: "10px 1%",
      width: "calc(100% - 40px)",
      justifyContent:"space-between"
    }
  })
)
export const ArticleList = () => {
  const classes = useStyles();
  const lists: Article[] = [
    {
      id: 1,
      title: "aaaaaa",
      content: "aaaaaaaa"
    },
    {
      id: 1,
      title: "aaaaaa",
      content: "aaaaaaaa"
    },
    {
      id: 1,
      title: "aaaaaa",
      content: "aaaaaaaa"
    },
    {
      id: 1,
      title: "aaaaaa",
      content: "aaaaaaaa"
    },
  ]
  const listItems = lists.map((item) => <ArticleCard title={item.title} content={item.content} id={item.id}/>)
  
  return (
    <div className={classes.root}>
      {listItems}
    </div>
  )
}