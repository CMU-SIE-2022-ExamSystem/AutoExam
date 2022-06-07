import React from 'react';
import { Row, Col, Card } from 'react-bootstrap';
import AppLayout from "../../components/AppLayout";
import TopNavbar from "../../components/TopNavbar";

interface CourseProps {
    name: string;
    semester: string;
    authLevel: string;
}

const Course = (props: CourseProps) => {
    const { name, semester, authLevel } = props
    return (
        <Card>
            <Card.Body>
                <Card.Title>{name}</Card.Title>
                <Card.Text>{semester}</Card.Text>
                <footer className="text-muted">{authLevel}</footer>
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
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1>My Courses</h1>
                    <Row xs={1} md={2} className="g-4">
                        {listOfCourses.map(course => (
                            <Col>
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
