import React from 'react';
import { Alert, Button } from 'react-bootstrap';
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

function ExamInstructions() {
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <div>
                        <h1 className="my-3">Exam 1</h1>
                        <h2 className=" text-start my-4"><strong>Instructions</strong></h2>
                        <p className="text-start">Some detailed instructions.</p>
                        <Alert key="primary" variant="primary" className="text-start my-4">Please turn on your camera to start the exam.</Alert>
                        <Button variant="primary">Start</Button>
                    </div>
                </>
            </AppLayout>
        </div>
    );
}

export default ExamInstructions;
