import { useValidationInput } from '@shared/hooks';
import { Button, Input, LineLink, ValidationText } from '@shared/ui';
import Card from '@shared/ui/card';

import classes from './css/login.module.css';

export default function Login() {
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
    <Card>
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

          <ValidationText
            condition={isEmailValid && emailValue === ''}
            type="warning"
            message="이메일을 입력해주세요."
          />

          <ValidationText
            condition={isEmailValid && emailValue !== ''}
            type="warning"
            message="이메일이 유효하지 않습니다."
          />
        </section>
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
        <Button type="submit">로그인</Button>
        <LineLink to="/register" span="아이디가 없으신가요?" strong="회원가입하기"></LineLink>
      </form>
    </Card>
  );
}
