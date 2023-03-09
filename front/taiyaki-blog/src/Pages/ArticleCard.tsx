import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import { Box } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { Article } from './ArticleList';
const useStyles = makeStyles({
  root: {
    minWidth: 345,
    //width: "calc(100% / 3  - 40px)",
    flex: 1,
    boxSizing:"border-box",
  },
  media: {
    height: 190,
  },
});

export const ArticleCard = ({title, content, id}: Article) => {
  const classes = useStyles();

  return (
    <Card className={classes.root}>
      <CardActionArea onClick={() => alert('aaaaa')}>
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
}
