import React from 'react';
import {Badge, Form} from 'react-bootstrap';

/**
 * The layout of a question
 * @param questionId  The HTML id of the question
 * @param displayIdx  The index of the question
 * @param description The question description contents
 * @param score       The total score of that question
 * @param children    The subquestions
 */
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
