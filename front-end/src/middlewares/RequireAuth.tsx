import {useGlobalState} from "../components/GlobalStateProvider";
import {Navigate} from "react-router-dom";


const RequireAuth = ({children} : {children: React.ReactNode}) => {
    const {globalState} = useGlobalState();
    if (process.env.NODE_ENV !== 'development' && (!globalState.token || globalState.token.length === 0)) {
        return <Navigate to="/oauth" replace />;
    }
    return <>{children}</>;
}

export default RequireAuth;
