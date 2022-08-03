import React, {useState} from 'react';
import { Button, Form } from 'react-bootstrap';

const AddSingleBlank = ({id, onDelete}: {id: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");
    const [solution, setSolution] = useState("");

    return (
        <>
        <Form.Group>
            <Form.Label><h5>Subquestion (Single Blank)</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Solution</Form.Label>
            <Form.Control id={"sub" + id + "_solution"} onChange={(e) => setSolution(e.target.value)}/>
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="secondary" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddSingleBlank;
