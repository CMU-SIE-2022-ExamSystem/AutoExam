import React, {useCallback, useEffect, useState} from 'react';
import {Alert, Button, Col, Container, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {useNavigate, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import ExamConfigInstructions from "./ExamConfigInstructions";
import ExamConfigGlobal from "./ExamConfigGlobal";
import ExamConfigQuestions from "./ExamConfigQuestions";
import {useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import ExamConfigExport from "./ExamConfigExport";
import {GlobalAlertProperties} from "../../../components/globalAlert";

const BackModal = ({show, onSubmit, onClose} :{ show: boolean, onSubmit: () => void, onClose: () => void }) => {
    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Warning</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>You are returning back to the assessment page, and all unsaved changes will be discarded.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={onSubmit}>Confirm</Button>
                <Button variant="danger" onClick={onClose}>Close</Button>
            </Modal.Footer>
        </Modal>
    );
}

const verifyState = () => {

    
}

function ExamConfig() {
    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    let {examConfigState, setExamConfigState}  = useConfigStates();
    let [dataReady, setDataReady] = useState<boolean>(false);

    const getSavedConfig = useCallback(async () =>  {
        const url = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId);
        const token = globalState.token;
        axios.get(url, {headers: {Authorization: "Bearer " + token}})
            .then(result => {
                setExamConfigState(result.data.data);
                setDataReady(true);
            })
            .catch(response => {
                let alertInfo: GlobalAlertProperties = {variant: "danger", content: "", show: true};
                if (response?.response?.data?.error?.message) {
                    alertInfo.content = response.response.data.error.message;
                }
                updateGlobalState({alert: alertInfo});
                navigate('/courses/' + courseName, {replace: true});
            })
    }, [globalState.token, courseName, examId, setExamConfigState]);

    const postConfig = async() => {
        const url = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId);
        const token = globalState.token;
        if (!examConfigState) return;
        const data = {
            general: examConfigState.general,
            settings: examConfigState.settings
        }
        const result = await axios.put(url, data, {headers: {Authorization: "Bearer " + token}});
        return result.data;
    }

    const [backModalShow, setBackModalShow] = useState(false);
    const backHandler = () => {
        setBackModalShow(false);
        navigate("/courses/" + courseName);
    }

    const [saveAlertShow, setSaveAlertShow] = useState(false);

    const saveHandler = () => {
        postConfig()
            .then(_ => {setSaveAlertShow(true)})
            .catch();
    }

    const submitHandler = () => {
        postConfig()
            .then(_ => {})
            .catch();
    }

    useEffect(() => {
        getSavedConfig().catch();
    }, [getSavedConfig]);

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={courseName} brandLink={"/courses/"+courseName}/>
            </Row>
            <main>
                <Tab.Container id="exam-config-tabs" defaultActiveKey="global">
                    <Row>
                        <Col xs={{span: "3"}}>
                            <Nav variant="pills" className="flex-column p-3">
                                <h3>Exam Config</h3>
                                <h4>{examId}</h4>
                                <hr />
                                <Nav.Item>
                                    <Nav.Link eventKey="global" href="#">
                                        Global Settings
                                    </Nav.Link>
                                </Nav.Item>
                                <Nav.Item>
                                    <Nav.Link eventKey="instructions" href="#">
                                        Instructions
                                    </Nav.Link>
                                </Nav.Item>
                                <Nav.Item>
                                    <Nav.Link eventKey="questions" href="#">
                                        Exam Questions
                                    </Nav.Link>
                                </Nav.Item>
                                <Nav.Item>
                                    <Nav.Link eventKey="export" href="#">
                                        Export
                                    </Nav.Link>
                                </Nav.Item>
                            </Nav>
                        </Col>
                        <Col xs={{span: "9"}}>
                            <div>
                                <Tab.Content>
                                    <Tab.Pane eventKey="global">
                                        <ExamConfigGlobal dataReady={dataReady}/>
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="instructions">
                                        <ExamConfigInstructions />
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="questions">
                                        <ExamConfigQuestions />
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="export">
                                        <ExamConfigExport />
                                    </Tab.Pane>
                                </Tab.Content>
                                <Container fluid className="text-end">
                                    <Button variant="outline-danger" className="m-1" onClick={() => setBackModalShow(true)}>Back</Button>
                                    <Button variant="outline-secondary" className="m-1" onClick={saveHandler}>Save</Button>
                                    <Button variant="primary" className="m-1" onClick={submitHandler}>Confirm</Button>
                                </Container>
                            </div>
                        </Col>
                    </Row>
                </Tab.Container>
                <div className="position-absolute bottom-0 end-0 w-25 p-3">
                    <Alert variant="success" show={saveAlertShow} onClose={() => setSaveAlertShow(false)} dismissible>
                        Your config has been saved.
                    </Alert>
                </div>
            </main>
            <BackModal show={backModalShow} onSubmit={backHandler} onClose={() => setBackModalShow(false)} />
        </AppLayout>
    );
}

export default ExamConfig;
