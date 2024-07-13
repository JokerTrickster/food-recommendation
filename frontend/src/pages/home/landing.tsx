import { Link } from 'react-router-dom';

import avocado from '@assets/icon/avocado.svg';

import classes from './css/landing.module.css';

export default function Landing() {
  return (
    <main className={classes.main}>
      <div className={classes.main__null} />
      <section className={classes.main__section}>
        <article>
          <h1>
            뭐 먹을지 곤란할 땐? <br />
            오늘 뭐먹지.
          </h1>

          <nav className={classes.section__nav}>
            <Link to="/login">시작하기</Link>
            <Link to="signup">
              아이디가 없으신가요? <strong>회원가입하기</strong>
            </Link>
          </nav>
        </article>
      </section>

      <section className={classes['main__image-container']}>
        <img src={avocado} alt="avocado" />
      </section>
    </main>
  );
}
