import classes from './css/validation-text.module.css';

type ValidationTextProps = {
  condition: boolean;
  type: 'warning' | 'success';
  message: string;
};

export function ValidationText(props: ValidationTextProps) {
  const { type = 'warning', condition, message } = props;

  const className =
    type === 'warning'
      ? classes.paragraph + ' ' + classes.warning
      : classes.paragraph + ' ' + classes.success;

  return <>{condition && <p className={className}>{message}</p>}</>;
}
