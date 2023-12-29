import React from "react";
import styled from "styled-components";

interface Props {
    className?: string;
    lang: string
}

const StyledLangWrapper = styled.div`
  align-items: center;
  border-bottom-style: solid;
  border-bottom-width: 3px;
  border-color: #00abe4;
  border-radius: 15px;
  display: flex;
  gap: 10px;
  justify-content: center;
  overflow: hidden;
  padding: 10px 0px;
  position: relative;
  width: 774px;

  & > text {
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
`

export const LangWrapper = ({ className = "", lang }: Props) => {
    return (
        <StyledLangWrapper className={`lang-wrapper ${className}`}>
            <div className="text">{lang}</div>
        </StyledLangWrapper>
    );
}
