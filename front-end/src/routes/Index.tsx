import React from 'react';
import TopNavbar from "../components/TopNavbar";
import AppLayout from "../components/AppLayout";
import {Link} from "react-router-dom";
import {Image, Row} from "react-bootstrap";
import autolab_logo from "../images/autolab_logo.png";

function Index() {
    return (
        <AppLayout>
            <Row>
                <TopNavbar />
            </Row>
            <main className="p-4">
                <h1>AutoExam System</h1>
                <h4 className="mb-3">This website needs authorization from Autolab.</h4>
                <div className="mb-3"><p>Click on the Autolab logo to login.</p></div>
                <div><Link to="/oauth"><Image src={autolab_logo} alt="Autolab Logo"/></Link></div>
            </main>
        </AppLayout>
    );
}

export default Index;
