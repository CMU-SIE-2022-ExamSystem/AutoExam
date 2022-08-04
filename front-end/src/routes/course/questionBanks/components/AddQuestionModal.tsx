import React, {useCallback, useEffect, useState} from 'react';
import {Button, Form, InputGroup, Modal} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import HTMLEditor from "../../../../components/HTMLEditor";
import AddSingleBlank from './AddSingleBlank';
import AddChoice from './AddChoice';
import AddCustomized from './AddCustomized';

interface blankProps {
    type: 'string' | 'code';
    multiple: boolean;
}

interface graderProps {
    name: string;
    blanks: blankProps[];
}

interface subqProps {
    id: number;
    type: string;
}

interface tagProps {
    id: string;
    name: string;
}

const AddQuestionModal = ({tag, show, errorMessage, onAdd, onClose, clearMessage} : {tag: tagProps, show: boolean, errorMessage: string, onAdd: (questionData: object) => void, onClose: () => void, clearMessage: () => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [title, setTitle] = useState("");
    const [description, setDescription]= useState<string>("");

    const updateDescription = (newDescription: string) => {
        setDescription(newDescription);
    }

    const [type, setType] = useState("");
    const [id, setId] = useState(0);
    const [subqList, setSubqList] = useState<subqProps[]>([]);

    const deleteSubq = (id: number) => {
        setSubqList(subqList.filter((subq) => subq.id !== id));
    }
    
    const subquestions = (subqList).map(({type, id}) => {
        if (type === "single_blank") return (<AddSingleBlank key={id} id={id} onDelete={deleteSubq}/>);
        if (type === "single_choice") return (<AddChoice key={id} type={type} id={id} onDelete={deleteSubq}/>);
        if (type === "multiple_choice") return (<AddChoice key={id} type={type} id={id} onDelete={deleteSubq}/>);
        if (type === "customized") return (<AddCustomized key={id} id={id} onDelete={deleteSubq}/>);
        return (<></>);
    });

    const [grader, setGrader] = useState<graderProps>();

    const getGrader = async (name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name);
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGrader(result.data.data)
    }

    const getSubquestions = () => {
        function getSingleBlankData(type: string, id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const solutionNodeList = document.getElementsByName("sub" + id + "_solutions");
            let solutions: string[] = []
            solutionNodeList.forEach((solution) => {
                solutions.push((solution as HTMLInputElement).value);
            })

            const data = {
                grader: type,
                description: description,
                choices: [null],
                solutions: [solutions]
            }
            return data;
        }

        function getSingleChoiceData(type: string, id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const choiceNodeList = document.getElementsByName("sub" + id + "_choices");
            let solutions: string[] = []
            let choices: object[] = []
            choiceNodeList.forEach((item, index) => {
                const isChecked = item as HTMLInputElement;
                if (isChecked.checked) {
                    solutions.push(String.fromCharCode(index + 65));
                }
                const choiceId = "sub" + id + "_choice" + index;
                const choiceContent = (document.getElementById(choiceId) as HTMLInputElement).value;
                choices.push({choice_id: String.fromCharCode(index + 65), content: choiceContent})
            })

            const data = {
                grader: type,
                description: description,
                choices: [choices],
                solutions: [solutions]
            }
            return data;
        }

        function getMultipleChoiceData(type: string, id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const choiceNodeList = document.getElementsByName("sub" + id + "_choices");
            let solutions: string = ""
            let choices: object[] = []
            choiceNodeList.forEach((item, index) => {
                const isChecked = item as HTMLInputElement;
                if (isChecked.checked) {
                    solutions = solutions.concat(String.fromCharCode(index + 65));
                }
                const choiceId = "sub" + id + "_choice" + index;
                const choiceContent = (document.getElementById(choiceId) as HTMLInputElement).value;
                choices.push({choice_id: String.fromCharCode(index + 65), content: choiceContent})
            })

            const data = {
                grader: type,
                description: description,
                choices: [choices],
                solutions: [[solutions]]
            }
            return data;
        }

        function getCustomizedata(type: string, id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const graderName = (document.getElementById("sub" + id + "_grader") as HTMLInputElement).value;
            getGrader(graderName);
            let choices: (object[] | null)[] = [];
            (grader as graderProps).blanks.map((blank: blankProps, index) => {
                if (blank.multiple) {
                    const choiceNodeList = document.getElementsByName("sub" + id + "_sub" + index + "_choices");
                    choiceNodeList.forEach((choice, choiceIdx) => {
                        const choiceId = "sub" + id + "_sub" + index + "_choice" + choiceIdx;
                        const choiceContent = (document.getElementById(choiceId) as HTMLInputElement).value;
                        (choices[index] as object[]).push({choice_id: String.fromCharCode(index + 65), content: choiceContent})
                    })
                } else {
                    choices.push(null);
                }
            })
            const solutionNodeList = document.getElementsByName("sub" + id + "_solutions");
            let solutions: string[] = []
            solutionNodeList.forEach((solution) => {
                solutions.push((solution as HTMLInputElement).value);
            })
            
            const data = {
                grader: graderName,
                description: description,
                choices: choices,
                solutions: [solutions]
            }
            return data;
        }

        const subqData =  subqList.map(({type, id}) => {
            if (type === "single_blank") return getSingleBlankData(type, id);
            if (type === "single_choice") return getSingleChoiceData(type, id);
            if (type === "multiple_choice") return getMultipleChoiceData(type, id);
            if (type === "customized") return getCustomizedata(type, id);
            return (<></>);
        });
        return subqData;
    }

    const onSubmit = (e: any) => {
        e.preventDefault();
        const question = {
            description: description,
            question_tag: tag.id,
            title: title,
            sub_questions: getSubquestions()
        }
        onAdd(question);
    }

    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Add New Question</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Label>Tag: {params.tag}</Form.Label>

                    <Form.Group className="mb-3">
                        <Form.Label>Title </Form.Label>
                        <Form.Control type="text" placeholder="Title" onChange={(e) => setTitle(e.target.value)} required/>
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>Description</Form.Label>
                            <div>
                                <HTMLEditor init={description} update={updateDescription}/>
                            </div>
                    </Form.Group>

                    <div>{subquestions}</div>

                    <InputGroup className="mb-3">
                        <Form.Select onChange={(e) => setType(e.target.value)}>
                            <option>Subquestion Type</option>
                            <option value="single_blank">Single Blank</option>
                            <option value="single_choice">Single Choice</option>
                            <option value="multiple_choice">Multiple Choice</option>
                            <option value="customized">Customized</option>
                        </Form.Select>
                        <Button variant="primary"
                            onClick={() => {if (type !== "") setSubqList([...subqList, {type: type, id: id}]); setId(id + 1);}}>
                            Add Subquestion
                        </Button>
                    </InputGroup>

                    <div>
                        <small className="text-danger">{errorMessage}</small>
                    </div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearMessage()}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default AddQuestionModal;
