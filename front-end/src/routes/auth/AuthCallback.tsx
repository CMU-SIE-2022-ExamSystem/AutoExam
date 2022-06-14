import React from "react";
import {useSearchParams} from "react-router-dom";
import ErrorLayout from "../../components/ErrorLayout";
import {AxiosError, AxiosResponse} from "axios";
import {getBackendApiUrl, getFrontendUrl} from "../../utils/url";

const axios = require('axios').default;

function AuthCallback() {
    let [searchParams, setSearchParams] = useSearchParams();

    const stateValue = localStorage.getItem('authStateValue')
    const stateQuery = searchParams.get('state')

    if (stateValue !== stateQuery) {
        return (<ErrorLayout><div>Bad State Value</div></ErrorLayout>)
    }

    const authCode = searchParams.get('code');

    // Here, call auth API to the back-end
    const backendTokenUrl = getBackendApiUrl("/auth");
    axios.post(backendTokenUrl, {code: authCode})
        .then((result: AxiosResponse) => {
            // Success, jump to dashboard
            window.location.assign(getFrontendUrl('/dashboard'));
        })
        .catch((error: AxiosError) => {
            //Error
        })

    return (
        <div>
            Code: {authCode}.
            Waiting for redirecting...
        </div>
    );
}

export default AuthCallback;
