import { Box } from "@material-ui/core";
import { ArticleCard } from "./ArticleCard";

export interface Article {
  id: number;
  title: string;
  content: string;
}
export const ArticleList = () => {
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
  ]
  const listItems = lists.map((item) => <ArticleCard title={item.title} content={item.content} id={item.id}/>)
  
  return (
    <Box sx={{
      display: "flex",
      flex: 1
    }}>
    {listItems}
    </Box>
  )
}