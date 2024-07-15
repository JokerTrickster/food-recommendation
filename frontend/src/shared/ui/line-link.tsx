import { ReactNode } from 'react';
import { Link } from 'react-router-dom';

import classes from './css/line-link.module.css';

type LineLinkProps = {
  to: string;
  children: ReactNode;
};

export function LineLink(props: LineLinkProps) {
  const { to, children } = props;

  return (
    <Link className={classes['line-like']} to={to}>
      {children}
    </Link>
  );
}
