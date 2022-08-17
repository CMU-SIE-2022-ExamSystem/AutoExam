import React, {useEffect} from "react";
import {useNavigate} from "react-router-dom";
import {useGlobalState} from "../../components/GlobalStateProvider";

const Logout = () => {
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    useEffect(() => {
        if (globalState.token) {
            sessionStorage.removeItem('globalState');
            updateGlobalState({name: null, token: null});
        }
        navigate("/");
    }, []);

    return (
        <div>

        </div>
    );
}

export default Logout;
