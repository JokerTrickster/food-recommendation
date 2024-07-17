import { ReactNode, memo } from 'react';

import classes from './css/card.module.css';

type CardProps = {
  children: ReactNode;
};

const Card = memo(function Card(props: CardProps) {
  const { children } = props;

  return <section className={classes.card}>{children}</section>;
});

export default Card;
