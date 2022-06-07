import React from 'react';
import './App.css';
import TopNavbar from "./components/TopNavbar";
import AppLayout from "./components/AppLayout";
import {Link} from "react-router-dom";

function App() {
    return (
        <div className="App">
            <TopNavbar />
            <AppLayout>
                <div>
                    <h1>This website needs authorization from Autolab.</h1>
                    <Link to="/oauth">OAuth 2.0 Redirect</Link>
                </div>
            </AppLayout>
        </div>
    );
}

export default App;
