import React, {useState} from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css';
import {AuthorizationWindow} from "./pages/AuthorizationWindow";
import {UserContext} from './UserContext'
import authService from "./services/AuthService";
import {TermsInfo} from "./pages/TermsInfo";

export const App = () => {
  const [userId, setUserId] = useState<number | null>(null);
  const [token, setToken] = useState<string | null>(authService.getToken());
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(Boolean(token));
  const [role, setRole] = useState<string | null>(null);

  return (
      <UserContext.Provider value={{ userId, setUserId, token, setToken, isAuthenticated, setIsAuthenticated, role, setRole}}>
        <BrowserRouter>
          <Routes>
            {/*<Route path="/" element={<OnlyMovies />} />*/}
            <Route path="/" element={<AuthorizationWindow />} />
            <Route path="/info" element={<TermsInfo />} />
          </Routes>
        </BrowserRouter>
      </UserContext.Provider>
  )
}

export default App;
