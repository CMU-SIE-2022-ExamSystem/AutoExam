import React from 'react';
import {Card} from 'react-bootstrap';

const Question = () => {
    return (
        <>
            <Card className="text-start h-100">
                <Card.Header>Question Title</Card.Header>
                <Card.Body className="d-flex flex-column">
                    <Card.Title className="fs-4 fw-bold">Subquestion No.</Card.Title>
                    <Card.Text>Detailed questions.</Card.Text>
                </Card.Body>
            </Card>
        </>
    );
}

export default Question;
