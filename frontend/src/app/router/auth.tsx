import { lazy } from 'react';

import Landing from '@pages/home/landing';
import Layout from '@shared/layout/layout';
const Login = lazy(() => import('@pages/auth/login'));
const Register = lazy(() => import('@pages/auth/register'));

export const AUTH_ROUTES = [
  {
    path: '/',
    element: <Landing />,
    children: [
      {
        path: '/login',
        element: <Login />,
      },

      {
        path: '/register',
        element: <Register />,
      },
    ],
  },
];
