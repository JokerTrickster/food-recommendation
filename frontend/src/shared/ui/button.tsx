import { ButtonHTMLAttributes } from 'react';

import classes from './css/button.module.css';

type ButtonProps = {
  className?: string;
} & ButtonHTMLAttributes<HTMLButtonElement>;

export function Button(props: ButtonProps) {
  const { children, className } = props;

  return (
    <button className={`${classes.button}  ${className}`} {...props}>
      {children}
    </button>
  );
}
