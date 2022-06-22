import React from 'react';
import { Button, Card } from 'react-bootstrap';
import { useParams } from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

const QuestionList = () => {

}

function ExamQuestions() {
    let params = useParams();
    const questionList = QuestionList();
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1 className="my-3">{params.exam_id}</h1>
                    <h2 className="text-start my-4"><strong>Instructions</strong></h2>
                    <p className="text-start">Some detailed instructions.</p>
                    <br/>
                    <Card className="text-start h-100">
                        <Card.Header>Question Title</Card.Header>
                        <Card.Body className="d-flex flex-column">
                            <Card.Title className="fs-4 fw-bold">Subquestion No.</Card.Title>
                            <Card.Text>Detailed questions.</Card.Text>
                        </Card.Body>
                    </Card>
                    <br/>
                    {/* {questionList} */}
                    <div><Button variant="primary" className="my-4">Submit</Button></div>
                </>
            </AppLayout>
        </div>
    );
}

export default ExamQuestions;
