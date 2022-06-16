import React from 'react';
import TopNavbar from "../components/TopNavbar";
import AppLayout from "../components/AppLayout";
import {Link} from "react-router-dom";

function Index() {
    return (
        <div>
            <TopNavbar />
            <AppLayout>
                <div>
                    <h1>Exam Server</h1>
                    <h4 className="mb-3">This website needs authorization from Autolab.</h4>
                    <div><Link to="/oauth">OAuth 2.0 Redirect</Link></div>
                </div>
            </AppLayout>
        </div>
    );
}

export default Index;
