import React from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Col, Nav, Row, Tab} from "react-bootstrap";

const ExamResultQuestions = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    return (
        <div>
            Questions
        </div>
    )
}

export default ExamResultQuestions;
