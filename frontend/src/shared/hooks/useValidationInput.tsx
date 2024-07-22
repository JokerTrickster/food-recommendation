import { useState, ChangeEvent, useEffect } from 'react';

export function useValidationInput(regex: RegExp) {
  const [userValue, setUserValue] = useState<string>('');
  const [isValid, setIsValid] = useState<boolean>(false);

  function getUserValue(e: ChangeEvent<HTMLInputElement>): void {
    setUserValue(e.currentTarget.value);
  }

  useEffect(() => {
    if (userValue.trim() === '') {
      setIsValid(true);
    } else {
      setIsValid(regex.test(userValue));
    }
  }, [userValue, regex]);

  return {
    userValue,
    getUserValue,
    isValid,
    setIsValid,
  };
}
