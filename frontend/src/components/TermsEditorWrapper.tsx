import React from "react";
import { ExitButton } from "./Buttons";
import { Input } from "./Input";
import { Text } from "./Text";
import styled from "styled-components";

const StyledTermEditor = styled.div`
  align-items: flex-start;
  border: 3px solid;
  border-color: #00abe4;
  border-radius: 15px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  justify-content: center;
  overflow: hidden;
  padding: 0px 45px 10px 0px;
  position: relative;
  width: 1490px;

  & .term-wrapper {
    align-items: flex-start;
    align-self: stretch;
    display: flex;
    flex: 0 0 auto;
    gap: 10px;
    position: relative;
    width: 100%;
  }

  & .term-number {
    align-items: center;
    border: 3px solid;
    border-color: #00abe4;
    border-radius: 15px;
    display: inline-flex;
    flex: 0 0 auto;
    flex-direction: column;
    overflow: hidden;
    padding: 5px 10px;
    position: relative;
  }

  & .text-2 {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    margin-top: -3px;
    position: relative;
    white-space: nowrap;
    width: fit-content;
  }

  & .term-input-wrapper {
    align-items: flex-start;
    border-radius: 10px;
    display: flex;
    flex: 1;
    flex-grow: 1;
    overflow: hidden;
    padding: 10px 0px 0px;
    position: relative;
  }

  & .term-input {
    align-items: center;
    background-color: #ffffff;
    border-radius: 10px;
    display: flex;
    flex: 1;
    flex-grow: 1;
    height: 30px;
    justify-content: flex-end;
    overflow: hidden;
    padding: 0px 0px 0px 10px;
    position: relative;
  }

  & .text-3 {
    color: #000000;
    flex: 1;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
  }

  & .term-info {
    align-items: flex-start;
    align-self: stretch;
    background-color: #e9f1fa;
    display: flex;
    flex: 0 0 auto;
    gap: 10px;
    justify-content: center;
    position: relative;
    width: 100%;
  }

  & .input-titles {
    align-items: flex-start;
    display: inline-flex;
    flex: 0 0 auto;
    flex-direction: column;
    gap: 10px;
    justify-content: center;
    padding: 0px 10px;
    position: relative;
  }

  & .input-wrapper {
    align-items: flex-start;
    display: flex;
    flex: 1;
    flex-direction: column;
    flex-grow: 1;
    gap: 10px;
    justify-content: center;
    position: relative;
  }

  & .input-instance {
    align-self: stretch !important;
    width: 100% !important;
  }

  & .delete-button {
    left: 1460px !important;
    position: absolute !important;
    top: 0 !important;
  }
`;

const StyledTermsEditor = styled.div`
  align-items: center;
  display: flex;
  flex-direction: column;
  gap: 10px;
  height: 593px;
  padding: 15px;
  position: relative;
  width: 497px;

  & .term-editor-instance {
    align-self: stretch !important;
    flex: 0 0 auto !important;
    width: 100% !important;
  }

  & .design-component-instance-node {
    left: 437px !important;
    position: absolute !important;
    top: 0 !important;
  }
`;

const StyledTermsEditorWrapper = styled.div`
  align-items: center;
  border-bottom-style: solid;
  border-bottom-width: 3px;
  border-color: #00abe4;
  border-left-style: solid;
  border-left-width: 3px;
  border-right-style: solid;
  border-right-width: 3px;
  display: flex;
  flex-direction: column;
  height: 632px;
  position: relative;
  width: 506.67px;

  & .lang-wrapper {
    align-items: center;
    align-self: stretch;
    border-bottom-style: solid;
    border-bottom-width: 3px;
    border-color: #00abe4;
    border-left-style: solid;
    border-left-width: 3px;
    border-right-style: solid;
    border-right-width: 3px;
    display: flex;
    flex: 0 0 auto;
    gap: 10px;
    justify-content: center;
    padding: 10px 0px;
    position: relative;
    width: 100%;

    & .lang-text {
      color: #000000;
      font-family: "Inter-Regular", Helvetica;
      font-size: 20px;
      font-weight: 400;
      letter-spacing: 0;
      line-height: normal;
      margin-top: -3px;
      position: relative;
      white-space: nowrap;
      width: fit-content;
    }
  }

  & .terms-editor-instance {
    align-self: stretch !important;
    flex: 1 !important;
    flex-grow: 1 !important;
    height: unset !important;
    width: 100% !important;
  }
`;


const TermEditor = () => {
    return (
        <StyledTermEditor>
            <div className="term-wrapper">
                <div className="term-number">
                    <div className="text-2">1.</div>
                </div>
                <div className="term-input-wrapper">
                    <div className="term-input">
                        <div className="text-3">Термин 1</div>
                    </div>
                </div>
            </div>
            <div className="term-info">
                <div className="input-titles">
                    <Text text="Структурная модель:" />
                    <Text text="Характеристики:" />
                    <Text text="Перевод:" />
                </div>
                <div className="input-wrapper">
                    <Input className="input-instance" text="Части речи через ‘+’" />
                    <Input className="input-instance" text="Характеристики через ‘;’" />
                    <Input className="input-instance" text="Номер соответствующего термина" />
                </div>
            </div>
            <div className="delete-button"><ExitButton/></div>
        </StyledTermEditor>
    );
};

export const TermsEditor = () => {
    return (
        <StyledTermsEditor>
            <div className="term-editor-instance"><TermEditor/></div>
            <div className="term-editor-instance"><TermEditor/></div>
            <div className="term-editor-instance"><TermEditor/></div>
        </StyledTermsEditor>
    );
};

export const TermsEditorWrapper = () => {
    return (
        <StyledTermsEditorWrapper>
            <div className="lang-wrapper">
                <div className="lang-text">Русский</div>
            </div>
            <div className="terms-editor-instance"><TermsEditor/></div>
        </StyledTermsEditorWrapper>
    );
};
