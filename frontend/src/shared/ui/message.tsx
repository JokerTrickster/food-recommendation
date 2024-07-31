import { memo } from 'react';
import Checkbox from './checkbox';

import avocado from '@assets/icon/avocado.svg';

import classes from './css/message.module.css';
interface GptMessage {
  id: number;
  message: string;
}

type UserMessageProps = {
  type: 'USER';
  message: string;
  selectList?: never;
};

type GptMessageProps = {
  type: 'GPT';
  message?: never;
  selectList?: GptMessage[];
};

type BackGroundProps = {
  type: 'Background';
  message?: never;
  selectList?: never;
};
type MessageProps = UserMessageProps | GptMessageProps | BackGroundProps;

const Message = memo(function message(props: MessageProps) {
  const { type = 'USER', selectList, message } = props;

  if (type === 'Background') {
    return (
      <div className={classes.gpt}>
        <div className={classes['gpt__profile-container']}>
          <img src={avocado} alt="GPT 프로필 사진 아보카도" />
        </div>
        <ul>
          <div className={classes.gpt__container}>
            <li style={{ width: '150px' }}>
              <p className={classes['container__gpt-message']}>
                <span>초밥</span>
                <Checkbox disabled />
              </p>
            </li>

            <li style={{ width: '150px' }}>
              <p className={classes['container__gpt-message']}>
                <span>우동</span>
                <Checkbox disabled />
              </p>
            </li>

            <li style={{ width: '150px' }}>
              <p className={classes['container__gpt-message']}>
                <span>돈까스</span>
                <Checkbox checked={true} disabled />
              </p>
            </li>

            <li>
              <p className={classes.container__recommend}>
                추천 받은 메뉴를 선택해주시면, 추천에 큰 도움이 됩니다!
              </p>
            </li>
          </div>
        </ul>
      </div>
    );
  }

  if (type === 'GPT') {
    return (
      <div className={classes.gpt}>
        <div className={classes['gpt__profile-container']}>
          <img src={avocado} alt="GPT 프로필 사진 아보카도" />
        </div>
        <ul>
          {selectList &&
            selectList.map(
              (item: GptMessage, index: number): JSX.Element => (
                <div className={classes.gpt__container} key={'g' + item.id}>
                  <li style={{ width: `${item.message.length * 10 + 100}px` }}>
                    <p className={classes['container__gpt-message']}>
                      <span>{item.message}</span>
                      <Checkbox />
                    </p>
                  </li>
                  {index === selectList.length - 1 && (
                    <li>
                      <p className={classes.container__recommend}>
                        추천 받은 메뉴를 선택해주시면, 추천에 큰 도움이 됩니다!
                      </p>
                    </li>
                  )}
                </div>
              )
            )}
        </ul>
      </div>
    );
  }

  return (
    <section className={classes['user-container']}>
      <p className={classes['container__user-message']}>{message}</p>
    </section>
  );
});

export default Message;
