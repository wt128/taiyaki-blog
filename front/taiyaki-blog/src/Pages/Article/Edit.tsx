import {
  Snackbar,
  TextField,
  Typography,
  createStyles,
  makeStyles,
} from '@material-ui/core';
import { ChangeEventHandler, FC, useEffect, useState } from 'react';
import { ReactMarkdown } from 'react-markdown/lib/react-markdown';
import { articleService } from '../../Features/api/article';
import { Alert, LoadingButton } from '@mui/lab';
import { useAuth0Token } from '../../Utils/auth';

const useStyles = makeStyles(() =>
  createStyles({
    head: {
      margin: '15px 40px',
      float: 'right',
    },
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
      width: 'calc(100% + 5px)',
      marginBottom: '4px',
    },
    textArea: {
      resize: 'none',
      width: '100%',
      height: '100%',
      fontSize: '24px',
      marginTop: '41px',
    },
    preview: {
      width: '50%',
      fontSize: '24px',
    },
  })
);

export const Edit: FC = () => {
  const [content, setContent] = useState('');
  const [title, setTitle] = useState('');
  
  const classes = useStyles();
  const token = useAuth0Token();
  const [loading, setLoading] = useState(false);
  console.log(token);
  document.title = '編集ページ';
  const [notifyState, setNotifyState] = useState({
    open: false,
    isSuccess: false,
  });
  const postArticle = () => {
    articleService
      .new(1, title, content, token)
      .then(() => {
        setNotifyState({ open: true, isSuccess: true });
      })
      .catch(() => {
        setNotifyState({ open: true, isSuccess: false });
      })
      .finally(() => setLoading(false));
  };
  const handlePost = () => {
    setLoading(true);
    postArticle();
  };

  const handleClose = () => {
    setNotifyState({ open: false, isSuccess: notifyState.isSuccess });
  };
  const handleTitleChange: ChangeEventHandler<HTMLTextAreaElement> = ({
    target,
  }) => {
    setTitle(target.value);
  };
  const notifyPosted = () => {
    return (
      <div>
        <Snackbar
          open={notifyState.open}
          autoHideDuration={6000}
          onClose={handleClose}
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'center',
          }}
        >
          <Alert
            onClose={handleClose}
            severity={notifyState.isSuccess ? 'success' : 'error'}
            sx={{ width: '100%' }}
          >
            {notifyState.isSuccess
              ? '投稿しました'
              : '通信中にエラーが発生しました'}
          </Alert>
        </Snackbar>
      </div>
    );
  };

  return (
    <>
      <div className={classes.head}>
        <LoadingButton
          onClick={handlePost}
          loading={loading}
          variant="contained"
          size="large"
        >
          投稿
        </LoadingButton>
      </div>
      <div className={classes.root}>
        {notifyPosted()}
        <div className={classes.form}>
          <TextField
            className={classes.titleForm}
            id="outlined-basic"
            label="タイトル"
            variant="outlined"
            onChange={handleTitleChange}
          />
          <textarea
            className={classes.textArea}
            onChange={(e) => setContent(e.target.value)}
          />
        </div>
        <div className={classes.preview}>
          <div>
            <Typography variant="h2">{title}</Typography>
          </div>
          <div>
            <ReactMarkdown>{content}</ReactMarkdown>
          </div>
        </div>
      </div>
    </>
  );
};
