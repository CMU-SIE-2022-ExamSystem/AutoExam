import React, { useState } from 'react';
import { Button, Col, Form, InputGroup, Row } from 'react-bootstrap';

const AddChoice = () => {
    const [choiceIdList, setChoiceIdList] = useState<string[]>([]);

    const handleAddChoice = () => {
        const len = choiceIdList.length;
        setChoiceIdList((prev) => [...prev, String.fromCharCode(len + 65)]);
    }

    const handleDeleteChoice = (id: string) => {
        setChoiceIdList((prev) => {
            const prevCopy = [...prev];
            const index = prevCopy.indexOf(id);
            prevCopy.splice(index, 1);
            for (let i = index; i < prevCopy.length; i++) {
                prevCopy[i] = String.fromCharCode(i + 65);
            }
            console.log(prevCopy);
            return prevCopy;
        })
    }

    const choices = choiceIdList.map((id) => {
        return (
            <Row className="d-flex flex-row align-items-center">
                <Col>
                    <InputGroup className="my-2" id={id}>
                        <InputGroup.Checkbox/>
                        <Form.Control/>
                    </InputGroup>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}}
                        onClick={() => handleDeleteChoice(id)}/>
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
            <Form.Control/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Choices</Form.Label><br/>
            <Form.Text>choose all possible answers</Form.Text>
        </Form.Group>

        <div>{choices}</div>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={handleAddChoice}>Add Choice</Button>
            <Button variant="secondary" className="ms-2">Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddChoice;
