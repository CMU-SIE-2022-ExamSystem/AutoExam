import React from 'react';

const usePersistState = <T>(defaultValue : T, storageKey : string) => {
    const [value, setValue] = React.useState<T>(() => {
        const storageValue = window.localStorage.getItem(storageKey);

        return storageValue !== null ? JSON.parse(storageValue) : defaultValue;
    });

    React.useEffect(() => {
        window.localStorage.setItem(storageKey, JSON.stringify(value));
    }, [storageKey, value]);

    const removeValue = () => {
        window.localStorage.removeItem(storageKey);
    }

    const getValue = () => {
        return JSON.parse(window.localStorage.getItem(storageKey) as string);
    }

    return {value, setValue, removeValue, getValue};
}

export default usePersistState;
