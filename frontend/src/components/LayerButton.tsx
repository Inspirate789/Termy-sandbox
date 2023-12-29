import React from "react";
import styled from "styled-components";

interface Props {
    text?: string;
    background: string;
    selected: boolean;
    setCurrentLayer: (layer: string) => void;
}

const StyledLayerButton = styled.button`
  align-items: flex-start;
  border-color: black;
  border-radius: 0px 15px 15px 0px;
  display: flex;
  gap: 10px;
  justify-content: center;
  padding: 10px;
  position: relative;
  width: 149px;

  .layer-name {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    margin-top: -1px;
    position: relative;
    white-space: nowrap;
    width: fit-content;
  }
`;

export const LayerButton = ({text = "", background, selected, setCurrentLayer}: Props) => {
    const onClick = () => {
        return setCurrentLayer(text);
    }

    return (
        <StyledLayerButton style={{background: background, borderWidth: selected ? 2 : 0}} onClick={onClick}>
            <div className="layer-name">{text}</div>
        </StyledLayerButton>
    );
};
