import classes from './css/radio-button.module.css';

type RadioButtonProps = {
  title: string;
  children?: React.ReactNode;
  onChange: (id: string) => void;
} & React.InputHTMLAttributes<HTMLInputElement>;

export default function RadioButton(props: RadioButtonProps) {
  const { name, id, title, children, checked, onChange } = props;

  return (
    <div className={classes['radio-container']}>
      <label className={checked ? classes.checked : ''} htmlFor={id}>
        {children}
        {title}
      </label>
      <input
        type="radio"
        name={name}
        id={id}
        checked={checked}
        onChange={() => onChange(id!)}
        value={id}
      />
    </div>
  );
}
