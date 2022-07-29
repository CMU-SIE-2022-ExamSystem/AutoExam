import React, {createContext, useState} from "react";
import {GlobalAlertProperties} from "./globalAlert";

export interface GlobalStateInterface {
    name: string | null;
    token: string | null;
    alert: GlobalAlertProperties | null;
}

const initialState: GlobalStateInterface = {name: null, token: null, alert: null};

export interface GlobalStateContextType {
    globalState: GlobalStateInterface;
    setGlobalState: React.Dispatch<React.SetStateAction<GlobalStateInterface>>;
    updateGlobalState: (updateOptions: any) => GlobalStateInterface;
}

const GlobalStateContext = createContext<GlobalStateContextType>(null!);

export const GlobalStateProvider = ({children} : {children: React.ReactNode}) => {
    let state = JSON.parse(sessionStorage.getItem('globalState') || JSON.stringify(initialState));
    let [globalState, setGlobalState] = useState<GlobalStateInterface>(state);
    const updateGlobalState = (updateOptions : any) => {
        const newState = Object.assign({}, globalState, updateOptions);
        setGlobalState(newState);
        return newState;
    }
    let value = {globalState, setGlobalState, updateGlobalState};
    return <GlobalStateContext.Provider value={value}>{children}</GlobalStateContext.Provider>;
}

export const useGlobalState = () => React.useContext(GlobalStateContext);
