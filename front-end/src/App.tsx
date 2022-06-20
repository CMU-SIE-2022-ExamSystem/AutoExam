import React, {useState} from 'react';
import './App.css';
import {Route, Routes} from "react-router-dom";
import Dashboard from "./routes/course/Dashboard";
import Assessments from "./routes/course/Assessments";
import AuthRedirect from "./routes/auth/AuthRedirect";
import AuthCallback from "./routes/auth/AuthCallback";
import Index from "./routes/Index";

import {GlobalStateProvider} from "./components/GlobalStateProvider";

const App = () => {
    const [config, setConfig] = useState<any>({});
    return (
        <div className="App">
            <GlobalStateProvider>
                <Routes>
                    <Route path='/' element={<Index />}/>
                    <Route path="dashboard" element={<Dashboard/>}/>
                    <Route path="assessments" element={<Assessments/>}/>
                    <Route path="oauth" element={<AuthRedirect/>}/>
                    <Route path="oauth-callback" element={<AuthCallback/>}/>
                </Routes>
            </GlobalStateProvider>
        </div>
    );
}

export default App;
