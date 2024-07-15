import { Link } from 'react-router-dom';
import { Button, Input, LineLink } from '@shared/ui';

import classes from './css/register.module.css';

export default function Register(): JSX.Element {
  return (
    <form className={classes.register}>
      <h2>회원 가입</h2>

      <Input id="email" label="이메일" />
      <Input id="password" label="패스워드" />
      <Input id="password-check" label="패스워드 확인" />

      <Button>회원가입</Button>
      <nav className={classes.nav}>
        <LineLink to="/">
          이미 회원이신가요? <strong>로그인 하기</strong>
        </LineLink>
      </nav>
    </form>
  );
}
