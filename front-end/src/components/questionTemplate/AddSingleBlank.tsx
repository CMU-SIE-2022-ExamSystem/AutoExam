import React from 'react';
import { Button, Form } from 'react-bootstrap';

const AddSingleBlank = () => {
    return (
        <>
        <Form.Group>
            <Form.Label><h5>Subquestion (Single Blank)</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control></Form.Control>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Answer</Form.Label>
            <Form.Control></Form.Control>
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="secondary">Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddSingleBlank;
