import React from 'react';

// Frontend Path
const getFrontendUrl = (path: string) => {
    return window.location.origin + path;
}

// Backend Path
const getBackendApiUrl = (path: string) => {
    return process.env.REACT_APP_BACKEND_API_ROOT + path;
}

export {getFrontendUrl, getBackendApiUrl};
