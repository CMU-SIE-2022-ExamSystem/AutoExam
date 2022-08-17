import {useGlobalState} from "../components/GlobalStateProvider";
import {Navigate} from "react-router-dom";
import React from "react";

/**
 * A middleware wrapper that forbids access to certain pages when the user is not logged in.
 * @param children
 * @constructor
 */
const RequireAuth = ({children} : {children: React.ReactNode}) => {
    const {globalState} = useGlobalState();
    if (process.env.NODE_ENV !== 'development' && (!globalState.token || globalState.token.length === 0)) {
        return <Navigate to="/oauth" replace />;
    }
    return <>{children}</>;
}

export default RequireAuth;
