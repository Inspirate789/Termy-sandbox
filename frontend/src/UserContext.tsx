import React from "react";
import {createContext, useContext} from "react";


interface UserContextType {
    userId: number | null;
    setUserId?: React.Dispatch<React.SetStateAction<number | null>>;
    token: string | null;
    setToken?: (t: string | null) => void;
    isAuthenticated: boolean;
    setIsAuthenticated?: (auth: boolean) => void;
    role?: string | null;
    setRole?: (t: string | null) => void;
}


const initialState: UserContextType = {
    userId: null,
    token: null,
    isAuthenticated: false,
    role: null
};

export const UserContext = createContext<UserContextType>(initialState);

export const useUser = () => {
    const context = useContext(UserContext);
    if (!context.setUserId) {
        throw new Error("useUser must be used within a UserContext.Provider that defines setUserId");
    }
    return context;
};
