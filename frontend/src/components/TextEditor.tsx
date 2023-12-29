import React from "react";
import styled from "styled-components";
import {LangWrapper} from "./LangWrapper";

const StyledTextEditor = styled.div`
  align-items: flex-start;
  background-color: #ffffff;
  border-radius: 15px;
  display: flex;
  gap: 10px;
  height: 742px;
  overflow: hidden;
  padding: 10px;
  position: relative;
  width: 773.5px;
`;

const StyledText = styled.p`
  align-self: stretch;
  color: transparent;
  flex: 1;
  font-family: "Inter-Regular", Helvetica;
  font-size: 20px;
  font-weight: 400;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
`;

const TextWrapper = styled.span`
  color: #e49196;
`;

const Span = styled.span`
  color: #000000;
`;

const TextEditor = () => {
    return (
        <StyledTextEditor>
            <StyledText>
                <TextWrapper>Отсутствие аккомодации</TextWrapper>
                <Span> свойственно не только явлению переноса количества движения, но и, как показывают </Span>
                <TextWrapper>экспериментальные исследования</TextWrapper>
                <Span>, процессу обмена энергией между </Span>
                <TextWrapper>падающими молекулами</TextWrapper>
                <Span> и стенкой.</Span>
            </StyledText>
        </StyledTextEditor>
    );
};

interface Props {
    lang: string
}

export const TextEditorWrapper = ({ lang }: Props) => {
    return (
        <div>
            <LangWrapper lang={lang}></LangWrapper>
            <TextEditor></TextEditor>
        </div>
    );
};
