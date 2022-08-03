import React, { useState } from 'react';
import { Button, Col, Form, InputGroup, Row } from 'react-bootstrap';

const AddCustomized = ({id, onDelete}: {id: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");
    const [answerType, setAnswerType] = useState("");
    const [solution, setSolution] = useState("");

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
            <Form.Label><h5>Subquestion (Customized)</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Grader</Form.Label><br/>
            <Form.Text>upload a grader file with .py extension</Form.Text>
            <Form.Control type="file"/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Answer Type</Form.Label>
            <Form.Select onChange={(e) => setAnswerType(e.target.value)}>
                <option>Answer Type</option>
                <option value="multiple-blank">Multiple Blank</option>
                <option value="multiple-choice">Multiple Choice</option>
            </Form.Select>
        </Form.Group>

        {answerType === "multiple-blank" &&
            <>
            <Form.Group className="mb-3">
                <Form.Label>Number of Answer Blanks</Form.Label>
                <Form.Control/>
            </Form.Group>

            <Form.Group className="mb-3">
                <Form.Label>Solution</Form.Label>
                <Form.Control onChange={(e) => setSolution(e.target.value)}/>
            </Form.Group>
            </>
        }

        {answerType === "multiple-choice" &&
            <>
            {choices}
            <div className="mb-3 text-end">
                <Button variant="primary" onClick={() => {setChoiceList([...choiceList, idx]); setIdx(idx + 1);}}>Add Choice</Button>
            </div>
            </>
        }

        <div className="mb-3 text-end">
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddCustomized;
