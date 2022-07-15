import React, {useCallback, useEffect, useState} from 'react';
import {Row, Col, Card, Modal, Button, Form} from 'react-bootstrap';
import AppLayout from "../../components/AppLayout";
import TopNavbar from "../../components/TopNavbar";
import {getBackendApiUrl} from "../../utils/url";
import {useGlobalState} from "../../components/GlobalStateProvider";
import {LinkContainer} from 'react-router-bootstrap';

interface CourseProps {
    name: string;
    display_name: string;
    semester: string;
    auth_level: string;
}

const axios = require('axios').default;

const Course = (props: CourseProps) => {
    const {name, semester, display_name, auth_level} = props

    const capitalize = (s: string) : string => s.length > 0 ? s[0].toUpperCase() + s.slice(1) : "";
    return (
        <LinkContainer to={`/courses/${name}`} style={{cursor: "pointer"}}>
            <Card className="text-start h-100">
                <Card.Body className="d-flex flex-column">
                    <Card.Title className="fs-4 fw-bold">{display_name}</Card.Title>
                    <Card.Text>{semester}</Card.Text>
                    <footer className="text-muted mt-auto">{capitalize(auth_level)}</footer>
                </Card.Body>
            </Card>
        </LinkContainer>
    )
}

function Dashboard() {
    const [listOfCourses, setListOfCourses] = useState<CourseProps[]>([]);
    const {globalState} = useGlobalState();

    const getCourses = useCallback(async () => {
        const url = getBackendApiUrl("/user/courses");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result);
        setListOfCourses(result.data.data);
    }, [globalState.token]);

    useEffect(() => {
        getCourses().catch();
    }, [getCourses])

    return (
        <AppLayout>
            <Row>
                <TopNavbar brandLink="/dashboard" />
            </Row>
            <main>
                <h1>My Courses</h1>
                <Row className="mt-4">
                    <Col xs={{span: "10", offset: "1"}}>
                        <Row xs={1} lg={2} xl={3} className="g-4">
                            {listOfCourses.map(course => (
                                <Col key={course.name}>
                                    <Course {...course}/>
                                </Col>
                            ))}
                        </Row>
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default Dashboard;
