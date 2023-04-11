import { useNavigate } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Typography from '@material-ui/core/Typography';
import { Article } from './List';
const useStyles = makeStyles({
  root: {
    minWidth: 345,
    //width: "calc(100% / 3  - 40px)",
    flex: 1,
    boxSizing: 'border-box',
  },
  media: {
    height: 190,
  },
});

export const ArticleCard = ({ title, content, id }: Article) => {
  const classes = useStyles();
  const navigation = useNavigate();

  return (
    <Card className={classes.root}>
      <CardActionArea onClick={() => navigation(`/article/${id}`)}>
        <CardMedia
          className={classes.media}
          image="assets/react.svg"
          title="Contemplative Reptile"
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            {title}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            {content}
          </Typography>
        </CardContent>
      </CardActionArea>
    </Card>
  );
};
