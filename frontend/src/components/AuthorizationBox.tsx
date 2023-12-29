import React, {useState} from "react";
import styled from "styled-components";
import {AuthRequest} from "../types";

const StyledAuthorizationBox = styled.div`
  align-items: center;
  background-color: #e9f1fa;
  border-radius: 30px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: 162px;
  justify-content: center;
  overflow: hidden;
  padding: 10px;
  position: relative;
  width: 288px;
`;

const AuthorizationText = styled.div`
  color: #000000;
  font-family: "Inter-Regular", Helvetica;
  font-size: 24px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  width: fit-content;
`;

const AuthorizationForm = styled.div`
  align-items: center;
  display: flex;
  flex: 0 0 auto;
  gap: 10px;
  justify-content: center;
  padding: 10px;
  position: relative;
  width: 237px;
`;

const TextWrapper = styled.div`
  align-items: flex-end;
  display: inline-flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  justify-content: center;
  position: relative;
`;

const LoginText = styled.div`
  color: #000000;
  font-family: "Inter-Regular", Helvetica;
  font-size: 16px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  white-space: nowrap;
  width: fit-content;
`;

const PasswordText = styled(LoginText)`
  // Inherits styles from LoginText
`;

const InputsWrapper = styled.div`
  align-items: flex-start;
  align-self: stretch;
  display: flex;
  flex: 1;
  flex-direction: column;
  flex-grow: 1;
  gap: 10px;
  justify-content: center;
  position: relative;
`;

const Frame = styled.input`
  align-self: stretch;
  background-color: #ffffff;
  border-radius: 5px;
  flex: 1;
  flex-grow: 1;
  position: relative;
  width: 100%;
`;

const ButtonWrapper = styled.button`
  border-width: 0;
  align-items: center;
  background-color: #00abe4;
  border-radius: 7px;
  display: inline-flex;
  flex: 0 0 auto;
  gap: 10px;
  justify-content: center;
  overflow: hidden;
  padding: 5px 15px;
  position: relative;
`;

const ButtonText = styled.div`
  color: #000000;
  font-family: "Inter-Regular", Helvetica;
  font-size: 12px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  width: fit-content;
`;

interface Props {
    onSuccess: (user: AuthRequest) => void
}

export const AuthorizationBox = (props: Props) => {
    const [user, setUser] = useState<AuthRequest>({ username: '', password: '' })
    const { onSuccess } = props;

    const handleInput = (e: any) => {
        e.persist()
        const { id, value } = e.target;
        setUser(prev => ({ ...prev, [id]: value }))
    }

    const onSubmit = () => {
        return onSuccess(user);
    }

    return (
        <StyledAuthorizationBox>
            <AuthorizationText>Авторизация</AuthorizationText>
            <AuthorizationForm>
                <TextWrapper>
                    <LoginText>Логин:</LoginText>
                    <PasswordText>Пароль:</PasswordText>
                </TextWrapper>
                <InputsWrapper>
                    <Frame id="username" onChange={handleInput}/>
                    <Frame id="password" onChange={handleInput} type="password"/>
                </InputsWrapper>
            </AuthorizationForm>
            <ButtonWrapper>
                <ButtonText onClick={onSubmit}>Войти</ButtonText>
            </ButtonWrapper>
        </StyledAuthorizationBox>
    );
};
