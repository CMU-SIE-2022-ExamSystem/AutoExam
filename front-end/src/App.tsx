import React from 'react';
import './App.css';
import {Route, Routes} from "react-router-dom";
import Dashboard from "./routes/course/Dashboard";
import Assessments from "./routes/course/Assessments";
import AuthRedirect from "./routes/auth/AuthRedirect";
import AuthCallback from "./routes/auth/AuthCallback";
import Index from "./routes/Index";
import {GlobalStateProvider} from "./components/GlobalStateProvider";
import ExamQuestions from "./routes/course/ExamQuestions";
import ExamInstructions from "./routes/course/ExamInstructions";

const App = () => {
    return (
        <div className="App">
            <GlobalStateProvider>
                <Routes>
                    <Route path='/' element={<Index />}/>
                    <Route path="dashboard" element={<Dashboard/>}/>
                    <Route path="assessments" element={<Assessments/>}/>
                    <Route path="assessments/:exam_id" element={<ExamInstructions/>}/>
                    <Route path="assessments/:exam_id/questions" element={<ExamQuestions/>}/>
                    <Route path="questions" element={<ExamQuestions/>}/>
                    <Route path="oauth" element={<AuthRedirect/>}/>
                    <Route path="oauth-callback" element={<AuthCallback/>}/>
                </Routes>
            </GlobalStateProvider>
        </div>
    );
}

export default App;
