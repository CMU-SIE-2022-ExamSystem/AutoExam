import React, {useCallback, useEffect, useState} from 'react';
import {Alert, Button, Col, Modal, Row} from 'react-bootstrap';
import {useNavigate, useParams} from "react-router-dom";
import AppLayout from "../../../components/AppLayout";
import Question from "../../../components/Question";
import CountdownTimer from "../../../components/CountdownTimer";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {blankDataType, choiceDataType} from '../../../components/questionTemplate/subQuestionDataType';
import {downloadFile} from "../../../utils/downloadFile";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import moment from "moment";

const TimeoutModal = ({show, toClose, onClose} :{ show: boolean, toClose: () => void, onClose: () => void }) => {
    return (
        <Modal show={show} onExit={onClose}>
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
        <Modal show={show} onExit={onClose}>
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
                <Button variant="danger" onClick={onClose}>Back</Button>
                <Button variant="primary" onClick={onSubmit}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const FileDownloadModal = ({show, onSubmit, onClose} :{ show: boolean, onSubmit: () => void, onClose: () => void }) => {
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Error</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>There has been a network issue when you are uploading your exam. Don't be worry, your answers have been recorded.</p>
                <p>Send the file to your instructor or TA immediately.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={onSubmit}>Confirm</Button>
           </Modal.Footer>
        </Modal>
    );
}

interface instructionType {
    title: string;
    instructions: string;
    zoom: string;
}

const Instructions = ({info}: {info: instructionType}) => {
    return (
        <div>
            <h1 className="my-3">{info.title}</h1>
            <h2 className="text-start my-4"><strong>Instructions</strong></h2>
            <div className="text-start" dangerouslySetInnerHTML={{__html: info.instructions}} />
            {info.zoom.length > 0 && <Alert key="primary" variant="primary" className="text-start my-4">Zoom Link: {info.zoom}</Alert>}
        </div>
    );
}

interface LooseObject {
    [key: string]: any
}

const prepareAnswer = (qList: questionDataType[]) : object => {
    let result: LooseObject = {};

    function getBlanksAnswer(blanks: blankDataType, questionId: string) {
        let element = document.getElementById(questionId);
        if (element !== null) {
            return (element as HTMLInputElement).value;
        } else {
            console.error(`questionId: ${questionId} not found`);
        }
    }

    function getMCAnswer(choices: choiceDataType[], subQuestionKey: string) {
        let answerList : string[] = [];
        choices.forEach((choice) => {
            let choiceId = subQuestionKey + '_choice' + choice.choice_id;
            let element = document.getElementById(choiceId);
            if (element !== null) {
                let answer = (element as HTMLInputElement).checked;
                if (answer) answerList.push(choice.choice_id);
            }
        })
        return answerList.join("");
    }

    qList.forEach((question: questionDataType, qIdx) => {
        const questionKey = "Q" + (qIdx + 1);
        let subResult: LooseObject = {};
        question.sub_questions.forEach((subQuestion, subQuestionIndex) => {
            const subQuestionKey = questionKey + "_sub" + (subQuestionIndex + 1);
            let answerObject: LooseObject = {};

            for (let i = 0; i < subQuestion.blanks.length; i++) {
                const questionId = subQuestionKey + "_sub" + (i + 1);
                const choices = subQuestion.choices[i];

                if (choices !== null) {
                    answerObject[questionId] = [getMCAnswer(choices, questionId)];
                } else {
                    answerObject[questionId] = [getBlanksAnswer(subQuestion.blanks[i], questionId)];
                }
            }

            subResult[subQuestionKey] = answerObject;
        })

        result[questionKey] = subResult;
    })
    return {
        answers: result
    };
}


const ExamQuestions = () => {
    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();
    const courseName = params.course_name, examId = params.exam_id;
    const [questionList, setQuestionList] = useState<questionDataType[]>([]);
    const [idList, setIdList] = useState<string[]>([]);

    const getQuestionList = useCallback(() => {
        const examUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/question`);
        const token = globalState.token;
        return axios.get(examUrl, {headers: {Authorization: "Bearer " + token}});
    }, [globalState.token]);

    useEffect(() => {
        getQuestionList()
            .then(result => {
                const questionList: questionDataType[] = result.data.data;
                setQuestionList(questionList);
                let subQuestionArray = questionList.flatMap((question, index) =>
                    question.sub_questions.map((subQuestion, subIndex) => ({
                        key: "Q" + (index + 1) + "_sub" + (subIndex + 1),
                        blanks: subQuestion.blanks,
                    })));
                let idList: string[] = [];
                for (let i = 0; i < subQuestionArray.length; i++) {
                    let {key, blanks} = subQuestionArray[i];
                    if (blanks && blanks.length > 0) {
                        for (let j = 1; j <= blanks.length; j++) {
                            idList.push(key + "_sub" + j);
                        }
                    }
                }
                setIdList(idList);
                return idList;
            })
            .then(idList => {
                // Recover

            })
            .catch(badExam => {
                console.error(badExam);
            });
    }, []);

    const [timeoutShow, setTimeoutShow] = useState(false);
    const [ackShow, setAckShow] = useState(false);
    const [confirmShow, setConfirmShow] = useState(false);
    const [fileDownloadShow, setFileDownloadShow] = useState<boolean>(false);

    const [inTest, setInTest] = useState(true);

    const [targetTime, setTargetTime] = useState(new Date(Date.now() + 1000 * 3600).toString());
    const [description, setDescription] = useState<string>("");
    const [zoomLink, setZoomLink] = useState<string>("");

    const getTestGeneralInfo = useCallback(() => {
        const examUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}`);
        const token = globalState.token;
        return axios.get(examUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getTestGeneralInfo()
            .then(result => {
                const testInfo: any = result.data.data;
                let nowMoment = moment();
                if (moment(testInfo.start_at).diff(nowMoment) > 0 || nowMoment.diff(moment(testInfo.end_at)) > 0) {
                    updateGlobalState({alert: {show: true, content: `'${examId}' is not ready.`, variant: 'danger'}});
                    navigate('/courses/' + courseName);
                    return;
                }
                if (!testInfo.can_submit) {
                    updateGlobalState({alert: {show: true, content: `You have used all attempts of '${examId}'.`, variant: 'danger'}});
                    navigate('/courses/' + courseName);
                    return;
                }
                setDescription(testInfo.description);
                setZoomLink(testInfo.zoom);
                setTargetTime(new Date(testInfo.end_at).toString());
            })
    }, [getTestGeneralInfo]);

    let instructionsInfo : instructionType = {
        title: params.exam_id!,
        instructions: description,
        zoom: zoomLink
    }

    const questions = questionList.map((question, index) => {
        return <Question key={`Q${index+1}`} questionData={question} questionId={index + 1}/>
    })

    const removeAllLocalStorage = () => {
        idList.forEach((item)=> {
            localStorage.removeItem(item);
        })
    }

    const autoSaveExam = (examAnswer: object) => {
        const saveUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/answers`);
        const token = globalState.token;
        return axios.put(saveUrl, examAnswer,{headers: {Authorization: "Bearer " + token}})
    }

    const submitExam = () => {
        const submitUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/submit`);
        const token = globalState.token;
        return axios.post(submitUrl, {},{headers: {Authorization: "Bearer " + token}});
    }

    const manualSubmitExam = () => {
        const studentAnswer = prepareAnswer(questionList);
        autoSaveExam(studentAnswer)
            .then(_ => {
                submitExam().then(_ => {
                    setConfirmShow(false);
                    setAckShow(true);
                    setInTest(false);
                })
            })
            .catch(badExam => {
                console.error(badExam);
                downloadFile(globalState.name + "_" + params.exam_id! + ".json", JSON.stringify(studentAnswer));
                setFileDownloadShow(true);
            })

        removeAllLocalStorage();
    }

    const timeoutSubmitExam = () => {
        const studentAnswer = prepareAnswer(questionList);
        autoSaveExam(studentAnswer)
            .then(_ => {
                submitExam().then(_ => {
                    setTimeoutShow(true);
                    setInTest(false);
                })
            })
            .catch(badExam => {
                console.error(badExam);
                downloadFile(globalState.name + "_" + params.exam_id! + ".json", JSON.stringify(studentAnswer));
                setFileDownloadShow(true);
            })

        removeAllLocalStorage();
    }

    useEffect(() => {
        const intervalTimeMS = 60 * 1000;
        const interval = setInterval(() => {
            const studentAnswer = prepareAnswer(questionList);
            autoSaveExam(studentAnswer).catch();
        }, intervalTimeMS);

        return () => clearInterval(interval);
    }, [])

    const ejectFromExam = () => {
        removeAllLocalStorage();
        navigate('/courses/' + courseName);
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
                    {process.env.NODE_ENV === 'development' &&
                        <div>
                            <Button variant="warning" className="my-4 w-50" onClick={() => {console.log(prepareAnswer(questionList));}}>Log Answers</Button>
                            <Button variant="outline-success" className="my-4 w-50" onClick={() => {autoSaveExam(prepareAnswer(questionList)).catch();}}>Manual Save</Button>
                        </div>
                    }
                </Col>
            </Row>

            <TimeoutModal onClose={ejectFromExam} show={timeoutShow} toClose={() => setTimeoutShow(false)} />
            <AcknowledgeModal onClose={ejectFromExam} show={ackShow} toClose={() => setAckShow(false)} />
            <ConfirmModal show={confirmShow} onSubmit={manualSubmitExam} onClose={() => setConfirmShow(false)} />
            <FileDownloadModal show={fileDownloadShow} onClose={ejectFromExam} onSubmit={() => setFileDownloadShow(false)} />
        </AppLayout>
    );
}

export default ExamQuestions;
