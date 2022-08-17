import React, {createContext, useState} from "react";
import {GlobalAlertProperties} from "./globalAlert";

/**
 * The global states that remains value between the web pages. Used to store some session information, such as token, name, etc.
 */
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
    // A shortcut to update partial information of the global state, instead of setting the whole.
    const updateGlobalState = (updateOptions : any) => {
        const newState = Object.assign({}, globalState, updateOptions);
        setGlobalState(newState);
        return newState;
    }
    let value = {globalState, setGlobalState, updateGlobalState};
    return <GlobalStateContext.Provider value={value}>{children}</GlobalStateContext.Provider>;
}

/**
 * To use: call hook useGlobalState() as other React hooks. See React documentation for help.
 */
export const useGlobalState = () => React.useContext(GlobalStateContext);
