import React, {useEffect, useState} from 'react';
import {Button, Col, Form, Row} from 'react-bootstrap';
import {subQuestionDataType} from '../../../../components/questionTemplate/subQuestionDataType';

interface solutionProps {
    solution_idx: number;
    solution_content: string;
}

const EditSingleBlank = ({id, displayIdx, subQuestion, onDelete}: {id: number, displayIdx: number,  subQuestion: subQuestionDataType | null, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");
    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<solutionProps[]>([]);

    const clearState = () => {
        setSolutionIdx(0);
        setSolutionList([]);
    }

    useEffect(() => {
        clearState();
        subQuestion !== null &&
            setDescription(subQuestion.description);
        subQuestion !== null &&
            setSolutionIdx(subQuestion.solutions[0].length);
        subQuestion !== null &&
            subQuestion.solutions[0].map((solution, index) =>
                setSolutionList((prevState) => ([
                    ...prevState,
                    {
                        solution_idx: index,
                        solution_content: solution
                    }
                ]))
            );
    }, [subQuestion])

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution.solution_idx !== idx));
    }

    const solutions = solutionList.map(({solution_idx, solution_content}, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={solution_idx}>
                <Col>
                    <Form.Control id={"sub" + id + "_solution" + index} name={"sub" + id + "_solutions"} defaultValue={solution_content}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteSolution(solution_idx)}/>
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
            <Form.Control id={"sub" + id + "_description"} defaultValue={subQuestion?.description} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Solution</Form.Label><br/>
            {solutions}
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setSolutionList([...solutionList, {solution_idx: solutionIdx, solution_content: ""}]); setSolutionIdx(solutionIdx + 1)}}>Add Solution</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default EditSingleBlank;
