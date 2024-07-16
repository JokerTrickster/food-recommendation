import { lazy } from 'react';

import Sidebar from '@shared/layouts/sidebar';
const Chat = lazy(() => import('@pages/chat/chat'));

export const CHAT_ROUTES = [
  {
    path: '/chat',
    element: <Sidebar />,
    children: [
      {
        path: '/chat',
        element: <Chat />,
      },
      {
        path: '/chat/new',
        element: <Chat />,
      },
      {
        path: '/chat:roomId',
        element: <Chat />,
      },
    ],
  },
];
