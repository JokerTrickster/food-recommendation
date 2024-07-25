import { useNavigate } from 'react-router';

import { useValidationInput } from '@shared/hooks';
import { Button, Input, LineLink, ValidationText } from '@shared/ui';

import { emailRegex, passwordRegex } from '@features/auth/constants';

import classes from './css/login.module.css';
import { END_POINT } from '@shared/constants';
import useAuthStore from '@app/store/user';

export default function Login() {
  const navigate = useNavigate();
  const setAccessToken = useAuthStore(state => state.setAccessToken);
  const setUser = useAuthStore(state => state.setUser);

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
        setAccessToken(data.accessToken);
        setUser(data.accessToken);
        navigate('/chat');
      }
    } catch (error) {
      if (error instanceof Error) {
        console.error(error.message);
      }
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

        <ValidationText
          condition={!isEmailValid && emailValue !== ''}
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
        />
        <ValidationText
          condition={!isPasswordValid}
          type="warning"
          message="암호가 유효하지 않습니다."
        />
      </section>
      <Button type="submit">로그인</Button>
      <LineLink to="/register" span="아이디가 없으신가요?" strong="회원가입하기"></LineLink>
    </form>
  );
}
