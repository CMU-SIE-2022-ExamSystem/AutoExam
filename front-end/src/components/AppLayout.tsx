import React from 'react';
import {Container} from "react-bootstrap";

function AppLayout({children} : {children: React.ReactNode}) {
    return (
        <Container className="my-3 d-flex flex-column text-center">
            {children}
        </Container>
    );
}

export default AppLayout;
