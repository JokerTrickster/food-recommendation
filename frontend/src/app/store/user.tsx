import { create } from 'zustand';
import { jwtDecode } from 'jwt-decode';

interface JwtToken {
  createTime: number;
  email: string;
  exp: number;
  userId: number;
  sex?: string;
  birth?: string;
}

interface AuthState {
  accessToken: string | null;
  user: JwtToken;
  setAccessToken: (token: string | null) => void;
  logout: () => void;
  setUser: (token: string | null) => void;
}

const useAuthStore = create<AuthState>(set => ({
  accessToken: null,
  user: {
    createTime: 0,
    email: '',
    exp: 0,
    userId: 0,
  },
  setAccessToken: (token: string | null) => set({ accessToken: token }),
  logout: () => set({ accessToken: null }),
  setUser: (token: string | null) => {
    const jwtToken: JwtToken = jwtDecode(token!);
    set({ user: jwtToken });
  },
}));

export default useAuthStore;
