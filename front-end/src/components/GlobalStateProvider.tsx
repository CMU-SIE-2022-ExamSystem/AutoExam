import React, {createContext, useState} from "react";

export interface GlobalStateInterface {
    name: string | null;
    token: string | null;
}

const initialState: GlobalStateInterface = {name: null, token: null};

export interface GlobalStateContextType {
    globalState: GlobalStateInterface;
    setGlobalState: React.Dispatch<React.SetStateAction<GlobalStateInterface>>;
}

const GlobalStateContext = createContext<GlobalStateContextType>(null!);

export const GlobalStateProvider = ({children} : {children: React.ReactNode}) => {
    let state = JSON.parse(sessionStorage.getItem('globalState') || JSON.stringify(initialState));
    let [globalState, setGlobalState] = useState<GlobalStateInterface>(state);
    let value = {globalState, setGlobalState};
    return <GlobalStateContext.Provider value={value}>{children}</GlobalStateContext.Provider>;
}

export const useGlobalState = () => React.useContext(GlobalStateContext);