import React, {useCallback, useEffect} from 'react';
import { Row, Col, Card } from 'react-bootstrap';
import AppLayout from "../../components/AppLayout";
import TopNavbar from "../../components/TopNavbar";
import {getBackendApiUrl} from "../../utils/url";
import {secureGet} from "../../utils/axios";
import {useCookies} from "react-cookie";

interface CourseProps {
    name: string;
    semester: string;
    authLevel: string;
}

const Course = (props: CourseProps) => {
    const { name, semester, authLevel } = props
    return (
        <Card className="text-start h-100">
            <Card.Body className="d-flex flex-column">
                <Card.Title className="fs-4 fw-bold">{name}</Card.Title>
                <Card.Text>{semester}</Card.Text>
                <footer className="text-muted mt-auto">{authLevel}</footer>
            </Card.Body>
        </Card>
    )
}

function Dashboard() {
    const listOfCourses = [{
        name: "Distributed Systems",
        semester: "Fall 2022",
        authLevel: "Student"
    }, {
        name: "Advanced Cloud Computing",
        semester: "Fall 2022",
        authLevel: "Student"
    }, {
        name: "Introduction to Computer Systems",
        semester: "Fall 2022",
        authLevel: "Student"
    }]

    const [cookies] = useCookies(['token']);

    const getUsers = useCallback(async () => {
        const url = getBackendApiUrl("/test/users");
        const token = cookies.token;
        console.log(token);
        const result = await secureGet(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result);
    }, [cookies]);

    useEffect(() => {
        getUsers();
    }, [getUsers])

    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1 className="mb-4">My Courses</h1>
                    <Row xs={1} md={2} lg={3} className="g-4">
                        {listOfCourses.map(course => (
                            <Col key={course.name}>
                                <Course {...course}/>
                            </Col>
                        ))}
                    </Row>
                </>
            </AppLayout>
        </div>
    );
}

export default Dashboard;
