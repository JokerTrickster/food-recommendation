import { memo } from 'react';

import classes from './css/card.module.css';

type CardProps = {
  children: React.ReactNode;
};

const Card = memo(function Card(props: CardProps) {
  const { children } = props;

  return <li className={classes.card}>{children}</li>;
});

export default Card;