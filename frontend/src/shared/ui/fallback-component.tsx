import classes from './css/fallback-component.module.css';

export default function FallbackComponent() {
  return (
    <main className={classes.fallback}>
      <h1 className={classes.loader__text}>페이지 이동 중 ...</h1>

      <div className={classes.loader}>
        <div className={classes.container}>
          <div className={classes.carousel}>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
            <div className={classes.love}></div>
          </div>
        </div>
        <div className={classes.container}>
          <div className={classes.carousel}>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
            <div className={classes.death}></div>
          </div>
        </div>
        <div className={classes.container}>
          <div className={classes.carousel}>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
            <div className={classes.robots}></div>
          </div>
        </div>
      </div>
    </main>
  );
}
