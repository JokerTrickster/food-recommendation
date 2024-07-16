import { Outlet, useLocation } from 'react-router';

import Message from '@shared/ui/message';
import LinkButton from '@shared/ui/link-button';
import { LineLink } from '@shared/ui';

import classes from './css/landing.module.css';

import { HIDE_AUTH_LAYOUT_PATH } from '@features/atuh/constants';

export default function Landing() {
  const location = useLocation();

  const hideLayout = !HIDE_AUTH_LAYOUT_PATH.includes(location.pathname);

  return (
    <main className={classes.main}>
      <header className={classes.main__header}>
        {hideLayout && (
          <>
            <h1>고민은 이제 그만! </h1>
            <p>
              <span>뭐 먹을지 곤란할 땐</span> <strong>오늘 뭐먹지?</strong>
            </p>

            <LinkButton to="/login">시작히기</LinkButton>

            <section>
              <LineLink to="register" />
            </section>
          </>
        )}
      </header>
      <Outlet />

      <section className={classes['main__chat-container']}>
        <article className={classes.container__chat}>
          <Message type="USER" message="혼밥 일식 오후 메뉴 추천해줘" />
          {/* 오른쪽에 있는 Message */}

          {/* 왼쪽에 있는 Message */}
          <Message type="Background" />

          <Message type="USER" message="돈까스" />
        </article>
      </section>
    </main>
  );
}
