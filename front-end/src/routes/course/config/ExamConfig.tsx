import React, {useCallback, useEffect} from 'react';
import {Button, Col, Container, Nav, Row, Tab} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import ExamConfigInstructions from "./ExamConfigInstructions";
import ExamConfigGlobal from "./ExamConfigGlobal";
import ExamConfigQuestions from "./ExamConfigQuestions";
import {useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";

function ExamConfig() {
    let params = useParams();
    const {globalState} = useGlobalState();

    const courseName = params.course_name;
    const examId = params.exam_id;

    let {setExamConfigState}  = useConfigStates();

    const getSavedConfig = useCallback(async () =>  {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/assessments/" + examId);
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});

        console.log(result.data.data);
        setExamConfigState(result.data.data);
    }, [globalState.token]);

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
                                </Nav>
                            </Col>
                            <Col xs={{span: "9"}}>
                                <div>
                                    <Tab.Content>
                                        <Tab.Pane eventKey="global">
                                            <ExamConfigGlobal />
                                        </Tab.Pane>
                                        <Tab.Pane eventKey="instructions">
                                            <ExamConfigInstructions />
                                        </Tab.Pane>
                                        <Tab.Pane eventKey="questions">
                                            <ExamConfigQuestions />
                                        </Tab.Pane>
                                    </Tab.Content>
                                    <Container fluid className="text-end">
                                        <Button variant="outline-danger" className="m-1">Back</Button>
                                        <Button variant="outline-secondary" className="m-1">Save</Button>
                                        <Button variant="primary" className="m-1">Confirm</Button>
                                    </Container>
                                </div>
                            </Col>
                        </Row>
                    </Tab.Container>
                </main>
        </AppLayout>
    );
}

export default ExamConfig;
