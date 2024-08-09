import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import { GoogleLogin, useGoogleLogin } from '@react-oauth/google';

import { useValidationInput } from '@shared/hooks';
import { Button, Input, LineLink, ValidationText } from '@shared/ui';

import { emailRegex, passwordRegex } from '@features/auth/constants';

import classes from './css/login.module.css';
import { END_POINT, END_POINT_V2 } from '@shared/constants';
import useAuthStore from '@app/store/user';

import google from '@assets/icon/google.svg';

export default function Login() {
  const navigate = useNavigate();
  const setAccessToken = useAuthStore(state => state.setAccessToken);
  const setUser = useAuthStore(state => state.setUser);

  const { userValue: emailValue, getUserValue: getEmailValue } = useValidationInput(emailRegex);

  const { userValue: passwordValue, getUserValue: getPasswordValue } =
    useValidationInput(passwordRegex);

  const [isSuccess, setIsSuccess] = useState<boolean>(false);
  const [googleLoginError, setGoogleLoginError] = useState<string>('');

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
      console.error('로그인에 실패했습니다:', error);
      setIsSuccess(true);
    }
  }

  async function guestLoginHandler(): Promise<void> {
    try {
      const response = await fetch(END_POINT + '/auth/guest', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
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
    } catch (error) {
      if (error instanceof Error) {
        console.error(error.message);
      }
      console.error('게스트 로그인에 실패했습니다.');
    }
  }

  const googleLogin = useGoogleLogin({
    ux_mode: 'redirect',
    flow: 'auth-code',
    redirect_uri: 'https://food-recommendation.jokertrickster.com',
    onSuccess: async codeResponse => {
      try {
        // codeResponse가 제대로 전달되는지 확인
        if (!codeResponse || !codeResponse.code) {
          throw new Error('No code received from Google login');
        }

        console.log('Google Auth Response:', codeResponse);

        // code를 쿼리 파라미터로 전송
        const response = await fetch(
          `${END_POINT_V2}/auth/google/callback?code=${encodeURIComponent(codeResponse.code)}`,
          {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
            },
          }
        );

        const data = await response.json();
        console.log('Server Response:', data);

        if (data.accessToken) {
          setAccessToken(data.accessToken);
          setUser(data.accessToken);
          localStorage.setItem('accessToken', data.accessToken);
          localStorage.setItem('refreshToken', data.refreshToken);
          navigate('/chat'); // 로그인 성공 후 /chat 페이지로 이동
        } else {
          throw new Error('Access token not received');
        }
      } catch (error) {
        console.error('Error during Google login:', error);
        setGoogleLoginError(`Google 로그인 중 오류가 발생했습니다: ${(error as Error).message}`);
      }
    },
    onError: error => {
      console.error('Google Login Failed:', error);

      // 추가로 에러 정보를 더 상세하게 기록
      if (error.error) {
        console.error(`Error code: ${error.error}`);
      }

      setGoogleLoginError('Google 로그인에 실패했습니다.');
    },
  });

  const params = new URLSearchParams(window.location.search);
  const code = params.get('code');

  const handleLoginPost = async (code: string) => {
    console.log(code);

    try {
      const response = await fetch(END_POINT_V2 + `/auth/google/callback?code=${code}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const data = await response.json();
      console.log('Server Response:', data);

      if (data.accessToken) {
        setAccessToken(data.accessToken);
        setUser(data.accessToken);
        localStorage.setItem('accessToken', data.accessToken);
        localStorage.setItem('refreshToken', data.refreshToken);

        navigate('/chat'); // 로그인 성공 후 /chat 페이지로 이동
      } else {
        throw new Error('Access token not received');
      }
    } catch (error) {
      console.error('Error during Google login:', error);
      setGoogleLoginError(`Google 로그인 중 오류가 발생했습니다: ${(error as Error).message}`);
    }
  };

  useEffect(() => {
    if (code) {
      handleLoginPost(code);
    } else {
      console.log('로그인 재시도하세요.');
    }
  }, [code, navigate]);

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

      <div className={classes.buttons}>
        <Button className={classes.guest} type="button" onClick={googleLogin}>
          <img className={classes.google} src={google} alt="구글 로그인" />
          <span>구글 간편 로그인</span>
        </Button>
        <Button className={classes.guest} type="button" onClick={guestLoginHandler}>
          게스트 로그인
        </Button>
      </div>

      {googleLoginError && <p className={classes.error}>{googleLoginError}</p>}

      <LineLink to="/register" span="아이디가 없으신가요?" strong="회원가입하기" />
      <LineLink to="/password" span="" strong="암호를 잊어버리셨나요?" />
    </form>
  );
}
