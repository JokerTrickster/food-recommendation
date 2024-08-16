import { useEffect } from 'react';
import { useNavigate } from 'react-router';

import { END_POINT_V2 } from '@shared/constants';
import useAuthStore from '@app/store/user';

export default function Google() {
  const navigate = useNavigate();

  // 현재 url에서 code 부분 추출
  const params = new URLSearchParams(window.location.search);
  const code = params.get('code');

  const setAccessToken = useAuthStore(state => state.setAccessToken);
  const setUser = useAuthStore(state => state.setUser);

  const handleLoginPost = async (code: string) => {
    try {
      const response = await fetch(END_POINT_V2 + `/auth/google/callback?code=${code}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const data = await response.json();

      setAccessToken(data.accessToken);
      setUser(data.accessToken);
      localStorage.setItem('accessToken', data.accessToken);
      localStorage.setItem('refreshToken', data.refreshToken);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (code) {
      handleLoginPost(code);
    } else {
      console.log('로그인 재시도하세요.');
    }
  }, [code, navigate]);
}
