import React from "react";
import styled from "styled-components";

interface Props {
    className: any;
    text: string;
}

const StyledInput = styled.div`
  align-items: center;
  background-color: #ffffff;
  border-radius: 10px;
  display: flex;
  height: 30px;
  overflow: hidden;
  padding: 0px 0px 0px 10px;
  position: relative;
  width: 1192px;

  & .text {
    color: black;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
    white-space: nowrap;
    width: fit-content;
  }
`;

export const Input = ({ className, text }: Props) => {
    return (
        <StyledInput className={`input ${className}`}>
            <div className="text">{text}</div>
        </StyledInput>
    );
};
