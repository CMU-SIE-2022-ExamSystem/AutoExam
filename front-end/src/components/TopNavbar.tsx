import React from 'react';
import {Container, Nav, Navbar, NavDropdown} from 'react-bootstrap';
import {useGlobalState} from "./GlobalStateProvider";

type Props = {
    brand? : string | null,
}

const TopNavbar = ({brand} : Props) => {
    const brandName: string = brand || "ExamServer"
    const {globalState} = useGlobalState();
    const username: string = globalState.name || "Guest";
    return (
        <>
            <Navbar collapseOnSelect expand="md" bg="dark" variant="dark" className="text-start">
                <Container fluid>
                    <Navbar.Brand href="#home">
                        {brandName}
                    </Navbar.Brand>
                    <Navbar.Toggle />
                    <Navbar.Collapse id="top-navbar-nav">
                        <Nav className="me-auto">
                            <Nav.Link href="/dashboard">Dashboard</Nav.Link>
                            <Nav.Link href="#">Help</Nav.Link>
                        </Nav>
                        <Nav>
                            <NavDropdown title={username} id="top-navbar-dropdown">
                                <NavDropdown.Item href="#">Logout</NavDropdown.Item>
                            </NavDropdown>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </>
    );
}

export default TopNavbar;
