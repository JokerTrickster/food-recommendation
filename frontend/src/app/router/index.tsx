import { createBrowserRouter } from 'react-router-dom';
import { AUTH_ROUTES } from './auth';
import { CHAT_ROUTES } from './chat';

export const router = createBrowserRouter([...AUTH_ROUTES, ...CHAT_ROUTES]);
