import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import {CookiesProvider} from "react-cookie";

import AuthRedirect from './routes/auth/AuthRedirect';
import AuthCallback from './routes/auth/AuthCallback';
import Assessments from "./routes/course/Assessments";
import Dashboard from "./routes/course/Dashboard";
import ExamInstructions from './routes/course/ExamInstructions';
import ExamQuestions from './routes/course/ExamQuestions';

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <CookiesProvider>
        <BrowserRouter>
            <Routes>
                <Route path='/' element={<App/>}/>
                <Route path="dashboard" element={<Dashboard/>}/>
                <Route path="assessments" element={<Assessments/>}/>
                <Route path="assessments/:exam_id" element={<ExamInstructions/>}/>
                <Route path="assessments/:exam_id/questions" element={<ExamQuestions/>}/>
                <Route path="questions" element={<ExamQuestions/>}/>
                <Route path="oauth" element={<AuthRedirect/>}/>
                <Route path="oauth-callback" element={<AuthCallback/>}/>
            </Routes>
        </BrowserRouter>
    </CookiesProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
