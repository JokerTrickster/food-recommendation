import { Button, Input, LineLink, ValidationText } from '@shared/ui';
import { emailRegex, passwordRegex } from '@features/auth/constants';

import { useValidationInput } from '@shared/hooks';

import classes from './css/register.module.css';
import { END_POINT } from '@shared/constants';
import { useNavigate } from 'react-router';
import { useState } from 'react';

export default function Register(): JSX.Element {
  const [isCheck, setIsCheck] = useState<boolean>(true);

  const navigate = useNavigate();

  const {
    userValue: emailValue,
    getUserValue: getEmailValue,
    isValid: isEmailValid,
  } = useValidationInput(emailRegex);

  const {
    userValue: passwordValue,
    getUserValue: getPasswordValue,
    isValid: isPasswordValid,
  } = useValidationInput(passwordRegex);

  const {
    userValue: passwordValueCheck,
    getUserValue: getPasswordValueCheck,
    isValid: isPasswordValidCheck,
  } = useValidationInput(passwordRegex);

  async function registerHandler(e: React.FormEvent): Promise<void> {
    e.preventDefault();

    try {
      const response = await fetch(END_POINT + '/auth/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: emailValue,
          password: passwordValue,
        }),
      });

      if (response.ok) {
        navigate('/');
      }
    } catch (error) {
      console.error('회원가입 실패', error);
    }
  }

  async function duplicationHandler(e: React.MouseEvent): void {
    e.preventDefault();

    try {
      const response = await fetch(END_POINT + '/auth/check', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: emailValue,
        }),
      });

      if (response.ok) {
        setIsCheck(false);
      }
    } catch (error) {
      console.error('중복 확인 실패', error);
    }
  }

  return (
    <>
      <header className={classes.header}>
        <h1>회원가입</h1>
      </header>

      <form className={classes.register} onSubmit={registerHandler}>
        <div>
          <Input
            id="email"
            type="email"
            label="이메일"
            className={classes.login__input}
            value={emailValue}
            onChange={getEmailValue}
          />
          <Button className={classes.check} type="button" onClick={duplicationHandler}>
            중복 체크
          </Button>
        </div>

        <ValidationText
          condition={!isEmailValid && emailValue !== ''}
          type="warning"
          message="이메일이 유효하지 않습니다."
        />
        <Input
          id="password"
          type="password"
          label="패스워드"
          className={classes.login__input}
          value={passwordValue}
          onChange={getPasswordValue}
        />
        <Input
          id="password-check"
          type="password-check"
          label="패스워드"
          className={classes.login__input}
          value={passwordValueCheck}
          onChange={getPasswordValueCheck}
        />
        <ValidationText
          condition={!isPasswordValidCheck}
          type="warning"
          message="암호가 유효하지 않습니다."
        />

        <Button>회원가입</Button>
        <nav className={classes.nav}>
          <LineLink to="/" span="이미 회원이신가요?" strong="로그인하기" />
        </nav>
      </form>
    </>
  );
}
