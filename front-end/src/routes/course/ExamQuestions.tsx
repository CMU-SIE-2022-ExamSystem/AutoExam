import React, {useCallback, useState} from 'react';
import {Button, Col, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import AppLayout from "../../components/AppLayout";
import Question from "../../components/Question";
import CountdownTimer from "../../components/CountdownTimer";
import questionDataType from "../../components/questionTemplate/questionDataType";
import usePersistState from "../../utils/usePersistState";
import { choiceDataType, subQuestionDataType } from '../../components/questionTemplate/subQuestionDataType';

const getQuestionList = () => {
    return [];
}

const TimeoutModal = ({show, toClose, onClose} :{ show: boolean, toClose: () => void, onClose: () => void }) => {
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Test over</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>This test is over. Your answers have been recorded.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={toClose}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const AcknowledgeModal = ({show, toClose, onClose} :{ show: boolean, toClose: () => void, onClose: () => void }) => {
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Submitted</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>Your answers have been recorded.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={toClose}>Confirm</Button>
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
    var subQuestionArray = questionList.flatMap((question) =>
        question.questions.map(subQuestion => ["Q" + question.headerId + "_sub" + subQuestion.questionId, subQuestion.questionType, subQuestion.choices]));
    var idList: string[] = [];
    for (var i = 0; i < subQuestionArray.length; i++) {
        if (subQuestionArray[i][2].length === 1 && subQuestionArray[i][2][0] === "") { // single blank
            idList.push(subQuestionArray[i][0].toString());
            continue;
        }
        
        for (var j = 0; j < subQuestionArray[i][2].length; j++) {
            if (subQuestionArray[i][1] === "single-choice" || subQuestionArray[i][1] === "multiple-choice") {
                idList.push(subQuestionArray[i][0].toString() + "_choice" + (subQuestionArray[i][2][j] as choiceDataType).choiceId);
            } else { // multiple-blank
                idList.push(subQuestionArray[i][0].toString() + "_sub" + (subQuestionArray[i][2][j] as choiceDataType).choiceId);
            }
        }
    }
    console.log(idList);

    const [timeoutShow, setTimeoutShow] = useState(false);
    const [ackShow, setAckShow] = useState(false);
    const [confirmShow, setConfirmShow] = useState(false);

    const [inTest, setInTest] = useState(true);

    const [targetTime] = usePersistState(new Date(Date.now() + 1000 * 100).toString(), "targetTime");

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
        setInTest(false);
    }

    return (
        <AppLayout>
            <Row className="flex-grow-1">
                <Col xs={9} className="p-3 overflow-auto vh-100 bottom-0">
                    <Instructions info={instructionsInfo} />
                    <Row>
                        <Col xs={{ span: 10, offset: 1 }}>
                            { questions }
                        </Col>
                    </Row>
                    <br/>
                </Col>
                <Col xs={3} className="p-3">
                    <CountdownTimer targetTime={targetTime} active={inTest} callback={() => setTimeoutShow(true)} />
                    <div><Button variant="primary" className="my-4 w-25" onClick={() => setConfirmShow(true)}>Submit</Button></div>
                </Col>
            </Row>

            <TimeoutModal onClose={() => {}} show={timeoutShow} toClose={() => setTimeoutShow(false)} />
            <AcknowledgeModal onClose={() => {}} show={ackShow} toClose={() => setAckShow(false)} />
            <ConfirmModal show={confirmShow} onSubmit={submitExam} onClose={() => setConfirmShow(false)} />
        </AppLayout>
    );
}

export default ExamQuestions;
