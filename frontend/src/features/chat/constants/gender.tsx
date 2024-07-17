import { ReactNode } from 'react';

import male from '@assets/icon/male.svg';
import female from '@assets/icon/female.svg';

import Male from '@assets/icon/male';
import Female from '@assets/icon/female';

export interface Gender {
  id: string;
  icon?: string;
  group: 'gender';
  type: '남성' | '여성' | '기타';
  value: number;
  Component?: ReactNode;
}

export const GENDER_CODE: Gender[] = [
  {
    id: 'g1',
    icon: male,
    type: '남성',
    value: 1,
    group: 'gender',
    Component: <Male />,
  },
  {
    id: 'g2',
    type: '여성',
    icon: female,
    value: 2,
    group: 'gender',
    Component: <Female />,
  },
  {
    id: 'g3',
    type: '기타',
    value: 0,
    group: 'gender',
  },
];
