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
import RequireAuth from "./middlewares/RequireAuth";

const App = () => {
    return (
        <div className="App">
            <GlobalStateProvider>
                <Routes>
                    <Route path='/' element={<Index/>}/>
                    <Route path="oauth" element={<AuthRedirect/>}/>
                    <Route path="oauth-callback" element={<AuthCallback/>}/>
                    <Route path="dashboard" element={
                        <RequireAuth>
                            <Dashboard/>
                        </RequireAuth>
                    }/>
                    <Route path="courses/:course_name">
                        <Route index element={<RequireAuth><Assessments/></RequireAuth>}/>
                        <Route path="exams/:exam_id">
                            <Route index element={<RequireAuth><ExamInstructions/></RequireAuth>}/>
                            <Route path="questions" element={<RequireAuth><ExamQuestions/></RequireAuth>}/>
                        </Route>
                    </Route>
                </Routes>
            </GlobalStateProvider>
        </div>
    );
}

export default App;
