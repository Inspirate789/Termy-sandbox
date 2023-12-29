import React from "react";
import styled from "styled-components";

const StyledButton = styled.button`
  border-width: 0;
  align-items: flex-start;
  background-color: #00abe4;
  border-radius: 10px;
  display: inline-flex;
  gap: 10px;
  overflow: hidden;
  padding: 5px;
  position: relative;
`;

const Icon = styled.img`
  height: 20px;
  object-fit: cover;
  position: relative;
  width: 20px;
`;

interface Props {
    onClick?: () => void
}

export const ExitButton = ({onClick}: Props) => {
    return (
        <StyledButton onClick={onClick}>
            <Icon src="assets/delete-icon.png" />
        </StyledButton>
    );
};

export const InfoButton = ({onClick}: Props) => {
    return (
        <StyledButton onClick={onClick}>
            <Icon src="assets/info-icon.png" />
        </StyledButton>
    );
};
