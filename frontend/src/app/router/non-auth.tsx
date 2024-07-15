import { lazy } from 'react';

import Landing from '@pages/home/landing';
import NonAuthLayout from '@shared/layout/non-auth-layout';

const Register = lazy(() => import('@pages/auth/register'));

export const NON_AUTH = [
  {
    path: '/',
    element: <NonAuthLayout />,
    children: [
      {
        index: true,
        element: <Landing />,
      },
      {
        path: '/register',
        element: <Register />,
      },
    ],
  },
];
