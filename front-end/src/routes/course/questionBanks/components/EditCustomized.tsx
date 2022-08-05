import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {Button, Col, Form, Row} from 'react-bootstrap';
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import { choiceDataType, subQuestionDataType } from '../../../../components/questionTemplate/subQuestionDataType';

interface blankProps {
    type: 'string' | 'code';
    multiple: boolean;
}

interface graderProps {
    name: string;
    blanks: blankProps[];
}

interface solutionProps {
    solution_idx: number;
    solution_content: string;
}

const Choices = ({id, choicesData}: {id: string, choicesData: choiceDataType[] | null}) => {
    interface choiceProps {
        choice_idx: number;
        choice_content: string;
    }
    
    const [choiceIdx, setChoiceIdx] = useState<number>();
    const [choiceList, setChoiceList] = useState<choiceProps[]>([]);

    useEffect(() => {
        choicesData !== undefined && choicesData !== null &&
            setChoiceIdx(choicesData.length);
        choicesData !== undefined && choicesData !== null &&
            choicesData.map((choiceData, index) =>
            setChoiceList((prevState) => ([
                ...prevState,
                {
                    choice_idx: index,
                    choice_content: choiceData.content
                }
            ]))
        )
    }, [choicesData])

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice.choice_idx !== idx));
    }

    const choices = choiceList.map((choice, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={index}>
                <Col>
                    <div className="my-2">
                        <Form.Control id={id + "_choice" + index}
                            name={id + "_choices"} defaultValue={choice.choice_content}/>
                    </div>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}}
                        onClick={() => deleteChoice(choice.choice_idx)}/>
                </Col>
            </Row>
        );
    });

    return (
        <>
        {choices}
        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, {choice_idx: choiceIdx as number, choice_content: ""}]); setChoiceIdx((choiceIdx as number) + 1);}}>Add Choice</Button>
        </div>
        </>
    );
}

const EditCustomized = ({id, subQuestion, onDelete}: {id: number, subQuestion: subQuestionDataType | null, onDelete: (id: number) => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [description, setDescription] = useState("");

    const [solutionIdx, setSolutionIdx] = useState<number>();
    const [solutionList, setSolutionList] = useState<solutionProps[]>([]);

    useEffect(() => {
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

    const solutions = solutionList.map((solution, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={index}>
                <Col>
                    <div className="my-2">
                        <Form.Control id={"sub" + id + "_solution" + index}
                            name={"sub" + id + "_solutions"} defaultValue={solution.solution_content}/>
                    </div>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}}
                        onClick={() => deleteSolution(solution.solution_idx)}/>
                </Col>
            </Row>
        );
    })

    // const [graders, setGraders] = useState<string[]>([]);
    const [graders, setGraders] = useState<graderProps[]>([]);
    const [grader, setGrader] = useState<graderProps>();

    useEffect(() => {
        subQuestion !== null &&
            setGrader({
                name: subQuestion.grader,
                blanks: subQuestion.blanks,
            })
    }, [subQuestion])

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
            <Form.Control id={"sub" + id + "_description"} defaultValue={subQuestion?.description} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Grader</Form.Label><br/>
            <Form.Text>Choose grader, or upload your grader by clicking "Grader" at the top right corner of the Question Bank page.</Form.Text>
            <Form.Select id={"sub" + id + "_grader"} value={grader?.name} onChange={(e) => getGrader(e.target.value)}>
                <option>Grader Type</option>
                {graders.map((item) => {
                    // if (grader !== "single_blank" && grader !== "single_choice" && grader !== "multiple_choice")
                    // return (<option key={grader} value={grader}>{grader}</option>)
                    if (item.name !== "single_blank" && item.name !== "single_choice" && item.name !== "multiple_choice")
                    return (<option key={item.name} value={item.name}>{item.name}</option>)
                })}
            </Form.Select>
        </Form.Group>

        <Form.Group>
            {
                grader?.blanks.map((blank, index) => {
                    if (blank.multiple) {
                        return (
                            <div key={index}>
                                <Form.Label>{"Blank " + (index + 1) + ": multiple choice"}</Form.Label><br/>
                                <Form.Text>Click "Add Choice" and input all choices content</Form.Text>
                                <Choices id={"sub" + id + "_sub" + index} choicesData={subQuestion?.choices !== undefined ? subQuestion?.choices[index] : null}/>
                            </div>)
                    }

                    // if (blank.type === "string" || blank.type === "code")
                    return (
                        <div key={index}>
                            <Form.Label>
                                {"Blank " + (index + 1) + (blank.type === "string"? ": blank" : ": code")}
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
            <Button variant="primary" onClick={() => {setSolutionList([...solutionList, {solution_idx: solutionIdx as number, solution_content: ""}]); setSolutionIdx((solutionIdx as number) + 1)}}>Add Solution</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default EditCustomized;
