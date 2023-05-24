import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import { useAuth0 } from '@auth0/auth0-react';
import { Box } from '@mui/material';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  })
);
export const Header = () => {
  const classes = useStyles();
  const { loginWithRedirect, isAuthenticated, logout } = useAuth0();
  const isShowAuth = () => {
    return !isAuthenticated ? (
      <Button
        variant="contained"
        onClick={() =>
          loginWithRedirect({
            appState: { target: '<redirectUrl>', redirect_uri: '/article/new' },
          })
        }
      >
        Log in
      </Button>
    ) : (
      <>
        <Button
          variant="contained"
          onClick={() =>
            logout({ logoutParams: { returnTo: window.location.origin } })
          }
        >
          Log Out
        </Button>
        <Button
          variant="contained"
          onClick={() => (location.pathname = '/article/new')}
        >
          新規作成
        </Button>
      </>
    );
  };
  return (
    <AppBar position="relative">
      <Toolbar>
        {/* <IconButton
          edge="start"
          className={classes.menuButton}
          color="inherit"
          aria-label="menu"
        >
          <MenuIcon />
        </IconButton> */}
        <Typography
          variant="h6"
          className={classes.title}
          onClick={() => location.pathname = "/"}
        >
          Taiyaki
        </Typography>
        <Box sx={{ display: 'flex', gap: '6px' }}>
          <>{isShowAuth()}</>
        </Box>
      </Toolbar>
    </AppBar>
  );
};
