import React, {useState} from "react";
import styled from "styled-components";
import {TermWrapper} from "./TermWrapper";
import {InputUnit} from "../types";

const TermWrapperWithInfoContainer = styled.div`
  align-items: flex-start;
  display: flex;
  flex-direction: column;
  padding: 3px 10px;
  position: relative;

  & .border-wrapper {
    align-items: center;
    align-self: stretch;
    border: 3px solid;
    border-color: #00abe4;
    border-radius: 15px;
    display: flex;
    flex: 0 0 auto;
    flex-direction: column;
    justify-content: center;
    margin-bottom: -3px;
    margin-left: -3px;
    margin-right: -3px;
    margin-top: -3px;
    padding: 0px 0px 5px 0px;
    position: relative;
    width: 98%;
  }

  & .info {
    align-items: flex-start;
    align-self: stretch;
    border-radius: 0px 0px 15px 15px;
    display: flex;
    flex: 0 0 auto;
    flex-direction: column;
    gap: 5px;
    padding: 0px 10px;
    position: relative;
    width: 100%;
    margin-top: 7px;
  }

  & .text-wrapper {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 14px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    margin-top: -3px;
    position: relative;
    width: fit-content;
  }

  & .div {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 14px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
    width: fit-content;
  }

  & .delete-button {
    left: 699px !important;
    position: absolute !important;
    top: 0 !important;
  }

  & .info-button-instance {
    left: 659px !important;
    position: absolute !important;
    top: 0 !important;
  }
`;

interface Props {
    className: any;
    unit: InputUnit;
    onDelete: () => void
}

export const TermWrapperWithInfo = ({className, unit, onDelete}: Props) => {
    const [opened, setOpened] = useState(false);

    const onOpen = () => {
        setOpened(!opened);
    }

    if (opened) {
        return (
            <TermWrapperWithInfoContainer className={`term-wrapper ${className}`}>
                <div className="border-wrapper">
                    <TermWrapper className={className} unit={unit} onClick={onOpen} onDelete={onDelete}/>
                    <div className="info">
                        <div className="text-wrapper">Структурная модель: {unit.model}</div>
                        <div className="div">Характеристики: {unit.properties?.join("; ")}</div>
                    </div>
                </div>
            </TermWrapperWithInfoContainer>
        );
    }

    return (
        <TermWrapper className={className} unit={unit} onClick={onOpen} onDelete={onDelete}/>
    );
};
