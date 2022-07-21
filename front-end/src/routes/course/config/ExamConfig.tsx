import React, {useEffect} from 'react';
import {Button, Col, Nav, Row, Tab} from 'react-bootstrap';
import {Link, Outlet, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import ExamConfigInstructions from "./ExamConfigInstructions";

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
                                        Global Settings
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="instructions">
                                        <ExamConfigInstructions />
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="questions">
                                        Questions
                                    </Tab.Pane>
                                </Tab.Content>
                            </div>
                        </Col>
                    </Row>
                </Tab.Container>
            </main>
        </AppLayout>
    );
}

export default ExamConfig;
