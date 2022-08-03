import React from 'react';
import {Badge, Form} from 'react-bootstrap';

const QuestionLayout = ({questionId, displayIdx, description, score, children}: {questionId: string, displayIdx: number, description: string, score: number, children: React.ReactNode}) => {
    const scoreBadge = (<Badge bg="success ms-1" id={questionId + "_score"}>{score} points</Badge>);
    return (
        <Form.Group key={questionId} id={questionId} className="border-top py-4">
            <Form.Label>{displayIdx + '. '} {scoreBadge}</Form.Label>
            <div dangerouslySetInnerHTML={{__html: description}} />
            {children}
        </Form.Group>
    );
}

export default QuestionLayout;
