import React, { useState } from 'react';
import { Button, Col, Form, InputGroup, Row } from 'react-bootstrap';

const AddChoice = ({id, onDelete}: {id: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");

    const [idx, setIdx] = useState(0);
    const [choiceList, setChoiceList] = useState<number[]>([]);

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice !== idx));
    }

    const choices = choiceList.map((idx) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={idx}>
                <Col>
                    <InputGroup className="my-2">
                        <InputGroup.Checkbox/>
                        <Form.Control/>
                    </InputGroup>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}}
                        onClick={() => deleteChoice(idx)}/>
                </Col>
            </Row>
        );
    });

    return (
        <>
        <Form.Group>
            <Form.Label><h5>Subquestion (Multiple Choice)</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Choices</Form.Label><br/>
            <Form.Text>choose all possible answers</Form.Text>
        </Form.Group>

        <div>{choices}</div>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, idx]); setIdx(idx + 1);}}>Add Choice</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddChoice;
