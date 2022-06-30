import React from 'react';
import { Form } from 'react-bootstrap';

const QuestionLayout = ({questionId, description, children}: {questionId: string, description: string, children: React.ReactNode}) => {
    return (
        <Form.Group id={questionId} className="border-top py-4">
            <Form.Label>{questionId + '. ' + description}</Form.Label>
            {children}
        </Form.Group>
    );
}

export default QuestionLayout;
