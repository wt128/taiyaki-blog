import {
  Chip,
  styled,
  PropTypes
} from '@material-ui/core';

interface ITag {
  name: string;
  color: Exclude<PropTypes.Color, 'inherit'>;
  key: number;
}

const ListItem = styled('li')(({ theme }) => ({
  margin: theme.spacing(0.5),
}));

export const Tag: React.FC<ITag[]> = (tags: ITag[]) => {
  return (
    <>
      {tags.map((t, key) => {
        return (
          <ListItem key={key}>
            <Chip label={`#${t.name}`} color={t.color} />
          </ListItem>
        );
      })}
    </>
  );
};
