import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {Button, Col, Form, InputGroup, Row} from 'react-bootstrap';
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';

interface blankProps {
    is_choice: boolean;
    multiple: boolean;
    type: 'string' | 'code';
}

interface graderProps {
    name: string;
    blanks: blankProps[];
}

const Blank = ({subSubId}: {subSubId: string}) => {
    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<number[]>([]);

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution !== idx));
    }

    const solutions = solutionList.map((idx, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={idx}>
                <Col>
                    <Form.Control id={subSubId + "_solution" + index} name={subSubId + "_solutions"}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteSolution(idx)}/>
                </Col>
            </Row>
        );
    })

    return (
        <div>
            {solutions}
            <div className="text-end">
                <Button variant="primary" onClick={() => {setSolutionList([...solutionList, solutionIdx]); setSolutionIdx(solutionIdx + 1)}}>Add Solution</Button>
            </div>
        </div>
    )
}

const Choice = ({subSubId}: {subSubId: string}) => {
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
                        <InputGroup.Checkbox name={subSubId + "_choices"}/>
                        <Form.Control id={subSubId + "_choice" + index}/>
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

    // const [graders, setGraders] = useState<string[]>([]);
    const [graders, setGraders] = useState<graderProps[]>([]);
    const [grader, setGrader] = useState<graderProps>();

    const getGraders = useCallback(async () => {
        // const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/list");
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    const getGrader = useCallback(async (name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name);
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
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
            {grader?.blanks.map((blank, index) => {
                if (blank.is_choice) {
                    return (
                        <div key={index}>
                            <Form.Label>{(index + 1) + (blank.multiple === true ? ". Multiple" : ". Single") + " Choice"}</Form.Label>
                            <br/>
                            <Choice subSubId={"sub" + id + "_sub" + index}/>
                        </div>
                    )
                }

                // blank or code
                return (
                    <div key={index}>
                        <Form.Label>{(index + 1) + (blank.type === "string"? ". Blank" : ". Code")}</Form.Label>
                        <Blank subSubId={"sub" + id + "_sub" + index}/>
                        <br/>
                    </div>
                )
            })
            }
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="secondary" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default AddCustomized;
