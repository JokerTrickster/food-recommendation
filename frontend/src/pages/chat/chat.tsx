import send from '@assets/icon/send.svg';

import { Button } from '@shared/ui';

import classes from './css/chat.module.css';
import RadioButton from '@shared/ui/radio-button';
import { useEffect, useState } from 'react';
import { Survey } from '@entities/chat/survey';
import { END_POINT } from '@shared/constants';
import useAuthStore from '@app/store/user';
import Card from '@shared/ui/card';
import { Link } from 'react-router-dom';
import { fetchToken } from '@features/chat/utils/token';

export default function Chat() {
  const [selectedScenario, setSelectedScenario] = useState('전체');
  const [selectedTime, setSelectedTime] = useState('전체');
  const [selectedType, setSelectedType] = useState('전체');

  const [accessToken, setAccessToken] = useState<string | null>(() =>
    localStorage.getItem('accessToken')
  );
  const [refreshToken, setRefreshToken] = useState<string | null>(() =>
    localStorage.getItem('refreshToken')
  );

  const [scenarios, setScenarios] = useState<string[]>([]);
  const [times, setTimes] = useState<string[]>([]);
  const [types, setTypes] = useState<string[]>([]);

  const [previousAnswer, setPreviousAnswer] = useState<string>('');
  const [answerList, setAnswerList] = useState<string[]>([]);

  const [isSurvey, setIsSurvey] = useState(false);

  const setUser = useAuthStore(state => state.setUser);

  function surveyClose() {
    setIsSurvey(false);
  }

  const token = useAuthStore(state => state.accessToken) || localStorage.getItem('accessToken');

  useEffect(() => {
    if (!token) {
      window.location.replace('/');
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
      token && window.location.replace('/');
    }

    setIsSurvey(true);
  }, [token]);

  async function submitHandler(e: React.FormEvent) {
    e.preventDefault();

    if (!accessToken || !refreshToken) {
      throw new Error('토큰이 없습니다');
    }

    console.log('previousAnswer', previousAnswer.replace(/,/g, ''));

    try {
      const response = await fetch(END_POINT + '/foods/recommend', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          tkn: accessToken!,
        },
        body: JSON.stringify({
          previousAnswer: previousAnswer.replace(/,/g, ''),
          scenario: selectedScenario,
          time: selectedTime,
          type: selectedType,
        }),
      });

      if (response.status === 401) {
        try {
          const newTokens = await fetchToken(accessToken, refreshToken);

          if (!newTokens.accessToken || !newTokens.refreshToken) {
            throw new Error('토큰이 정의되지 않았습니다.');
          }

          console.log('newTokens', newTokens);
          setAccessToken(newTokens.accessToken);
          setRefreshToken(newTokens.refreshToken);
          setUser(newTokens.accessToken);

          // 새 토큰으로 요청 재시도
          const newResponse = await fetch(END_POINT + '/foods/recommend', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              tkn: newTokens.accessToken,
            },
            body: JSON.stringify({
              previousAnswer,
              scenario: selectedScenario,
              time: selectedTime,
              type: selectedType,
            }),
          });

          if (newResponse.ok) {
            const data = await newResponse.json();
            setAnswerList(data.foodNames);
            setPreviousAnswer(() => data.foodNames.join(', '));
          }
        } catch (error) {
          console.error('토큰 갱신 실패:', error);
          throw new Error('토큰 갱신에 실패했습니다.');
        }
      }

      if (response.ok) {
        const data = await response.json();
        console.log(data);

        setAnswerList(data.foodNames);
        setPreviousAnswer(() => data.foodNames.join(', '));
      }
    } catch (error) {
      if (error instanceof Error) {
        console.error(error.message);
      }
    }
  }

  useEffect(() => {
    if (accessToken === null || refreshToken === null) {
      throw new Error('토큰이 없습니다');
    }

    (async function getProfile() {
      try {
        const response = await fetch(END_POINT + '/users/check', {
          method: 'GET',
          headers: {
            tkn: accessToken!,
          },
        });

        const data = await response.json();

        if (data.sex && data.birth) {
          console.log(data.sex && data.birth);
          setIsSurvey(false);
        }
      } catch (error) {
        if (error instanceof Error) {
          console.error(error.message);
        }
      }
    })();
  }, []);

  async function selectFoodHandler(food: string) {
    try {
      console.log(accessToken);

      const response = await fetch(END_POINT + '/foods/select', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          tkn: accessToken!,
        },
        body: JSON.stringify({
          name: food,
          scenario: selectedScenario,
          time: selectedTime,
          type: selectedType,
        }),
      });
      const data = await response.json();

      if (!response.ok) {
        console.error('Error:', data);
        throw new Error(data.message || 'Unknown error occurred');
      }

      console.log('Success:', data);
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <section className={classes.background}>
      {isSurvey && <Survey onClose={surveyClose} />}

      <main className={classes['answer-list']}>
        {answerList &&
          answerList.map((food, index) => (
            <Card key={'_' + index} onClick={() => selectFoodHandler(food)}>
              {food}
            </Card>
          ))}
      </main>

      <section className={classes['chat__form-container']}>
        <form className={classes.chat__form} onSubmit={submitHandler}>
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
                      value={scenario}
                      title={scenario}
                      checked={scenario === selectedScenario}
                      onChange={() => setSelectedScenario(scenario)}
                      onClick={() => {}}
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
                      value={time}
                      checked={time === selectedTime}
                      onChange={() => setSelectedTime(time)}
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
                      value={type}
                      title={type}
                      checked={type === selectedType}
                      onChange={() => setSelectedType(type)}
                    />
                  );
                })}
            </div>
          </section>

          <div className={classes.form__wrap}>
            <Button className={classes.form__submit}>
              <img src={send} alt="추천 받기" />
            </Button>
          </div>
        </form>
      </section>

      <Link to="/logout">로그아웃</Link>
    </section>
  );
}
