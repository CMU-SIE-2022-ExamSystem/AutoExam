import React from "react";
import {useSearchParams} from "react-router-dom";

function AuthCallback() {
    let [searchParams, setSearchParams] = useSearchParams();

    return (
        <div>
            {searchParams.toString()}
        </div>
    );
}

export default AuthCallback;
