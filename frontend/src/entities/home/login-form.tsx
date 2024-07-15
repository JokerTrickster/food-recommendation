import { Link } from 'react-router-dom';

import { useValidationInput } from '@shared/hooks';
import { Button, Input, ValidationText } from '@shared/ui';

import classes from './css/login-form.module.css';

export default function LoginForm() {
  const {
    userValue: emailValue,
    getUserValue: getEmailValue,
    isValid: isEmailValid,
    setIsValid: setIsEmailValid,
  } = useValidationInput();

  const {
    userValue: passwordValue,
    getUserValue: getPasswordValue,
    isValid: isPasswordValid,
    setIsValid: setIsPasswordValid,
  } = useValidationInput();

  function emailValidationHandler(): void {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

    if (emailRegex.test(emailValue)) {
      setIsEmailValid(false);
    } else {
      setIsEmailValid(true);
    }
  }

  function passwordValidationHandler() {
    if (passwordValue.trim() !== '') {
      setIsPasswordValid(false);
    } else {
      setIsPasswordValid(true);
    }
  }

  return (
    <form className={classes.login}>
      <section>
        <Input
          id="email"
          type="email"
          label="이메일"
          className={classes.login__input}
          value={emailValue}
          onChange={getEmailValue}
          onValidation={emailValidationHandler}
        />

        <ValidationText condition={isEmailValid} type="warning" message="이메일을 입력해주세요." />
      </section>

      {!isEmailValid && emailValue !== '' && (
        <section className={classes['container__password']}>
          <Input
            id="password"
            type="password"
            label="패스워드"
            className={classes.login__input}
            value={passwordValue}
            onChange={getPasswordValue}
            onValidation={passwordValidationHandler}
          />
          <ValidationText
            condition={isPasswordValid}
            type="warning"
            message="비밀번호를 입력해주세요."
          />
        </section>
      )}

      <Button type="submit">로그인</Button>

      <nav className={classes.nav}>
        <Link to="/signup">
          아이디가 없으신가요? <strong>회원가입하기</strong>
        </Link>
      </nav>
    </form>
  );
}
