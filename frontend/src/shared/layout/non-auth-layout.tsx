import avocado from '@assets/icon/avocado.svg';

import classes from './css/non-auth-layout.module.css';

type NonAuthLayoutProps = {
  children: React.ReactNode;
};

export default function NonAuthLayout(props: NonAuthLayoutProps) {
  const { children } = props;

  return (
    <main className={classes.main}>
      <div className={classes.main__null} />
      <section className={classes.main__section}>
        <article>
          <h1>
            뭐 먹을지 곤란할 땐? <br />
            오늘 뭐먹지.
          </h1>

          {children}
        </article>
      </section>

      <section className={classes['main__image-container']}>
        <img src={avocado} alt="avocado" />
      </section>
    </main>
  );
}
