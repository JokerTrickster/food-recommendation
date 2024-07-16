import { lazy } from 'react';

const Chat = lazy(() => import('@pages/chat/chat'));

export const CHAT_ROUTES = [
  {
    path: '/chat',
    element: <Chat />,
  },
];
