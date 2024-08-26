import { Link } from 'react-router-dom';

import classes from './css/line-link.module.css';

type LineLinkProps = {
  to: string;
  span?: string;
  strong?: string;
};

export function LineLink(props: LineLinkProps) {
  const { to, span = '아이디가 없으신가요?', strong = '회원가입하기' } = props;

  return (
    <Link className={classes['line-like']} to={to}>
      <span>{span}</span>
      <strong>{strong}</strong>
    </Link>
  );
}
