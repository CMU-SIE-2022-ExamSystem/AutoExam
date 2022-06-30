import React from 'react';
import {Container} from "react-bootstrap";

function AppLayout({children} : {children: React.ReactNode}) {
    return (
        <Container fluid className="d-flex flex-column text-center max-vh-100">
            {children}
        </Container>
    );
}

export default AppLayout;
