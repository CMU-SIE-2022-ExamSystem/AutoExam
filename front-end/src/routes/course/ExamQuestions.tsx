import React from 'react';
import {Button, Col, Row} from 'react-bootstrap';
import { useParams } from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import Question from "../../components/Question";
import CountdownTimer from "../../components/CountdownTimer";

const QuestionList = () => {

}

interface instructionType {
    title: string;
    instructions: string;
}

const Instructions = ({info}: {info: instructionType}) => {

    return (
        <div>
            <h1 className="my-3">{info.title}</h1>
            <h2 className="text-start my-4"><strong>Instructions</strong></h2>
            <p className="text-start">Some detailed instructions.</p>
        </div>
    );
}

function ExamQuestions() {
    let params = useParams();
    const questionList = QuestionList();

    const targetTime = new Date(Date.now() + 1000 * 100).toString();

    let instructionsInfo : instructionType = {
        title: params.exam_id!,
        instructions: "",
    }
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={null}/>
            </Row>
            <Row className="flex-grow-1 justify-content-center">
                <Col xs={9} className="overflow-auto p-3">
                    <Instructions info={instructionsInfo} />
                    <Question questionData={{}} />
                    <br/>
                </Col>
                <Col xs={3} className="p-3">
                    <CountdownTimer targetTime={targetTime} callback={() => {}} />
                </Col>
            </Row>
            <div><Button variant="primary"className="my-4 w-25">Submit</Button></div>
        </AppLayout>
    );
}

export default ExamQuestions;
