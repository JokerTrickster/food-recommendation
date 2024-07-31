import { Button, Input, LineLink, ValidationText } from '@shared/ui';
import { emailRegex, passwordRegex } from '@features/auth/constants';

import { useValidationInput } from '@shared/hooks';

import classes from './css/register.module.css';
import { END_POINT } from '@shared/constants';
import { useNavigate } from 'react-router';
import { useState } from 'react';

export default function Register(): JSX.Element {
  const [isCheck, setIsCheck] = useState<boolean>(false);

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
      if (error instanceof Error) {
        console.error(error.message);
      }
    }
  }

  async function duplicationHandler(e: React.MouseEvent): Promise<void> {
    e.preventDefault();

    try {
      const response = await fetch(`${END_POINT}/auth/email/check?email=${emailValue}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.status === 400) {
        setIsCheck(true);
      } else {
        setIsCheck(false);
      }

      if (response.ok) {
        console.log(response.ok);
        setIsCheck(false);
      }
    } catch (error) {
      throw new Error('중복 체크에 실패했습니다.');
    }
  }

  const disabledButton = isEmailValid || isPasswordValid || isPasswordValidCheck;

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
          condition={!isCheck}
          type="success"
          message="사용할 수 있는 이메일입니다."
        />
        <ValidationText
          condition={!isEmailValid && emailValue !== ''}
          type="warning"
          message="이메일이 유효하지 않습니다."
        />
        <ValidationText condition={isCheck} type="warning" message="중복된 이메일입니다." />

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
          type="password"
          label="패스워드"
          className={classes.login__input}
          value={passwordValueCheck}
          onChange={getPasswordValueCheck}
        />

        <ValidationText
          condition={passwordValue.length > 0 && !isPasswordValid}
          type="warning"
          message="암호가 유효하지 않습니다."
        />
        <ValidationText
          condition={passwordValueCheck.length > 0 && passwordValue !== passwordValueCheck}
          type="warning"
          message="암호가 일치하지 않습니다."
        />

        <Button disabled={disabledButton}>회원가입</Button>
        <nav className={classes.nav}>
          <LineLink to="/" span="이미 회원이신가요?" strong="로그인하기" />
        </nav>
      </form>
    </>
  );
}
