import React from 'react';
import { Form } from 'react-bootstrap';

const QuestionLayout = ({index, description, children}: {index: string, description: string, children: React.ReactNode}) => {
    return (
        <Form.Group id={index}>
            <Form.Label>{index}</Form.Label>
            <Form.Text>{description}</Form.Text>
            {children}
        </Form.Group>
    );
}

export default QuestionLayout;
