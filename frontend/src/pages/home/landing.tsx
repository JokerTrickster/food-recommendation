import NonAuthLayout from '@shared/layout/non-auth-layout';
import LoginForm from '@entities/home/login-form';
import { Link } from 'react-router-dom';

export default function Landing() {
  return (
    <>
      <NonAuthLayout>
        <LoginForm />
      </NonAuthLayout>
    </>
  );
}
