import { useState, ChangeEvent } from 'react';

export function useValidationInput() {
  const [userValue, setUserValue] = useState<string>('');
  const [isValid, setIsValid] = useState<boolean>(false);

  const getUserValue = (e: ChangeEvent<HTMLInputElement>): void =>
    setUserValue(e.currentTarget.value);

  return { userValue, getUserValue, isValid, setIsValid };
}
