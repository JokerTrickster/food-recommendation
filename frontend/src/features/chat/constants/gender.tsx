import male from '@assets/icon/male.svg';
import female from '@assets/icon/female.svg';

import Male from '@assets/icon/male';
import Female from '@assets/icon/female';

export interface Gender {
  id: string;
  icon?: string;
  group: 'sex';
  type: '남성' | '여성' | '기타';
  value: number;
  Component?: React.ReactNode;
}

export const GENDER_CODE: Gender[] = [
  {
    id: 'male',
    icon: male,
    type: '남성',
    value: 1,
    group: 'sex',
    Component: <Male />,
  },
  {
    id: 'female',
    type: '여성',
    icon: female,
    value: 2,
    group: 'sex',
    Component: <Female />,
  },
  {
    id: 'outer',
    type: '기타',
    value: 0,
    group: 'sex',
  },
];
