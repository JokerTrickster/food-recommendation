import { Outlet } from 'react-router';

import avocado from '@assets/icon/avocado.svg';

import classes from './css/non-auth-layout.module.css';
import { Suspense } from 'react';

export default function NonAuthLayout() {
  return (
    <main className={classes.main}>
      <div className={classes.main__null} />
      <section className={classes.main__section}>
        <article>
          <h1>
            뭐 먹을지 곤란할 땐? <br />
            오늘 뭐먹지.
          </h1>

          <Suspense fallback={<div>Loading...</div>}>
            <Outlet />
          </Suspense>
        </article>
      </section>

      <section className={classes['main__image-container']}>
        <img src={avocado} alt="avocado" />
      </section>
    </main>
  );
}
