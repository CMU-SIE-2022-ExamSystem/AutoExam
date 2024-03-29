import React, {useCallback, useEffect, useState} from "react";
import {useNavigate, useSearchParams} from "react-router-dom";
import ErrorLayout from "../../components/ErrorLayout";
import {AxiosError, AxiosResponse} from "axios";
import {getBackendApiUrl} from "../../utils/url";
import {useGlobalState} from "../../components/GlobalStateProvider";

const axios = require('axios').default;

/**
 * The callback component when the front-end receives the Authorization code from the Autolab.
 */
const AuthCallback = () => {
    let [authCode, setAuthCode] = useState<string>("N/A");
    let [searchParams] = useSearchParams();
    const {updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const CallBack = useCallback(async () => {
        const stateValue = localStorage.getItem('authStateValue')
        const stateQuery = searchParams.get('state')
        const ignoreState = searchParams.get('ignore_state');

        if (stateValue !== stateQuery && ignoreState !== "true") {
            return (<ErrorLayout><div>Bad State Value</div></ErrorLayout>)
        }

        const authCode = searchParams.get('code') || "ERROR";
        setAuthCode(authCode);

        // Here, call auth API to the back-end
        const backendTokenUrl = getBackendApiUrl("/auth/token");
        axios.post(backendTokenUrl, {code: authCode})
            .then((result: AxiosResponse) => {
                // Success, jump to dashboard

                const data = result.data.data;
                const myName = data.firstName + " " + data.lastName;

                let newState = updateGlobalState({name: myName, token: data.token});
                sessionStorage.setItem('globalState', JSON.stringify(newState));
                navigate("/dashboard");
            })
            .catch((error: AxiosError) => {
                //Error
            })
    }, [navigate, searchParams])

    useEffect(() => {
        CallBack();
    }, [CallBack]);

    return (
        <div>
            Code: {authCode}.
            Waiting for redirecting...
        </div>
    );
}

export default AuthCallback;
