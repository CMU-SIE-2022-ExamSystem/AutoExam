import React, {useEffect} from 'react';
import {Button, Col, Container, Nav, Row, Tab} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import ExamConfigInstructions from "./ExamConfigInstructions";
import ExamConfigGlobal from "./ExamConfigGlobal";
import ExamConfigQuestions from "./ExamConfigQuestions";

function ExamConfig() {
    let params = useParams();

    const getSavedConfig = async () =>  {

    }

    useEffect(() => {
        getSavedConfig().catch();
    }, [getSavedConfig]);

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <main>
                <Tab.Container id="exam-config-tabs" defaultActiveKey="global">
                    <Row>
                        <Col xs={{span: "3"}}>
                            <Nav variant="pills" className="flex-column">
                                <div>Exam Config</div>
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
                                    <Button variant="outline-danger">Back</Button>
                                    <Button variant="outline-secondary">Save</Button>
                                    <Button variant="primary">Confirm</Button>
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
