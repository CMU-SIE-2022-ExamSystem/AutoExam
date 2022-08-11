import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {Button, Col, Form, InputGroup, Row} from 'react-bootstrap';
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import {choiceDataType, subQuestionDataType} from '../../../../components/questionTemplate/subQuestionDataType';
import CodeEditor from "@uiw/react-textarea-code-editor";

interface blankProps {
    type: 'string' | 'code';
    is_choice: boolean;
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

interface choiceProps {
    choice_idx: number;
    choice_content: string;
    choice_checked: boolean;
}

const Blank = ({subSubId, type, solutionData}: {subSubId: string, type: string, solutionData: string[] | undefined}) => {
    const [solutionIdx, setSolutionIdx] = useState(0);
    const [solutionList, setSolutionList] = useState<solutionProps[]>([]);

    const clearState = () => {
        setSolutionIdx(0);
        setSolutionList([]);
    }

    useEffect(() => {
        clearState();
        solutionData !== undefined &&
            setSolutionIdx(solutionData.length);
        solutionData !== undefined &&
            solutionData.map((solution, index) =>
                setSolutionList((prevState) => ([
                    ...prevState,
                    {
                        solution_idx: index,
                        solution_content: solution
                    }
                ]))
            );
    }, [solutionData])

    const deleteSolution = (idx: number) => {
        setSolutionList(solutionList.filter((solution) => solution.solution_idx !== idx));
    }

    const solutions = solutionList.map((solution, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={solution.solution_idx}>
                <Col>
                {type==="string" ?
                    <Form.Control id={subSubId + "_solution" + index} name={subSubId + "_solutions"} defaultValue={solution.solution_content}/> :
                    <CodeEditor
                        id={subSubId + "_solution" + index}
                        name={subSubId + "_solutions"}
                        language={"c"}
                        value={solution.solution_content}
                        className="mb-3"
                        padding={10}
                        style={{
                            height: "200px",
                            fontSize: 12,
                            backgroundColor: "#f5f5f5",
                            fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
                        }}
                    />
                }
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteSolution(solution.solution_idx)}/>
                </Col>
            </Row>
        );
    })

    return (
        <div className="mb-3">
            {solutions}
            <div className="text-end">
                <Button variant="primary" onClick={() => {setSolutionList([...solutionList, {solution_idx: solutionIdx, solution_content: ""}]); setSolutionIdx(solutionIdx + 1)}}>Add Solution</Button>
            </div>
        </div>
    )
}

const Choice = ({subSubId, multiple, choiceData, solutionData}: {subSubId: string, multiple: boolean, choiceData: choiceDataType[] | undefined | null, solutionData: string[] | undefined}) => {    
    const [choiceIdx, setChoiceIdx] = useState(0);
    const [choiceList, setChoiceList] = useState<choiceProps[]>([]);

    const clearState = () => {
        setChoiceIdx(0);
        setChoiceList([]);
    }

    useEffect(() => {
        clearState();
        choiceData !== undefined && choiceData !== null &&
            setChoiceIdx(choiceData.length);
        choiceData !== undefined && choiceData !== null && solutionData !== undefined &&
            choiceData.map((choiceData, index) =>
            setChoiceList((prevState) => ([
                ...prevState,
                {
                    choice_idx: index,
                    choice_content: choiceData.content,
                    choice_checked: multiple ? solutionData[0].includes(choiceData.choice_id) : solutionData.includes(choiceData.choice_id)
                }
            ]))
        )
    }, [choiceData, solutionData, multiple])

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice.choice_idx !== idx));
    }

    const choices = choiceList.map((choice, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={choice.choice_idx}>
                <Col>
                    <InputGroup>
                        <InputGroup.Checkbox name={subSubId + "_choices"} defaultChecked={choice.choice_checked}/>
                        <Form.Control id={subSubId + "_choice" + index} defaultValue={choice.choice_content}/>
                    </InputGroup>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteChoice(choice.choice_idx)}/>
                </Col>
            </Row>
        );
    });

    return (
        <>
        {choices}
        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, {choice_idx: choiceIdx, choice_content: "", choice_checked: false}]); setChoiceIdx(choiceIdx + 1);}}>Add Choice</Button>
        </div>
        </>
    );
}

const EditCustomized = ({id, displayIdx, subQuestion, onDelete}: {id: number, displayIdx: number, subQuestion: subQuestionDataType | null, onDelete: (id: number) => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [description, setDescription] = useState("");

    const [graders, setGraders] = useState<string[]>([]);
    const [grader, setGrader] = useState<graderProps>();

    useEffect(() => {
        subQuestion !== null &&
            setDescription(subQuestion.description);
        subQuestion !== null &&
            setGrader({
                name: subQuestion.grader,
                blanks: subQuestion.blanks,
            })
    }, [subQuestion])

    const getGraders = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/list");
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
            <Form.Label><h5>{displayIdx + ". Customized"}</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} defaultValue={subQuestion?.description} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Grader</Form.Label><br/>
            <Form.Text>Choose grader, or upload your grader by clicking "Grader" at the top right corner of the Question Bank page.</Form.Text>
            <Form.Select id={"sub" + id + "_grader"} value={grader?.name} onChange={(e) => {getGrader(e.target.value);}}>
                <option>Grader Type</option>
                {graders.map((grader) => {
                    if (grader !== "single_blank" && grader !== "single_choice" && grader !== "multiple_choice")
                    return (<option key={grader} value={grader}>{grader}</option>)
                })}
            </Form.Select>
        </Form.Group>

        <Form.Group>
            {grader?.blanks.map((blank, index) => {
                if (blank.is_choice) {
                    return (
                        <div key={index}>
                            <Form.Label>{"(" + (index + 1) + (blank.multiple === true ? ") Multiple" : ") Single") + " Choice"}</Form.Label>
                            <br/>
                            <Choice subSubId={"sub" + id + "_sub" + index} multiple={blank.multiple}
                                choiceData={subQuestion?.choices[index]}
                                solutionData={subQuestion?.solutions[index]}/>
                        </div>
                    )
                } else {
                    return (
                        <div key={index}>
                            <Form.Label>{"(" + (index + 1) + (blank.type === "string"? ") Blank" : ") Code")}</Form.Label>
                            <Blank subSubId={"sub" + id + "_sub" + index} type={blank.type} solutionData={subQuestion?.solutions[index]}/>
                            <br/>
                        </div>
                    )
                }
            })}
        </Form.Group>

        <div className="mb-3 text-end">
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default EditCustomized;
