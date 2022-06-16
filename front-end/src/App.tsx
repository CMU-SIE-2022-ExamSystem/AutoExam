import React, {useState} from 'react';
import './App.css';
import {Route, Routes} from "react-router-dom";
import Dashboard from "./routes/course/Dashboard";
import Assessments from "./routes/course/Assessments";
import AuthRedirect from "./routes/auth/AuthRedirect";
import AuthCallback from "./routes/auth/AuthCallback";
import Index from "./routes/Index";

import GlobalConfigContext from "./components/GlobalConfigContext";

const App = () => {
    const [globalConfig] = useState(null);
    return (
        <div className="App">
            <GlobalConfigContext.Provider value={globalConfig}>
                <Routes>
                    <Route path='/' element={<Index />}/>
                    <Route path="dashboard" element={<Dashboard/>}/>
                    <Route path="assessments" element={<Assessments/>}/>
                    <Route path="oauth" element={<AuthRedirect/>}/>
                    <Route path="oauth-callback" element={<AuthCallback/>}/>
                </Routes>
            </GlobalConfigContext.Provider>
        </div>
    );
}

export default App;
