import React, {useState} from "react";
import { HeaderAccount } from "../components/HeaderAccount";
import { SearchResults } from "../components/SearchResults";
import { SidePanel } from "../components/SidePanel";
import styled from "styled-components";
import {useLocation, useNavigate} from "react-router-dom";

const StyledTermsInfo = styled.div`
  align-items: flex-start;
  background-color: #e9f1fa;
  display: flex;
  flex-direction: column;
  position: relative;
  min-height: 98.4vh;
  height: 100%;
  width: 99.3%;

  & .design-component-instance-node-3 {
    align-self: stretch !important;
    flex: 0 0 auto !important;
    width: 100% !important;
  }

  & .searcher-wrapper {
    align-items: center;
    align-self: stretch;
    display: flex;
    flex: 1;
    flex-grow: 1;
    gap: 10px;
    justify-content: center;
    padding: 0px 10px 10px 0px;
    position: relative;
    width: 100%;
  }

  & .side-panel-instance {
    align-self: stretch !important;
    flex: 0 0 auto !important;
    height: unset !important;
  }

  & .searcher {
    align-items: center;
    align-self: stretch;
    background-color: #e9f1fa;
    display: flex;
    flex: 1;
    flex-direction: column;
    flex-grow: 1;
    gap: 10px;
    padding: 0px 10px;
    position: relative;
  }

  & .input-line-wrapper {
    align-items: flex-start;
    align-self: stretch;
    display: flex;
    flex: 0 0 auto;
    gap: 10px;
    justify-content: center;
    padding: 10px;
    position: relative;
    width: 100%;
  }

  & .input-line {
    align-items: center;
    background-color: #ffffff;
    border-radius: 15px;
    display: flex;
    justify-content: flex-end;
    overflow: hidden;
    padding: 5px 10px;
    position: relative;
    width: 429px;
  }

  & .text-7 {
    border: none;
    outline: none;
    color: #000000;
    flex: 1;
    font-family: "Inter-Regular", Helvetica;
    font-size: 20px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    margin-top: -0.5px;
    position: relative;
  }

  & .search-icon {
    height: 25px;
    object-fit: cover;
    position: relative;
    width: 25px;
  }

  & .search-results-2 {
    align-items: flex-start;
    align-self: stretch;
    display: flex;
    flex: 1;
    flex-grow: 1;
    gap: 10px;
    justify-content: center;
    overflow: hidden;
    padding: 0px 10px;
    position: relative;
    width: 100%;
  }

  & .search-results-3 {
    align-items: center;
    align-self: stretch;
    border: 3px solid;
    border-color: #00abe4;
    border-radius: 15px;
    display: flex;
    flex: 1;
    flex-direction: column;
    flex-grow: 1;
    gap: 10px;
    margin-bottom: -1.5px;
    margin-left: -1.5px;
    margin-top: -1.5px;
    overflow: hidden;
    padding: 0px 0px 10px;
    position: relative;
  }

  & .term-wrapper-2 {
    align-self: stretch !important;
    height: 30px !important;
    width: 100% !important;
  }

  & .search-results-instance {
    align-self: stretch !important;
    flex: 1 !important;
    flex-grow: 1 !important;
    height: unset !important;
    margin-bottom: -1.5px !important;
    margin-right: -1.5px !important;
    margin-top: -1.5px !important;
    width: unset !important;
  }
`;

function useQuery() {
    const { search } = useLocation();
    return React.useMemo(() => new URLSearchParams(search), [search]);
}

export const TermsInfo = () => {
    document.body.style.backgroundColor = '#e9f1fa';

    const query = useQuery();
    const navigate = useNavigate();
    const [layer, setLayer] = useState<string>(query.get("layer") || "");
    const [pattern, setPattern] = useState(query.get("pattern") || "");

    const handleInput = (e: any) => {
        e.persist();
        const { value } = e.target;
        setPattern(value);
    }

    const setCurrentLayer = (layer: string) => {
        query.set("layer", layer);
        const newSearch = `?${query.toString()}`;
        navigate({search: newSearch});
        setLayer(layer);
    };

    const search = () => {
        query.set("pattern", pattern);
        const newSearch = `?${query.toString()}`;
        navigate({search: newSearch});
        setPattern(pattern);
    };

    return (
        <StyledTermsInfo>
            <HeaderAccount/>
            <div className="searcher-wrapper">
                <SidePanel currentLayer={layer} setCurrentLayer={setCurrentLayer}/>
                <div className="searcher">
                    <div className="input-line-wrapper">
                        <div className="input-line">
                            <input className="text-7" onChange={handleInput} value={pattern}/>
                            <button 
                                style={{width: "fit-content", height: "fit-content", background: "none", border: "none"}}
                                onClick={search}
                            >
                                <img className="search-icon" alt="" src="assets/search-icon.png" />
                            </button>
                        </div>
                    </div>
                    <div className="search-results-2">
                        <SearchResults lang="ru" layer={layer} pattern={pattern}/>
                        <SearchResults lang="en" layer={layer} pattern={pattern}/>
                    </div>
                </div>
            </div>
        </StyledTermsInfo>
    );
};
