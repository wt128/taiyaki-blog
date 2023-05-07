import { Box, Typography } from '@mui/material';
import AccessAlarmIcon from '@mui/icons-material/AccessAlarm';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import moment from 'moment';
export const Author: React.FC<{ author: string; createdAt: string }> = ({
  author,
  createdAt,
}) => {
  return (
    <>
      <Box
        sx={{
          display: 'flex',
          justifyContent: 'center',
          mt: '1rem',
          gap: '9px',
        }}
      >
        <AccountCircleIcon />
        <Typography variant="body1" fontWeight="light">
          {author}
        </Typography>
        <AccessAlarmIcon />
        <Typography variant="body1" fontWeight="light">
          {moment(createdAt).format('YYYY年MM月DD日')}
        </Typography>
      </Box>
    </>
  );
};
