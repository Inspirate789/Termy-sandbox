import React from "react";
import styled from "styled-components";

interface Props {
    text: string;
}

const StyledText = styled.div`
  height: 30px;
  position: relative;
  width: 213px;

  & .text {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    left: 0;
    letter-spacing: 0;
    line-height: normal;
    position: absolute;
    top: -1px;
    width: 213px;
  }
`;

export const Text = ({text}: Props) => {
    return (
        <StyledText>
            <div className="text">{text}</div>
        </StyledText>
    );
};
