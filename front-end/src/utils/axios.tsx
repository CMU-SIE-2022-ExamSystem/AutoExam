import React from "react";
import { useCookies } from 'react-cookie';
const axios = require('axios').default;

// Frontend Path
const secureGet = (url: string, config: any = {}) => {
    config.withCredentials = true;
    return axios.get(url, config);
}

// Backend Path
const securePost = (url: string, data: any, config: any = {}) => {
    config.withCredentials = true;
    return axios.post(url, data,config);
}

export {secureGet, securePost};
