import { useState } from 'react';

import classes from './css/radio-button.module.css';

type RadioButtonProps = {
  title: string;
} & React.InputHTMLAttributes<HTMLInputElement>;

export default function RadioButton(props: RadioButtonProps) {
  const { name, id, title, checked, onChange } = props;

  return (
    <div className={classes['radio-container']}>
      <label className={checked ? classes.checked : ''} htmlFor={id}>
        {title}
      </label>
      <input type="radio" name={name} id={id} checked={checked} onChange={() => onChange(id)} />
    </div>
  );
}
