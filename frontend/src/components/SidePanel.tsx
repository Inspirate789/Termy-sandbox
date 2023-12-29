import React, {useEffect, useState} from "react";
import styled from "styled-components";
import { LayerButton } from "./LayerButton";
import { ModeButtonsWrapper } from "./ModeButtonsWrapper";
import layersService from "../services/LayersService";

const StyledSidePanel = styled.div`
  align-items: flex-start;
  display: inline-flex;
  flex-direction: column;
  gap: 10px;
  height: 786px;
  position: relative;
`;

const ButtonsWrapper = styled.div`
  align-items: flex-start;
  align-self: stretch;
  display: flex;
  flex: 1;
  flex-direction: column;
  flex-grow: 1;
  gap: 10px;
  position: relative;
  width: 100%;    
`;

const DesignComponentInstanceNode = styled(ModeButtonsWrapper)`
  align-self: stretch !important;
  flex: 0 0 auto !important;
  width: 100% !important;
`;

const LayerButtonWrap = styled(LayerButton)`
  background-color: {$background}; !important;
  border: 5px brown;
  flex: 0 0 auto !important;
`;

// const LayerButton2 = styled(LayerButton)`
//   background-color: #e49196 !important;
//   flex: 0 0 auto !important;
// `;
//
// const LayerButton3 = styled(LayerButton)`
//   background-color: #ff9d57 !important;
//   flex: 0 0 auto !important;
// `;
//
// const LayerButton4 = styled(LayerButton)`
//   background-color: #97f9a1 !important;
//   flex: 0 0 auto !important;
// `;

const LayerButtonAdd = styled(LayerButton)`
  background-color: #dbdbdb !important;
  flex: 0 0 auto !important;
`;

interface Props {
    currentLayer: string
    setCurrentLayer: (layer: string) => void;
}

export const SidePanel = ({currentLayer, setCurrentLayer}: Props) => {
    const [loading, setLoading] = useState(false)
    const [layers, setLayers] = useState<string[]>([]);
    const colors = ["#7bdaef", "#e49196", "#ff9d57", "#97f9a1"];

    useEffect(() => {
        setLoading(true)
        layersService.getLayers()
            .then(data => {
                setLayers(data)
                setLoading(false)
            })
            .catch(err => {
                console.error("Failed to fetch layers: ", err)
                setLoading(false)
            })
    }, [])

    if (loading || layers.length === 0) {
        return (
            <StyledSidePanel>
                <ButtonsWrapper>
                </ButtonsWrapper>
                {/*<DesignComponentInstanceNode mode={2} />*/}
            </StyledSidePanel>
        )
    }

    return (
        <StyledSidePanel>
            <ButtonsWrapper>
                {layers.map((object, i) =>
                    <LayerButton
                        text={object}
                        background={colors[i % colors.length]}
                        selected={object === currentLayer}
                        setCurrentLayer={setCurrentLayer}
                    />)}
                {/*<LayerButtonAdd text="+" />*/}
            </ButtonsWrapper>
            {/*<DesignComponentInstanceNode mode={2} />*/}
        </StyledSidePanel>
    );
};
