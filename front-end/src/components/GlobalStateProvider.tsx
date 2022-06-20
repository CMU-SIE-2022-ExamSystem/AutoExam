import React, {createContext, Dispatch, SetStateAction, useContext, useState} from "react";

export interface GlobalStateInterface {
    name: string | null;
    token: string | null;
};

const GlobalStateContext = createContext({
    state: {} as Partial<GlobalStateInterface>,
    setState: {} as Dispatch<SetStateAction<Partial<GlobalStateInterface>>>
});

const GlobalStateProvider = ({
                                 children,
                                 value = {} as GlobalStateInterface,
                             }: {
    children: React.ReactNode;
    value?: Partial<GlobalStateInterface>;
}) => {
    const [state, setState] = useState(value);
    return (
        <GlobalStateContext.Provider value={{state, setState}}>
            {children}
        </GlobalStateContext.Provider>
    );
};

const useGlobalState = () => {
    const context = useContext(GlobalStateContext);
    if (!context) {
        throw new Error("useGlobalState must be used within a GlobalStateContext");
    }
    return context;
};

export { GlobalStateProvider, useGlobalState };