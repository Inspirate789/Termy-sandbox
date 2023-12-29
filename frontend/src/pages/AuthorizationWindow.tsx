import React from "react";
import styled from "styled-components";
import { AuthorizationBox } from "../components/AuthorizationBox";
import { TitleWrapper } from "../components/TitleWrapper";
import {AuthRequest} from "../types";
import authService from "../services/AuthService";
import {useNavigate} from "react-router-dom";

const StyledAuthorizationWindow = styled.div`
  align-items: center;
  background-color: #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 95vh;
  height: 100%;
  padding: 10px;
  position: relative;

  .title-wrapper-instance {
    align-self: stretch !important;
    flex: 0 0 auto !important;
    width: 100% !important;
  }

  .authorization {
    align-items: center;
    align-self: stretch;
    display: flex;
    flex: 1;
    flex-direction: column;
    flex-grow: 1;
    gap: 10px;
    justify-content: center;
    padding: 10px;
    position: relative;
    width: 100%;
  }

  .title-wrapper-clone {
    align-items: flex-end;
    align-self: stretch;
    border-radius: 10px;
    display: flex;
    flex: 0 0 auto;
    overflow: hidden;
    padding: 10px 0px;
    position: relative;
    width: 100%;
  }
`;

export const AuthorizationWindow = () => {
    const navigate = useNavigate();

    const onSuccess = async (user: AuthRequest) => {
        try {
            await authService.login(user);
            navigate("/info");
        } catch (error) {
            console.error("Auth error", error);
            alert(error);
        }
    }

    return (
        <StyledAuthorizationWindow>
            <TitleWrapper className="title-wrapper-instance" />
            <div className="authorization">
                <AuthorizationBox  onSuccess={onSuccess}/>
            </div>
        </StyledAuthorizationWindow>
    );
};
