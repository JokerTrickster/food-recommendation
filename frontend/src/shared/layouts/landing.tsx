import classes from './css/landing.module.css';
import { Outlet } from 'react-router';

export default function Landing() {
  return (
    <main className={classes.main}>
      <section className={classes.main__section}>
        <h1>고민은 이제 그만! </h1>
        <p>
          <span>뭐 먹을지 곤란할 땐</span> <strong>오늘 뭐먹지?</strong>
        </p>

        <Outlet />
      </section>
      <section className={classes.main__background}></section>
    </main>
  );
}
