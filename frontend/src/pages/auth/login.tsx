import { useNavigate } from 'react-router';

import { useValidationInput } from '@shared/hooks';
import { Button, Input, LineLink, ValidationText } from '@shared/ui';

import { emailRegex, passwordRegex } from '@features/auth/constants';

import classes from './css/login.module.css';
import { END_POINT } from '@shared/constants';
import useAuthStore from '@app/store/user';
import { useState } from 'react';

export default function Login() {
  const navigate = useNavigate();
  const setAccessToken = useAuthStore(state => state.setAccessToken);
  const setUser = useAuthStore(state => state.setUser);

  const { userValue: emailValue, getUserValue: getEmailValue } = useValidationInput(emailRegex);

  const { userValue: passwordValue, getUserValue: getPasswordValue } =
    useValidationInput(passwordRegex);

  const [isSuccess, setIsSuccess] = useState<boolean>(false);

  async function loginHandler(e: React.FormEvent): Promise<void> {
    e.preventDefault();

    try {
      const response = await fetch(END_POINT + '/auth/signin', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: emailValue, password: passwordValue }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log(data);

        setAccessToken(data.accessToken);
        setUser(data.accessToken);
        localStorage.setItem('accessToken', data.accessToken);
        localStorage.setItem('refreshToken', data.refreshToken);
        navigate('/chat');
      }

      if (response.status === 400) {
        setIsSuccess(true);
      } else {
        setIsSuccess(false);
      }
    } catch (error) {
      throw new Error('로그인에 실패했습니다.');
    }
  }

  return (
    <form className={classes.login} onSubmit={loginHandler}>
      <section>
        <Input
          id="email"
          type="email"
          label="이메일"
          className={classes.login__input}
          value={emailValue}
          onChange={getEmailValue}
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
        />
      </section>

      <ValidationText
        condition={isSuccess}
        type="warning"
        message="이메일 또는 암호가 유효하지 않습니다."
      />

      <Button type="submit" disabled={isSuccess}>
        로그인
      </Button>
      <LineLink to="/register" span="아이디가 없으신가요?" strong="회원가입하기"></LineLink>
      <LineLink to="/register" span="" strong="암호를 잊어버리셨나요?"></LineLink>
    </form>
  );
}
