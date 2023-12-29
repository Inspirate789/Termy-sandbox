import React, {useEffect, useState} from "react";
import styled from "styled-components";
import { LangWrapper } from "./LangWrapper";
import { TermWrapper } from "./TermWrapper";
import { TermWrapperWithInfo } from "./TermWrapperWithInfo";
import unitsService from "../services/UnitsService";
import {InputUnit} from "../types";
import layersService from "../services/LayersService";
import config from "../config";

const StyledSearchResults = styled.div`
    //align-items: flex-start;
    border: 3px solid;
    border-color: #00abe4;
    border-radius: 15px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    height: fit-content;
    min-height: 95%;
    overflow: hidden;
    padding: 0px 0px 10px;
    position: relative;
    width: 100%;
  
    & .lang-wrapper-instance {
    align-self: stretch !important;
    flex: 0 0 auto !important;
    width: 100% !important;
    }

    & .term-wrapper-instance {
    align-self: stretch !important;
    height: 100% !important;
    width: 100% !important;
    }
`;

interface Props {
    lang: string
    layer: string
    pattern: string
}

interface UnitWithState {
    unit: InputUnit
    opened: boolean
    setOpened: (opened: boolean) => void
}

export const SearchResults = ({ lang, layer, pattern }: Props) => {
    const [loading, setLoading] = useState(false)
    const [units, setUnits] = useState<InputUnit[]>([]);

    useEffect(() => {
        setLoading(true)
        unitsService.getUnits(layer, lang, pattern, 0, config.units_per_page)
            .then(data => {
                setUnits(data)
                setLoading(false)
            })
            .catch(err => {
                console.error("Failed to fetch units: ", err)
                setUnits([]);
                setLoading(false)
            })
    }, [lang, layer, pattern])

    if (loading || units.length === 0) {
        return (
            <StyledSearchResults>
                <LangWrapper className="lang-wrapper-instance" lang={lang === "ru" ? "Русский" : "Английский"}/>
            </StyledSearchResults>
        )
    }

    const deleteUnit = (unit: InputUnit, i: number) => {
        unitsService.deleteUnit(layer, {lang: lang, text: unit.text})
            .then(() => {
                const unitsCopy = Array.from(units);
                unitsCopy.splice(i, 1);
                setUnits(unitsCopy);
            })
            .catch(err => {
                console.error("Failed to delete unit: ", err);
            })
    }

    return (
        <StyledSearchResults>
            <LangWrapper className="lang-wrapper-instance" lang={lang === "ru" ? "Русский" : "Английский"}/>
            {units.map((object, i) => {
                return <TermWrapperWithInfo className="term-wrapper-instance" unit={object} onDelete={() => deleteUnit(object, i)}/>;
            })}
        </StyledSearchResults>
    );
};