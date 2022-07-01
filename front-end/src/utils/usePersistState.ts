import React from 'react';

const usePersistState = (defaultValue : any, storageKey : string) => {
    const [value, setValue] = React.useState(() => {
        const storageValue = window.localStorage.getItem(storageKey);

        return storageValue !== null ? JSON.parse(storageValue) : defaultValue;
    });

    React.useEffect(() => {
        window.localStorage.setItem(storageKey, JSON.stringify(value));
    }, [storageKey, value]);

    return [value, setValue];
}

export default usePersistState;
