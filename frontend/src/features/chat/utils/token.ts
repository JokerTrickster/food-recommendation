import { END_POINT } from '@shared/constants';

type Token = {
  accessToken: string;
  refreshToken: string;
};

export async function fetchToken(accessToken: string, refreshToken: string): Promise<Token> {
  if (!accessToken || !refreshToken) {
    throw new Error('토큰이 없습니다');
  }

  console.log('토큰재발급');
  try {
    const response = await fetch(END_POINT + '/auth/token/reissue', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        accessToken,
        refreshToken,
      }),
    });

    console.log(
      JSON.stringify({
        accessToken,
        refreshToken,
      })
    );

    if (!response.ok) {
      const errorText = await response.text();
      console.error('서버 오류 응답:', errorText);
      console.log(response);

      throw new Error(`서버 오류: ${response.status}, 응답: ${errorText} ${response}`);
    }

    const data = await response.json();

    return {
      accessToken: data.accessToken,
      refreshToken: data.refreshToken,
    };
  } catch (error) {
    if (error instanceof Error) {
      throw new Error(error.message);
    }
    throw new Error('알 수 없는 오류 발생');
  }
}
