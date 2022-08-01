import React from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Row} from "react-bootstrap";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";

const ExamResults = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={courseName} brandLink={"/courses/"+courseName}/>
            </Row>
            <Row>
                Exam Results Page: {courseName}, {examId}
            </Row>
        </AppLayout>
    )
}

export default ExamResults;
