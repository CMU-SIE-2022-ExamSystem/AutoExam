import React, { useState } from 'react';
import { Button, Col, Form, InputGroup, Row } from 'react-bootstrap';

const AddChoice = ({type, id, onDelete}: {type: string, id: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");

    const [choiceIdx, setChoiceIdx] = useState(0);
    const [choiceList, setChoiceList] = useState<number[]>([]);

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice !== idx));
    }

    const choices = choiceList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={idx}>
                <Col>
                    <InputGroup>
                        <InputGroup.Checkbox name={"sub" + id + "_choices"}/>
                        <Form.Control id={"sub" + id + "_choice" + index}/>
                    </InputGroup>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteChoice(idx)}/>
                </Col>
            </Row>
        );
    });

    return (
        <>
        <Form.Group>
            <Form.Label><h5>Subquestion ({type === "single_choice" ? "Single Choice" : "Multiple Choice"})</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Choices</Form.Label><br/>
            <Form.Text>{type === "single_choice" ? "choose all possible answers" : "choose all correct answers"}</Form.Text>
        </Form.Group>

        <div>{choices}</div>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, choiceIdx]); setChoiceIdx(choiceIdx + 1);}}>Add Choice</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddChoice;
