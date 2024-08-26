import { lazy } from 'react';
import { GoogleOAuthProvider } from '@react-oauth/google';

import Landing from '@shared/layouts/landing';
import Login from '@pages/auth/login';
import Logout from '@pages/auth/logout';
import ForgotPassword from '@pages/auth/forgot-password';
import Google from '@pages/auth/google';

const Register = lazy(() => import('@pages/auth/register'));

const clientId = import.meta.env.VITE_GOOGLE_AUTH_CLIENT_ID;

export const AUTH_ROUTES = [
  {
    path: '/',
    element: <Landing />,
    children: [
      {
        index: true,
        element: (
          <GoogleOAuthProvider clientId={clientId}>
            <Login />
          </GoogleOAuthProvider>
        ),
      },

      {
        path: '/register',
        element: <Register />,
      },
      {
        path: 'password',
        element: <ForgotPassword />,
      },
    ],
  },
  {
    path: '/google',
    element: <Google />,
  },

  {
    path: '/logout',
    element: <Logout />,
  },
];
