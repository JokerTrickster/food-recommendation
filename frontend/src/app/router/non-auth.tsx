import { lazy } from 'react';

import Landing from '@pages/home/landing';
import NonAuthLayout from '@shared/layout/non-auth-layout';

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
        element: lazy(() => import('@pages/auth/register')),
      },
    ],
  },
];
