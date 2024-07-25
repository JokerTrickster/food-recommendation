import send from '@assets/icon/send.svg';
import stop from '@assets/icon/stop.svg';

import { Button } from '@shared/ui';

import classes from './css/chat.module.css';
import RadioButton from '@shared/ui/radio-button';
import { useEffect, useState } from 'react';
import { Survey } from '@entities/chat/survey';
import { END_POINT } from '@shared/constants';
import useAuthStore from '@app/store/user';

export default function Chat() {
  const [selectedScenario, setSelectedScenario] = useState(')전체0');
  const [selectedTime, setSelectedTime] = useState('-전체0');
  const [selectedType, setSelectedType] = useState('+전체0');

  const [scenarios, setScenarios] = useState<string[]>([]);
  const [times, setTimes] = useState<string[]>([]);
  const [types, setTypes] = useState<string[]>([]);

  const [isSurvey, setIsSurvey] = useState(false);

  function surveyClose() {
    setIsSurvey(false);
  }

  const token = useAuthStore(state => state.accessToken);
  const user = useAuthStore(state => state.user);

  useEffect(() => {
    if (!token) {
      throw new Error('토큰이 없습니다');
    }
  }, [token]);

  useEffect(() => {
    (async function fetchData() {
      try {
        const response = await fetch(END_POINT + '/foods/meta');
        const data = await response.json();

        setScenarios(data.metaData.scenarios);
        setTimes(data.metaData.times);
        setTypes(data.metaData.types);
      } catch (error) {
        if (error instanceof Error) {
          throw new Error(error.message);
        }
      }
    })();
  }, []);

  useEffect(() => {
    if (!token) {
      throw new Error('토큰이 없습니다');
    }

    if (user.sex && user.birth) {
      return setIsSurvey(false);
    }

    setIsSurvey(true);
  }, [token]);

  console.log(token);

  return (
    <section className={classes.background}>
      {isSurvey && <Survey onClose={surveyClose} />}

      <section className={classes['chat__form-container']}>
        <form className={classes.chat__form}>
          <section>
            <strong>상황</strong>

            <div className={classes['radio-group']}>
              {scenarios &&
                scenarios.map((scenario, index) => {
                  return (
                    <RadioButton
                      key={')' + scenario + index}
                      id={')' + scenario + index}
                      name="scenario"
                      title={scenario}
                      checked={')' + scenario + index === selectedScenario}
                      onChange={() => setSelectedScenario(')' + scenario + index)}
                    />
                  );
                })}
            </div>
          </section>

          <section>
            <strong>시간</strong>

            <div className={classes['radio-group']}>
              {times &&
                times.map((time, index) => {
                  return (
                    <RadioButton
                      key={'-' + time + index}
                      id={'-' + time + index}
                      name="times"
                      title={time}
                      checked={'-' + time + index === selectedTime}
                      onChange={() => setSelectedTime('-' + time + index)}
                    />
                  );
                })}
            </div>
          </section>
          <section>
            <strong>카테고리</strong>

            <div className={classes['radio-group']}>
              {types &&
                types.map((type, index) => {
                  return (
                    <RadioButton
                      key={'+' + type + index}
                      id={'+' + type + index}
                      name="types"
                      title={type}
                      checked={'+' + type + index === selectedType}
                      onChange={() => setSelectedType('+' + type + index)}
                    />
                  );
                })}
            </div>
          </section>

          <div className={classes.form__wrap}>
            <Button className={classes.form__submit}>
              <img src={send} alt="메시지 전송" />
            </Button>
          </div>
        </form>
      </section>
    </section>
  );
}
