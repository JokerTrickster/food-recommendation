import { useValidationInput } from '@shared/hooks';
import { emailRegex, passwordRegex } from '@features/auth/constants';

import { Button, Input, ValidationText } from '@shared/ui';

import classes from './css/forgot-password.module.css';
import { END_POINT } from '@shared/constants';
import { useState } from 'react';
import { useNavigate } from 'react-router';

export default function ForgotPassword() {
  const navigate = useNavigate();

  const {
    userValue: emailValue,
    getUserValue: getEmailValue,
    isValid: isEmailValid,
  } = useValidationInput(emailRegex);

  const { userValue: emailAuthenticationValue, getUserValue: getEmailAuthenticationValue } =
    useValidationInput(emailRegex);

  const {
    userValue: passwordValue,
    getUserValue: getPasswordValue,
    isValid: isPasswordValid,
  } = useValidationInput(passwordRegex);

  const { userValue: passwordValueCheck, getUserValue: getPasswordValueCheck } =
    useValidationInput(passwordRegex);

  const [emailAuthentication, setEmailAuthentication] = useState<boolean>(false);

  async function emailCheckHandler(e: React.FormEvent): Promise<void> {
    e.preventDefault();

    try {
      const response = await fetch(END_POINT + '/auth/password/request', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: emailValue }),
      });

      if (response.ok) {
        setEmailAuthentication(true);
      }
      console.log(response.json());
    } catch (error) {
      throw new Error('이메일 전송에 실패했습니다.');
    }
  }

  async function forgotPasswordHandler(e: React.FormEvent): Promise<void> {
    e.preventDefault();

    console.log(
      'emailAuthentication: ',
      emailAuthentication,
      'emailValue: ',
      emailValue,
      'passwordValue: ',
      passwordValue
    );

    try {
      const response = await fetch(END_POINT + '/auth/password/validate', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          code: emailAuthenticationValue,
          email: emailValue,
          password: passwordValue,
        }),
      });

      if (response.ok) {
        navigate('/');
      } else {
        throw new Error('비밀번호 변경에 실패했습니다.');
      }
    } catch (error) {
      if (error instanceof Error) {
        console.error(error.message);
      }

      throw new Error('오류. 비밀번호 변경에 실패했습니다.');
    }
  }

  return (
    <>
      <header className={classes.header}>
        <h1>회원가입</h1>
      </header>

      <form className={classes['forgot-password']} onSubmit={forgotPasswordHandler}>
        <Input id="email" type="email" label="이메일" value={emailValue} onChange={getEmailValue} />
        <Button className={classes.check} type="button" onClick={emailCheckHandler}>
          이메일 확인
        </Button>

        <ValidationText
          condition={isEmailValid && emailValue === ''}
          type="warning"
          message="이메일이 유효하지 않습니다."
        />

        {emailAuthentication && (
          <Input
            id="authentication"
            type="text"
            label="인증번호"
            value={emailAuthenticationValue}
            onChange={getEmailAuthenticationValue}
          />
        )}

        <Input
          id="password"
          type="password"
          label="패스워드"
          value={passwordValue}
          onChange={getPasswordValue}
        />

        <Input
          id="password-check"
          type="password"
          label="패스워드"
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

        <Button>비밀번호 변경</Button>
      </form>
    </>
  );
}
