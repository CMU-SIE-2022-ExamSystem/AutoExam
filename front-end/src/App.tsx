import React from 'react';
import './App.css';
import {Route, Routes} from "react-router-dom";
import Dashboard from "./routes/course/Dashboard";
import Assessments from "./routes/course/Assessments";
import AuthRedirect from "./routes/auth/AuthRedirect";
import AuthCallback from "./routes/auth/AuthCallback";
import Index from "./routes/Index";
import {GlobalStateProvider} from "./components/GlobalStateProvider";
import ExamQuestions from "./routes/course/exams/ExamQuestions";
import ExamInstructions from "./routes/course/exams/ExamInstructions";
import RequireAuth from "./middlewares/RequireAuth";
import QuestionBank from "./routes/course/questionBanks/QuestionBank";
import ExamConfig from "./routes/course/config/ExamConfig";
import {ExamConfigStateProvider} from "./routes/course/config/ExamConfigStates";

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
                        <Route path="questionBank/:tag" element={<RequireAuth><QuestionBank/></RequireAuth>}/>
                        <Route path="examConfig">
                            <Route path=":exam_id" element={<RequireAuth><ExamConfigStateProvider><ExamConfig/></ExamConfigStateProvider></RequireAuth>} />
                        </Route>
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
