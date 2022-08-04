import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {Button, Col, Form, InputGroup, Row} from 'react-bootstrap';
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';

interface blankProps {
    type: 'string' | 'code';
    multiple: boolean;
}

interface graderProps {
    name: string;
    blanks: blankProps[];
}

const Choices = ({id}: {id: string}) => {
    const [choiceIdx, setChoiceIdx] = useState(0);
    const [choiceList, setChoiceList] = useState<number[]>([]);

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice !== idx));
    }

    const choices = choiceList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={idx}>
                <Col>
                    <div className="my-2">
                        <Form.Control id={id + "_choice" + index}
                            name={id + "_choices"}/>
                    </div>
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
        {choices}
        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, choiceIdx]); setChoiceIdx(choiceIdx + 1);}}>Add Choice</Button>
        </div>
        </>
    );
}

const AddCustomized = ({id, onDelete}: {id: number, onDelete: (id: number) => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [description, setDescription] = useState("");

    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<number[]>([]);

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution !== idx));
    }

    const solutions = solutionList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={idx}>
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

    // const [graders, setGraders] = useState<string[]>([]);
    const [graders, setGraders] = useState<graderProps[]>([]);
    const [grader, setGrader] = useState<graderProps>();

    const getGraders = useCallback(async () => {
        // const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/list");
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data)
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    const getGrader = useCallback(async (name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name);
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data)
        setGrader(result.data.data);
    }, [globalState.token, params.course_name])

    return (
        <>
        <Form.Group>
            <Form.Label><h5>Subquestion (Customized)</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Grader</Form.Label><br/>
            <Form.Text>Choose grader, or upload your grader by clicking "Grader" at the top right corner of the Question Bank page.</Form.Text>
            <Form.Select id={"sub" + id + "_grader"} onChange={(e) => getGrader(e.target.value)}>
                <option>Grader Type</option>
                {graders.map((grader) => {
                    // if (grader !== "single_blank" && grader !== "single_choice" && grader !== "multiple_choice")
                    // return (<option key={grader} value={grader}>{grader}</option>)
                    if (grader.name !== "single_blank" && grader.name !== "single_choice" && grader.name !== "multiple_choice")
                    return (<option key={grader.name} value={grader.name}>{grader.name}</option>)
                })}
            </Form.Select>
        </Form.Group>

        <Form.Group>
            {
                grader?.blanks.map((blank, index) => {
                    if (blank.multiple) {
                        return (
                            <div key={index}>
                                <Form.Label>{"Sub " + (index + 1) + ": Multiple Choice"}</Form.Label><br/>
                                <Form.Text>Click "Add Choice" and input all choices content</Form.Text>
                                <Choices id={"sub" + id + "_sub" + index}/>
                            </div>)
                    }

                    // if (blank.type === "string" || blank.type === "code")
                    return (
                        <div key={index}>
                            <Form.Label>
                                {"Sub " + (index + 1) + (blank.type === "string"? ": Blank" : ": Code")}
                            </Form.Label>
                        <br/></div>
                    )
                })
            }
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Solution</Form.Label><br/>
            <Form.Text>Click "Add Solution" and iuput all possible solutions.</Form.Text><br/>
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

export default AddCustomized;
