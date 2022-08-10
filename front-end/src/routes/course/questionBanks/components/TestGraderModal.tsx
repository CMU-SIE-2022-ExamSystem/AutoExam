import React, {useState} from 'react';
import {Button, Col, Form, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType, {blankDataType} from './graderDataType';

interface LooseObject {
    [key: string]: any
}

interface solutionProps {
    solution_idx: number;
    content: string;
}

const TestSuccessModal = ({show, onClose, setTestGraderShow, result, grader, getGraders, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, setTestGraderShow: any, result: string, grader: graderDataType, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();

    const validateGrader = async(name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name + "/valid");
        const token = globalState.token;
        const data = {
            valid: true
        }
        axios.put(url, data, {headers: {Authorization: "Bearer " + token}})
            .then(response => {
                setErrorMsg("");
                onClose();
                getGraders();
            })
            .catch((error) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(typeof response.error.message === "string" ? response.error.message : response.error.message[0].message);
            });
    }
    
    return (
        <Modal show={show} onHide={() => {onClose(); setErrorMsg("")}}>
            <Modal.Header>
                <Modal.Title>Test Result</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                {result}<br/><br/>
                <div className="text-primary">Click "Back" to do more tests, or click "Confirm" to validate this grader.</div>
                <div><small className="text-danger">{errorMsg}</small></div>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={() => {onClose(); setTestGraderShow(true);  setErrorMsg("")}}>Back</Button>
                <Button variant="primary" className="ms-2" onClick={() => validateGrader(grader.name)}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    )
}

const OneInBlank = ({blank, blankIdx}: {blank: blankDataType, blankIdx: number}) => {
    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<solutionProps[]>([]);

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution.solution_idx !== idx));
    }

    const solutions = solutionList.map((solution, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={solution.solution_idx}>
                <Col>
                    <Form.Control id={"blank" + blankIdx + "_solution" + index} name={"blank" + blankIdx + "_solutions"}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteSolution(solution.solution_idx)}/>
                </Col>
            </Row>
        );
    })

    return (
        <>
        <Form.Label>{"(" + blankIdx + ") " + (blank.is_choice ? (blank.multiple? "Multiple Choice" : "Single Choice") : (blank.type === "string" ? "Blank" : "Code"))}</Form.Label>
        <br/>
        {blank.is_choice &&
            <Form.Text>It's a choice. You can directly input the answer to test grader.</Form.Text>
        }

        <Form.Group className="mb-3">
            <Form.Label>Answer</Form.Label>
            <Form.Control id={"blank" + blankIdx + "_answer"}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Solution</Form.Label>
            {solutions}
            <div className="text-end">
                <Button variant="primary" onClick={() => {setSolutionList([...solutionList, {solution_idx: solutionIdx, content: ""}]); setSolutionIdx(solutionIdx + 1)}}>Add Solution</Button>
            </div>
        </Form.Group>
        <hr/>
        </>
    )
}

const TestGraderModal = ({show, setTestGraderShow, onClose, grader, getGraders, errorMsg, setErrorMsg}: {show: boolean, setTestGraderShow: any, onClose: () => void, grader: graderDataType, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();

    const [testResult, setTestResult] = useState("");
    const [testSuccessShow, setTestSuccessShow] = useState(false);

    const onSubmit = (e: any) => {
        e.preventDefault();

        function getAnswers() {
            let answerMap = new Map<string, string[]>();
            grader.blanks.forEach((blank, index) => {
                const answer = (document.getElementById("blank" + (index + 1) + "_answer") as HTMLInputElement).value;
                const key = "additionalProp" + (index + 1);
                answerMap.set(key, [answer])
            })

            let answerObj: LooseObject = {};
            answerMap.forEach(function(value: string[], key: string) {
                answerObj[key] = value;
            })

            const data = {
                test_autograder: answerObj
            }
            return data;
        }

        function getSolutions() {
            let solutionMap = new Map<string, string[]>();
            grader.blanks.forEach((blank, index) => {
                let solutions: string[] = []
                const solutionNodeList = document.getElementsByName("blank" + (index + 1) + "_solutions");
                solutionNodeList.forEach((solution, solutionIdx) => {
                    const solutionContent = (document.getElementById("blank" + (index +1) + "_solution" + solutionIdx) as HTMLInputElement).value;
                    solutions.push(solutionContent);
                })
                const key = "additionalProp" + (index + 1);
                solutionMap.set(key, solutions)
            })

            let solutionObj: LooseObject = {};
            solutionMap.forEach(function(value: string[], key: string) {
                solutionObj[key] = value;
            })

            const data = {
                test_autograder: solutionObj
            }
            return data
        }

        const testData = {
            answers: getAnswers(),
            solutions: getSolutions()
        }
        testGrader(grader.name, testData);
    }

    const testGrader = async (name: string, testData: object) => {
        console.log(testData)
        const url = getBackendApiUrl("/courses/" + params.course_name + "/autograder/" + name + "/test");
        const token = globalState.token;
        axios.post(url, testData, {headers: {Authorization: "Bearer " + token}})
            .then(response => {
                setErrorMsg("");
                setTestResult(response.data.data);
                onClose();
                setTestSuccessShow(true);
                getGraders();
            })
            .catch((error) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(typeof response.error.message === "string" ? response.error.message : response.error.message[0].message);
            });
    }
    
    return (
        <>
        <Modal show={show} onHide={() => {onClose(); setErrorMsg("");}}>
            <Modal.Header>
                <Modal.Title>Test Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    {grader !== undefined &&
                    <>
                        <Form.Label>{"Name: " + grader.name}</Form.Label><br/>

                        {
                            grader.blanks.map((blank, index) => (
                                <OneInBlank key={index} blank={blank} blankIdx={index + 1}/>
                            ))
                        }
                    </>
                    }

                    <Form.Group className="mb-3">
                        <Form.Text>Please wait for a moment for test result to be shown...</Form.Text>
                    </Form.Group>

                    <div><small className="text-danger">{errorMsg}</small></div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); setErrorMsg("");}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Confirm</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>

        <TestSuccessModal
            show={testSuccessShow}
            onClose={() => setTestSuccessShow(false)}
            setTestGraderShow={setTestGraderShow}
            result={testResult}
            grader={grader as graderDataType}
            getGraders={getGraders}
            errorMsg={errorMsg}
            setErrorMsg={setErrorMsg}
        />
        </>
    )
}

export default TestGraderModal;
