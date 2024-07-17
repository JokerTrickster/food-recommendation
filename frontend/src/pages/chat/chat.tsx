import send from '@assets/icon/send.svg';
import stop from '@assets/icon/stop.svg';

import { Button } from '@shared/ui';

import classes from './css/chat.module.css';
import RadioButton from '@shared/ui/radio-button';
import { useState } from 'react';
import { Survey } from '@entities/chat/survey';

const radioButtons = [
  { id: 'radio1', name: 'group1', title: 'Option 1' },
  { id: 'radio2', name: 'group1', title: 'Option 2' },
  { id: 'radio3', name: 'group1', title: 'Option 3' },
];

export default function Chat() {
  const [selectedId, setSelectedId] = useState('');

  const [isSurvey, setIsSurvey] = useState(true);

  const selectedGenderHandler = (id: string) => {
    setSelectedId(id);
  };

  function surveyClose() {
    setIsSurvey(false);
  }

  return (
    <section className={classes.background}>
      {isSurvey && <Survey onClose={surveyClose} />}
      <main className={classes['chat__message-container']}></main>

      <section className={classes['chat__form-container']}>
        <form className={classes.chat__form}>
          <section>
            <div className={classes['radio-group']}>
              {radioButtons.map(radio => (
                <RadioButton
                  key={radio.id}
                  id={radio.id}
                  name={radio.name}
                  title={radio.title}
                  checked={radio.id === selectedId}
                  onChange={() => selectedGenderHandler(radio.id)}
                />
              ))}
            </div>
          </section>

          <div className={classes.form__wrap}>
            <input type="text" />

            <Button className={classes.form__submit}>
              <img src={send} alt="메시지 전송" />
            </Button>
          </div>
        </form>
      </section>
    </section>
  );
}
