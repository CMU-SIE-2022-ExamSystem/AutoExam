import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import AppLayout from "../../../components/AppLayout";
import Question from "../../../components/Question";
import CountdownTimer from "../../../components/CountdownTimer";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {choiceDataType} from '../../../components/questionTemplate/subQuestionDataType';
import downloadFile from "../../../utils/downloadFile";

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

const prepareAnswer = (qList: questionDataType[]) : object => {
    let result = {};

    function getMultipleBlankAnswer(choices: choiceDataType[], subQuestionKey: string) {
        let returnObject = {};
        choices.forEach((choice) => {
            let choiceId = subQuestionKey + '_sub' + choice.choiceId;
            let element = document.getElementById(choiceId);
            if (element !== null) {
                let answer = (element as HTMLInputElement).value;
                Object.defineProperty(returnObject, choiceId, {value: answer})
            } else {
                Object.defineProperty(returnObject, choiceId, {value: ""});
            }
        })
        return returnObject;
    }

    function getSingleBlankAnswer(subQuestionKey: string) {
        let returnObject = {};
        let element = document.getElementById(subQuestionKey);
        if (element !== null) {
            let answer = (element as HTMLInputElement).value;
            Object.defineProperty(returnObject, subQuestionKey, {value: answer})
        } else {
            Object.defineProperty(returnObject, subQuestionKey, {value: ""});
        }
        return returnObject;
    }

    function getSingleChoiceAnswer(choices: choiceDataType[], subQuestionKey: string) {
        return {};
    }

    function getMultipleChoiceAnswer(choices: choiceDataType[], subQuestionKey: string) {
        return {};
    }

    qList.forEach((question: questionDataType) => {
        const questionKey = "Q" + question.headerId;
        let subResult: object = {};
        question.questions.forEach((subQuestion) => {
            const subQuestionKey = questionKey + "_sub" + subQuestion.questionId;
            let answerObject: object = {};

            switch (subQuestion.questionType) {
                case "multiple-blank":
                    answerObject = getMultipleBlankAnswer(subQuestion.choices, subQuestionKey);
                    break;
                case "single-blank":
                    answerObject = getSingleBlankAnswer(subQuestionKey);
                    break;
                case "single-choice":
                    answerObject = getSingleChoiceAnswer(subQuestion.choices, subQuestionKey);
                    break;
                case "multiple-choice":
                    answerObject = getMultipleChoiceAnswer(subQuestion.choices, subQuestionKey);
                    break;
            }

            Object.defineProperty(subResult, subQuestionKey, {
                value: answerObject
            })
        })

        Object.defineProperty(result, questionKey, {
            value: subResult
        });
    })
    return result;
}


const ExamQuestions = () => {
    let params = useParams();
    let questionList: questionDataType[];
    //useCallback(() => questionList = getQuestionList(), []);

    questionList = require('./questions_new.json').data;
    let subQuestionArray = questionList.flatMap((question) =>
        question.questions.map(subQuestion => ["Q" + question.headerId + "_sub" + subQuestion.questionId, subQuestion.questionType, subQuestion.choices]));
    let idList: string[] = [];
    for (let i = 0; i < subQuestionArray.length; i++) {
        if (subQuestionArray[i][2].length === 1 && subQuestionArray[i][2][0] === "") { // single blank
            idList.push(subQuestionArray[i][0].toString());
            continue;
        }
        
        for (let j = 0; j < subQuestionArray[i][2].length; j++) {
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

    /*const {value: targetTime, removeValue: removeTargetTime} = usePersistState(new Date(Date.now() + 1000 * 100).toString(), "targetTime");
    useEffect(() => {
        return () => {removeTargetTime()}
    }, [removeTargetTime])*/
    const [targetTime] = useState(new Date(Date.now() + 1000 * 100).toString());

    let instructionsInfo : instructionType = {
        title: params.exam_id!,
        instructions: "",
    }

    const questions = questionList.map((question) => {
        return <Question key={`Q${question.headerId}`} questionData={question} />
    })

    const removeAllLocalStorage = () => {
        idList.forEach((item)=> {
            localStorage.removeItem(item);
        })
    }

    const manualSubmitExam = () => {
        setConfirmShow(false);
        setAckShow(true);
        setInTest(false);
        const studentAnswer = prepareAnswer(questionList);
        console.log(studentAnswer);
        //downloadFile(params.exam_id!, JSON.stringify(studentAnswer));
        removeAllLocalStorage();
    }

    const timeoutSubmitExam = () => {
        setTimeoutShow(true);
        setInTest(false);
        const studentAnswer = prepareAnswer(questionList);
        console.log(studentAnswer);
        //downloadFile(params.exam_id!, JSON.stringify(studentAnswer));
        removeAllLocalStorage();
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
                    <CountdownTimer targetTime={targetTime} active={inTest} callback={timeoutSubmitExam} />
                    <div><Button variant="primary" className="my-4 w-50" onClick={() => setConfirmShow(true)}>Submit</Button></div>
                </Col>
            </Row>

            <TimeoutModal onClose={() => {}} show={timeoutShow} toClose={() => setTimeoutShow(false)} />
            <AcknowledgeModal onClose={() => {}} show={ackShow} toClose={() => setAckShow(false)} />
            <ConfirmModal show={confirmShow} onSubmit={manualSubmitExam} onClose={() => setConfirmShow(false)} />
        </AppLayout>
    );
}

export default ExamQuestions;
