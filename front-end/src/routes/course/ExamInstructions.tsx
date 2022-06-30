import React from 'react';
import {Alert, Button, Row} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

function ExamInstructions() {
    let params = useParams();
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name}/>
            </Row>
            <main>
                <div>
                    <h1 className="my-3">{params.exam_id}</h1>
                    <h2 className="text-start my-4"><strong>Instructions</strong></h2>
                    <p className="text-start">Some detailed instructions.</p>
                    <Alert key="primary" variant="primary" className="text-start my-4">Please turn on your camera to
                        start the exam.</Alert>
                    <Link to="questions">
                        <Button type="button" className="btn btn-primary">Start</Button>
                    </Link>
                </div>
            </main>
        </AppLayout>
    );
}

export default ExamInstructions;
