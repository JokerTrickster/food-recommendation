import { memo } from 'react';

import classes from './css/card.module.css';

type CardProps = {
  children: React.ReactNode;
  onClick?: () => void;
};

const Card = memo(function Card(props: CardProps) {
  const { children, onClick } = props;

  return (
    <li className={classes.card} onClick={onClick}>
      {children}
    </li>
  );
});

export default Card;
