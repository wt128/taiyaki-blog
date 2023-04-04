import { useContext } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import {
  AuthInfoContext,
  LoggedInContext,
} from '../Contexts/AuthContextProvider';
import { useAuth0 } from '@auth0/auth0-react';

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
  const isLoggedIn = useContext(LoggedInContext);
  const [authInfo, setAuthInfo] = useContext(AuthInfoContext);
  const { loginWithRedirect, isAuthenticated, logout } = useAuth0();
  const isShowAuth = () => {
    return !isAuthenticated ? (
      <Button onClick={() => loginWithRedirect()}>Log in</Button>
    ) : (
      <Button
        onClick={() =>
          logout({ logoutParams: { returnTo: window.location.origin } })
        }
      >
        Log Out
      </Button>
    );
  };
  return (
    <AppBar position="relative">
      <Toolbar>
        <IconButton
          edge="start"
          className={classes.menuButton}
          color="inherit"
          aria-label="menu"
        >
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
          News
        </Typography>
        <>{isShowAuth()}</>
      </Toolbar>
    </AppBar>
  );
};
