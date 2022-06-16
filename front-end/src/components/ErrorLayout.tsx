import React from 'react';
import {Container, Alert} from "react-bootstrap";

type Props = {
    children?: JSX.Element
}

function ErrorLayout({children} : Props) {
    return (
        <Container className="my-3 d-flex flex-column text-center">
            <Alert key="danger" variant="danger">
                {children}
            </Alert>
        </Container>
    );
}

export default ErrorLayout;
