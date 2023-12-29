import PropTypes from "prop-types";
import React from "react";
import { ExitButton, InfoButton } from "./Buttons";
import styled from "styled-components";
import {InputUnit} from "../types";

const StyledTermWrapper = styled.div`
  align-items: flex-start;
  display: flex;
  flex-direction: column;
  position: relative;
  margin-left: 10px;

  & .term {
    align-items: flex-start;
    background-color: #ffffff;
    border-radius: 10px;
    display: flex;
    overflow: hidden;
    position: relative;
    width: 97.5%;
  }

  & .term-text {
    align-items: center;
    align-self: stretch;
    background-color: #ffffff;
    border-radius: 10px 0px 0px 10px;
    display: flex;
    flex: 1;
    flex-grow: 1;
    justify-content: flex-end;
    padding: 0px 0px 0px 10px;
    position: relative;
  }

  & .text {
    color: #000000;
    flex: 1;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
  }

  & .buttons-wrapper {
    align-items: flex-start;
    align-self: stretch;
    display: inline-flex;
    flex: 0 0 auto;
    gap: 10px;
    padding: 0px 0px 0px 10px;
    position: relative;
  }
`;

interface Props {
    className: any;
    unit: InputUnit;
    onClick: () => void
    onDelete: () => void
}

export const TermWrapper = ({className, unit, onClick, onDelete}: Props) => {
    return (
        <StyledTermWrapper className={`term-wrapper ${className}`}>
            <div className="term">
                <div className="term-text">
                    <div className="text">{unit.text}</div>
                </div>
                <div className="buttons-wrapper">
                    <InfoButton onClick={onClick}/>
                    <ExitButton onClick={onDelete}/>
                </div>
            </div>
        </StyledTermWrapper>
    );
};

TermWrapper.propTypes = {
    text: PropTypes.string,
    infoButtonInfoIcon: PropTypes.string,
    exitButtonDeleteIcon: PropTypes.string,
};
