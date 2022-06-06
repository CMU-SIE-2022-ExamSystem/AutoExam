import React from 'react';

type Props = {
    children?: JSX.Element
}

function AppLayout({children} : Props) {
    return (
        <div className="px-4 py-3 d-flex flex-column text-center">
            {children}
        </div>
    );
}

export default AppLayout;