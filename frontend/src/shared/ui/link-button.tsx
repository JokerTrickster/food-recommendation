import { memo } from 'react';
import { Link } from 'react-router-dom';

import classes from './css/link-button.module.css';

type LinkButtonProps = {
  to: string;
  className?: string;
  children: React.ReactNode;
};

const LinkButton = memo(function LinkButton(props: LinkButtonProps) {
  const { to, className = '', children } = props;

  return (
    <Link to={to} className={`${classes['link-button']} ${className}`}>
      {children}
    </Link>
  );
});

export default LinkButton;
