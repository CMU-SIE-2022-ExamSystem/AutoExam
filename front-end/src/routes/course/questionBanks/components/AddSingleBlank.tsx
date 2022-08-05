import React, {useState} from 'react';
import { Button, Col, Form, Row } from 'react-bootstrap';

const AddSingleBlank = ({id, onDelete}: {id: number, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");
    const [idx, setIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<number[]>([]);

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution !== idx));
    }

    const solutions = solutionList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={index}>
                <Col>
                    <div className="my-2">
                        <Form.Control id={"sub" + id + "_solution" + index}
                            name={"sub" + id + "_solutions"}/>
                    </div>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}}
                        onClick={() => deleteSolution(idx)}/>
                </Col>
            </Row>
        );
    })

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
            <Form.Label>Solution</Form.Label><br/>
            <Form.Text>Click "Add Solution" and iuput all possible solutions.</Form.Text><br/>
            {solutions}
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setSolutionList([...solutionList, idx]); setIdx(idx + 1)}}>Add Solution</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddSingleBlank;
