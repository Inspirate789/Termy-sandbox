import React from "react";
import styled from "styled-components";
import {AuthRequest} from "../types";

const ModeButtonsWrapperStyled = styled.div`
  align-items: center;
  display: flex;
  flex-direction: column;
  gap: 10px;
  justify-content: center;
  padding: 0px 0px 0px 10px;
  position: relative;
  width: 149px;
`;

const TextWrapper = styled.button`
  border-width: 0;
  align-items: center;
  align-self: stretch;
  background-color: #00abe4;
  border-radius: 15px;
  display: flex;
  flex: 0 0 auto;
  flex-direction: column;
  gap: 10px;
  justify-content: center;
  overflow: hidden;
  padding: 10px 0px;
  position: relative;
  width: 100%;
`;

const Text = styled.div`
  color: #000000;
  font-family: "Inter-Regular", Helvetica;
  font-size: 20px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  text-align: center;
  width: fit-content;
`;

interface Props {
    mode: number
}

export const ModeButtonsWrapper = ({mode}: Props) => {
    if (mode === 2) {
        return (
            <ModeButtonsWrapperStyled>
                <TextWrapper>
                    <Text>
                        Разметка
                        <br />
                        текстов
                    </Text>
                </TextWrapper>
            </ModeButtonsWrapperStyled>
        );
    }

    return (
        <ModeButtonsWrapperStyled>
            <TextWrapper>
                <Text>
                    Сохранить
                    <br />
                    разметку
                </Text>
            </TextWrapper>
            <TextWrapper>
                <Text>
                    Анализ
                    <br />
                    терминов
                </Text>
            </TextWrapper>
        </ModeButtonsWrapperStyled>
    );
};
