import React, {useCallback, useEffect, useState} from "react";
import {useSearchParams} from "react-router-dom";
import ErrorLayout from "../../components/ErrorLayout";
import {AxiosError, AxiosResponse} from "axios";
import {getBackendApiUrl, getFrontendUrl} from "../../utils/url";
import {useCookies} from "react-cookie";

const axios = require('axios').default;

function AuthCallback() {
    let [authCode, setAuthCode] = useState("N/A");
    let [searchParams] = useSearchParams();
    const [cookies, setCookie] = useCookies(['token']);

    const CallBack = useCallback(async () => {
        const stateValue = localStorage.getItem('authStateValue')
        const stateQuery = searchParams.get('state')

        if (stateValue !== stateQuery) {
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
                setCookie('token', data.token, {path: '/'});

                window.location.assign(getFrontendUrl('/dashboard'));
            })
            .catch((error: AxiosError) => {
                //Error
            })
    }, [searchParams, setCookie])

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
