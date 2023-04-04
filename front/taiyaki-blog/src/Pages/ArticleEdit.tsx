import { Box, TextField, Typography, createStyles, makeStyles } from '@material-ui/core';
import { useState } from 'react';
import { ReactMarkdown } from 'react-markdown/lib/react-markdown';
const useStyles = makeStyles(() =>
  createStyles({
    root: {
      display: 'flex',
      margin: '1.5rem',
      width: '100%',
      height: '1200px',
      gap: '0 18px',
    },
    form: {
      width: '50%',
    },
    titleForm: {
      width: "calc(100% + 5px)",
      marginBottom: "4px"
    },
    textArea: {
      resize: 'none',
      width: '100%',
      height: '100%',
      fontSize: '24px',
      marginTop: "41px"
    },
    preview: {
      width: '50%',
      fontSize: '24px',
    },
  })
);
export const ArticleEdit = () => {
  const [content, setContent] = useState('');
  const [title, setTitle] = useState('');
  const classes = useStyles();
  return (
    <>
      <div className={classes.root}>
        <div className={classes.form}>
          <TextField className={classes.titleForm} id="outlined-basic" label="タイトル" variant="outlined" onChange={(e) => setTitle(e.target.value)}/>
          <textarea
            className={classes.textArea}
            onChange={(e) => setContent(e.target.value)}
          />
        </div>
        <div className={classes.preview}>
          <Typography variant="h2"></Typography>
          <ReactMarkdown>{content}</ReactMarkdown>
        </div>
      </div>
    </>
  );
};
