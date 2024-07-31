import { lazy } from 'react';

import Landing from '@shared/layouts/landing';
import Login from '@pages/auth/login';
import Logout from '@pages/auth/logout';

const Register = lazy(() => import('@pages/auth/register'));

export const AUTH_ROUTES = [
  {
    path: '/',
    element: <Landing />,
    children: [
      {
        index: true,
        element: <Login />,
      },

      {
        path: '/register',
        element: <Register />,
      },
    ],
  },

  {
    path: '/logout',
    element: <Logout />,
  },
];
