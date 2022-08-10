import React, { useState } from 'react';
import { Button, Col, Form, Row } from 'react-bootstrap';

const AddSingleBlank = ({id, displayIdx, onDelete}: {id: number, displayIdx: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");
    
    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<number[]>([]);

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution !== idx));
    }

    const solutions = solutionList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={idx}>
                <Col>
                    <Form.Control id={"sub" + id + "_solution" + index} name={"sub" + id + "_solutions"}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteSolution(idx)}/>
                </Col>
            </Row>
        );
    })

    return (
        <>
        <Form.Group>
            <Form.Label><h5>{displayIdx + ". Single Blank"}</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Solution</Form.Label><br/>
            {solutions}
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setSolutionList([...solutionList, solutionIdx]); setSolutionIdx(solutionIdx + 1)}}>Add Solution</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddSingleBlank;
