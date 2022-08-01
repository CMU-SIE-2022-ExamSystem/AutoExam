import React, {useCallback, useEffect, useState} from "react";
import {useNavigate, useSearchParams} from "react-router-dom";
import ErrorLayout from "../../components/ErrorLayout";
import {AxiosError, AxiosResponse} from "axios";
import {getBackendApiUrl} from "../../utils/url";
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
