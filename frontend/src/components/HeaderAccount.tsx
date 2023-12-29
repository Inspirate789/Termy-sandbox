import React from "react";
import styled from "styled-components";
import { TitleWrapper } from "./TitleWrapper";
import {useNavigate} from "react-router-dom";

const StyledHeaderAccount = styled.div`
  align-items: flex-start;
  background-color: #e9f1fa;
  display: flex;
  justify-content: flex-end;
  padding: 10px 10px 25px;
  position: relative;
  width: 99%;

  .title-wrapper-instance {
    flex: 1 !important;
    flex-grow: 1 !important;
    width: unset !important;
  }

  .name-wrapper {
    align-items: flex-start;
    display: inline-flex;
    flex: 0 0 auto;
    flex-direction: column;
    padding: 0px 10px;
    position: relative;
  }

  .text-wrapper {
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

  .div {
    color: #000000;
    font-family: "Inter-Regular", Helvetica;
    font-size: 14px;
    font-weight: 400;
    letter-spacing: 0;
    line-height: normal;
    position: relative;
    width: 100%;
  }

  .exit-button-wrapper {
    border-width: 0;
    align-items: flex-start;
    background-color: #00abe4;
    border-radius: 10px;
    display: inline-flex;
    flex: 0 0 auto;
    justify-content: flex-end;
    overflow: hidden;
    padding: 7px;
    position: relative;
  }

  .logout-icon {
    height: 30px;
    object-fit: cover;
    position: relative;
    width: 30px;
  }
`;

export const HeaderAccount = () => {
    const navigate = useNavigate();

    const onSubmit = () => {
        return navigate("/");
    }

    return (
        <StyledHeaderAccount>
            <TitleWrapper className="title-wrapper-instance" />
            <div className="name-wrapper">
                <div className="text-wrapper">Admin</div>
                <div className="div">Admin</div>
            </div>
            <button className="exit-button-wrapper" onClick={onSubmit}>
                <img className="logout-icon" alt="Logout icon" src="assets/logout-icon.png" />
            </button>
        </StyledHeaderAccount>
    );
};
