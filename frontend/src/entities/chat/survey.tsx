import { ChangeEvent, useState } from 'react';
import { createPortal } from 'react-dom';

import doughnut from '@assets/icon/doughnut.svg';

import RadioButton from '@shared/ui/radio-button';
import { Button } from '@shared/ui';

import { GENDER_CODE, Gender } from '@features/chat/constants';

import classes from './css/survey.module.css';
import { END_POINT } from '@shared/constants';

type SurveyProps = {
  onClose: () => void;
};

export function Survey(props: SurveyProps) {
  const { onClose } = props;

  const token = localStorage.getItem('accessToken');

  const [sexValue, setSexValue] = useState<string>('g1');

  const [bornValue, setBornValue] = useState<string>('');

  function bornHandler(e: ChangeEvent) {
    const target = e.target as HTMLInputElement;
    setBornValue(target.value);
  }

  function genderSelectedHandler(id: string) {
    setSexValue(id);
  }

  const isDisabled = !bornValue || !sexValue;

  async function surveySubmitHandler(e: React.FormEvent) {
    e.preventDefault();

    try {
      await fetch(END_POINT + '/users/profile', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          tkn: token!,
        },
        body: JSON.stringify({
          birth: bornValue,
          sex: sexValue,
        }),
      });

      onClose();
    } catch (error) {
      throw new Error('설문조사에 실패했습니다.');
    }
  }

  return createPortal(
    <>
      <div className={classes.overlay} />
      <section className={classes.survey}>
        <header>
          <p>
            생년월일과 성별을 입력해주시면, <br />
            <strong>추천</strong>에 큰 도움이돼요!
          </p>
        </header>

        <form className={classes.survey__form} onSubmit={surveySubmitHandler}>
          <section>
            <label htmlFor="born" className={classes.form__label}>
              생년월일
            </label>
            <input
              type="date"
              id="born"
              value={bornValue}
              onChange={bornHandler}
              className={bornValue ? classes.selected : classes.date}
            />
          </section>

          <section>
            <label htmlFor="gender" className={classes.form__label}>
              성별을 선택해주세요.
            </label>
            <ul>
              {GENDER_CODE.map((type: Gender) => {
                const Component: React.ReactNode = type.Component!;

                return (
                  <RadioButton
                    key={type.id}
                    title={type.type}
                    onChange={() => genderSelectedHandler(type.id)}
                    checked={sexValue === type.id}
                    {...type}
                  >
                    {Component}
                  </RadioButton>
                );
              })}
            </ul>
          </section>

          <div className={classes['form__button-container']}>
            <Button type="button" className={classes.form__cancel} onClick={onClose}>
              다음에 할래요!
            </Button>
            <Button
              className={isDisabled ? classes.disabled : classes.form__submit}
              disabled={isDisabled}
            >
              <img src={doughnut} alt="제출" />
              <span>제출</span>
            </Button>
          </div>
        </form>
      </section>
    </>,
    document.getElementById('modal-root')!
  );
}
