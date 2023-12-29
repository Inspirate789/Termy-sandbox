import React from "react";
import styled from "styled-components";

interface Props {
    className?: string;
}

const StyledTitleWrapper = styled.div`
  align-items: flex-end;
  display: flex;
  overflow: hidden;
  position: relative;

  .title {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 36px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    margin-top: -1px;
    position: relative;
    width: fit-content;
  }

  .version {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 12px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
    width: fit-content;
  }
`;

export const TitleWrapper = ({ className = "" }: Props) => {
    return (
        <StyledTitleWrapper className={className}>
            <div className="title">Termy</div>
            <div className="version">v.0.1.0</div>
        </StyledTitleWrapper>
    );
};
