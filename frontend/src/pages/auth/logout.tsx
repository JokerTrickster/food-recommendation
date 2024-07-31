import { END_POINT } from '@shared/constants';
import { useEffect } from 'react';

export default function Logout() {
  const token = localStorage.getItem('accessToken');

  useEffect(() => {
    (async function logoutHandler() {
      if (!token) {
        throw new Error('토큰이 없습니다');
      }

      try {
        await fetch(END_POINT + '/auth/logout', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            tkn: token,
          },
        });

        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
        window.location.replace('/');
      } catch (error) {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
        window.location.replace('/');
      }
    })();
  }, []);

  return <></>;
}
