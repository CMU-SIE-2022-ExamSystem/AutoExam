import React from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Col, Nav, Row, Tab} from "react-bootstrap";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import ExamResultStats from "./ExamResultStats";

const ExamResults = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={courseName} brandLink={"/courses/"+courseName}/>
            </Row>
            <Tab.Container id="TabContainer" defaultActiveKey="statistics">
                <Row>
                    <Col xs={12}>
                        <Nav variant="pills" className="flex-row justify-content-evenly">
                            <Nav.Item>
                                <Nav.Link eventKey="statistics" href="#">
                                    Statistics
                                </Nav.Link>
                            </Nav.Item>
                            <Nav.Item>
                                <Nav.Link eventKey="questions" href="#">
                                    Your Work
                                </Nav.Link>
                            </Nav.Item>
                        </Nav>
                    </Col>
                    <Col xs={12}>
                        <Tab.Content>
                            <Tab.Pane eventKey="statistics">
                                <ExamResultStats />
                            </Tab.Pane>
                            <Tab.Pane eventKey="questions">
                                Questions
                            </Tab.Pane>
                        </Tab.Content>
                    </Col>
                </Row>
            </Tab.Container>
        </AppLayout>
    )
}

export default ExamResults;
