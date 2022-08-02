import React from 'react';
import {Badge, Form} from 'react-bootstrap';

const QuestionLayout = ({questionId, displayIdx, description, score, children}: {questionId: string, displayIdx: number, description: string, score: number, children: React.ReactNode}) => {
    const scoreBadge = (<Badge bg="success ms-1">{score} points</Badge>);
    return (
        <Form.Group key={questionId} id={questionId} className="border-top py-4">
            <Form.Label>{displayIdx + '. ' + description} {scoreBadge}</Form.Label>
            {children}
        </Form.Group>
    );
}

export default QuestionLayout;
