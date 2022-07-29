import React from 'react';
import {Container} from "react-bootstrap";
import {RightBottomAlert} from "./RightBottomAlert";
import {useGlobalState} from "./GlobalStateProvider";

function AppLayout({children} : {children: React.ReactNode}) {
    const {globalState, updateGlobalState} = useGlobalState();
    let alertState = globalState.alert;
    if (!alertState) alertState = {variant: "primary", content: "", show: false};
    const closeUpdate = () => {
        if (!alertState) return;
        alertState.show = false;
        updateGlobalState({alert: alertState});
    }
    return (
        <Container fluid className="d-flex flex-column text-center max-vh-100">
            {children}
            <RightBottomAlert variant={alertState.variant} content={alertState.content} show={alertState.show} onClose={closeUpdate} />
        </Container>
    );
}

export default AppLayout;
