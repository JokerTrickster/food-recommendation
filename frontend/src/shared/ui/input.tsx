import { InputHTMLAttributes, useState } from 'react';

import classes from './css/input.module.css';

type InputProps = {
  className?: string;
  label?: string;
  onValidation?: () => void;
} & InputHTMLAttributes<HTMLInputElement>;

export function Input(props: InputProps) {
  const { id, type, label, value, onValidation, className, ...rest } = props;

  const [isFocus, setIsFocus] = useState<boolean>(false);

  function focusHandler() {
    setIsFocus(true);
  }

  function blurHandler() {
    onValidation && onValidation();

    if (value!.toString().trim() === '') {
      setIsFocus(false);
    }
  }

  return (
    <div className={classes['input-container']} onFocus={focusHandler} onBlur={blurHandler}>
      <label htmlFor={id} className={isFocus ? classes.focus : classes['not-focus']}>
        {label}
      </label>
      <input id={id} type={type} className={`${classes.input} ${className}`} {...rest} />
    </div>
  );
}
