import React, {useCallback, useState} from 'react';
import {Button, Col, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import Question from "../../components/Question";
import CountdownTimer from "../../components/CountdownTimer";
import questionDataType from "../../components/questionTemplate/questionDataType";

const getQuestionList = () => {
    return [];
}

const TimeoutModal = ({show, onClose} :{ show: boolean, onClose: () => void }) => {
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Test over</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>This test is over. Your answers have been recorded.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary">Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const AcknowledgeModal = ({show, onClose} :{ show: boolean, onClose: () => void }) => {
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Submitted</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>Your answers have been recorded.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary">Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const ConfirmModal = ({show, onSubmit, onClose} :{ show: boolean, onSubmit: () => void, onClose: () => void }) => {
    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Confirmation</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>Do you want to submit early?</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={onSubmit}>Submit</Button>
                <Button variant="danger" onClick={onClose}>Back</Button>
            </Modal.Footer>
        </Modal>
    );
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

const ExamQuestions = () => {
    let params = useParams();
    let questionList: questionDataType[];
    //useCallback(() => questionList = getQuestionList(), []);

    questionList = require('./questions_new.json').data;

    const [timeoutShow, setTimeoutShow] = useState(false);
    const [ackShow, setAckShow] = useState(false);
    const [confirmShow, setConfirmShow] = useState(false);

    const [targetTime] = useState(new Date(Date.now() + 1000 * 100).toString());

    let instructionsInfo : instructionType = {
        title: params.exam_id!,
        instructions: "",
    }

    const questions = questionList.map((question) => {
        return <Question questionData={question} />
    })

    const submitExam = () => {
        setConfirmShow(false);
        setAckShow(true);
    }

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={null}/>
            </Row>
            <Row className="flex-grow-1 justify-content-center">
                <Col xs={9} className="overflow-auto p-3">
                    <Instructions info={instructionsInfo} />
                    { questions }
                    <br/>
                </Col>
                <Col xs={3} className="p-3">
                    <CountdownTimer targetTime={targetTime} callback={() => {}} />
                </Col>
            </Row>
            <div><Button variant="primary" className="my-4 w-25" onClick={() => setConfirmShow(true)}>Submit</Button></div>
            <TimeoutModal onClose={() => setTimeoutShow(false)} show={timeoutShow} />
            <AcknowledgeModal onClose={() => setAckShow(false)} show={ackShow} />
            <ConfirmModal show={confirmShow} onSubmit={submitExam} onClose={() => setConfirmShow(false)} />
        </AppLayout>
    );
}

export default ExamQuestions;
