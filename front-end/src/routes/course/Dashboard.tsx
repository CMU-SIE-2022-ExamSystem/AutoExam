import React, {useCallback, useEffect} from 'react';
import {Row, Col, Card} from 'react-bootstrap';
import AppLayout from "../../components/AppLayout";
import TopNavbar from "../../components/TopNavbar";
import {getBackendApiUrl} from "../../utils/url";
import {useGlobalState} from "../../components/GlobalStateProvider";
import {LinkContainer} from 'react-router-bootstrap';

interface CourseProps {
    name: string;
    semester: string;
    authLevel: string;
}

const axios = require('axios').default;

const Course = (props: CourseProps) => {
    const {name, semester, authLevel} = props
    return (
        <LinkContainer to={`/courses/${name}`} style={{cursor: "pointer"}}>
            <Card className="text-start h-100">
                <Card.Body className="d-flex flex-column">
                    <Card.Title className="fs-4 fw-bold">{name}</Card.Title>
                    <Card.Text>{semester}</Card.Text>
                    <footer className="text-muted mt-auto">{authLevel}</footer>
                </Card.Body>
            </Card>
        </LinkContainer>
    )
}

function Dashboard() {
    const listOfCourses = [{
        name: "Introduction to Computer Systems",
        semester: "Fall 2022",
        authLevel: "Student"
    }, {
        name: "Advanced Cloud Computing",
        semester: "Fall 2022",
        authLevel: "Student"
    }, {
        name: "Distributed Systems",
        semester: "Fall 2022",
        authLevel: "Student"
    }]
    const {globalState} = useGlobalState();

    const getUsers = useCallback(async () => {
        const url = getBackendApiUrl("/test/users");
        const token = globalState.token;
        console.log(token);
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result);
    }, [globalState]);

    useEffect(() => {
        getUsers();
    }, [globalState, getUsers])

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={null}/>
            </Row>
            <main>
                <h1 className="mt-4">My Courses</h1>
                <Row>
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
