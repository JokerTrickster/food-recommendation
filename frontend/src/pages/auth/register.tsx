import { Button, Input, LineLink } from '@shared/ui';

import classes from './css/register.module.css';
import Container from '@shared/ui/card';

export default function Register(): JSX.Element {
  return (
    <Container>
      <form className={classes.register}>
        <Input id="email" label="이메일" />
        <Input id="password" label="패스워드" />
        <Input id="password-check" label="패스워드 확인" />

        <Button>회원가입</Button>
        <nav className={classes.nav}>
          <LineLink to="/login" span="이미 회원이신가요?" strong="로그인하기" />
        </nav>
      </form>
    </Container>
  );
}
