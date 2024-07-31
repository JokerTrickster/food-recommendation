import { lazy } from 'react';

// import Sidebar from '@shared/layouts/sidebar';
const Chat = lazy(() => import('@pages/chat/chat'));

export const CHAT_ROUTES = [
  {
    path: '/chat',
    element: <Chat />,
  },
];
