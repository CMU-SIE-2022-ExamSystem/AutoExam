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
    is_choice: boolean;
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

const AddQuestionModal = ({show, onClose, tags, getTags, getQuestionsByTag, errorMsg, setErrorMsg} : {show: boolean, onClose: () => void, tags: tagProps[], getTags: () => any, getQuestionsByTag: (tags: tagProps[]) => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [title, setTitle] = useState("");
    const [description, setDescription]= useState("");

    const updateDescription = (newDescription: string) => {
        setDescription(newDescription);
    }

    const [type, setType] = useState("");
    const [id, setId] = useState(0);
    const [subqList, setSubqList] = useState<subqProps[]>([]);

    const deleteSubq = (id: number) => {
        setSubqList(subqList.filter((subq) => subq.id !== id));
    }
    
    const subquestions = subqList.map(({type, id}, index) => {
        if (type === "single_blank") return (<AddSingleBlank key={id} id={id} displayIdx={index + 1} onDelete={deleteSubq}/>);
        if (type === "single_choice" || type === "multiple_choice") return (<AddChoice key={id} type={type} id={id} displayIdx={index + 1} onDelete={deleteSubq}/>);
        if (type === "customized") return (<AddCustomized key={id} id={id} displayIdx={index + 1} onDelete={deleteSubq}/>);
        return (<></>);
    });

    const [graders, setGraders] = useState<graderProps[]>([]);

    const getGraders = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    const getSubquestionsData = () => {
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
            let choices: object[] = []
            let solutions: string[] = []
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
            let choices: object[] = []
            let solutions: string = ""
            choiceNodeList.forEach((item, index) => {
                const choice = item as HTMLInputElement;
                if (choice.checked) {
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

        function getCustomizedData(id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const graderName = (document.getElementById("sub" + id + "_grader") as HTMLInputElement).value;
            const grader = graders.filter((grader) => grader.name === graderName)[0];
            let choices: (object[] | null)[] = [];
            let solutions: string[][] = []
            grader.blanks.forEach((blank: blankProps, index) => {
                if (blank.is_choice) {
                    choices[index] = []
                    solutions[index] = []
                    if (blank.multiple) {
                        solutions[index][0] = ""
                    }
                    const choiceNodeList = document.getElementsByName("sub" + id + "_sub" + index + "_choices");
                    choiceNodeList.forEach((item, choiceIdx) => {
                        const choice = item as HTMLInputElement;
                        const choice_id = String.fromCharCode(choiceIdx + 65);
                        if (choice.checked) {
                            if (blank.multiple) {
                                solutions[index][0] = solutions[index][0].concat(choice_id);
                            } else {
                                solutions[index].push(choice_id);
                            }
                        }
                        const choiceId = "sub" + id + "_sub" + index + "_choice" + choiceIdx;
                        const choice_content = (document.getElementById(choiceId) as HTMLInputElement).value;
                        (choices[index] as object[]).push({choice_id: choice_id, content: choice_content})
                    })
                } else {
                    choices[index] = null
                    solutions[index] = []
                    const solutionNodeList = document.getElementsByName("sub" + id + "_sub" + index + "_solutions");
                    solutionNodeList.forEach((solution) => {
                        solutions[index].push((solution as HTMLInputElement).value);
                    })
                }
            })
            
            const data = {
                grader: graderName,
                description: description,
                choices: choices,
                solutions: solutions
            }
            return data;
        }

        const subqData =  subqList.map(({type, id}) => {
            if (type === "single_blank") return getSingleBlankData(type, id);
            if (type === "single_choice") return getSingleChoiceData(type, id);
            if (type === "multiple_choice") return getMultipleChoiceData(type, id);
            if (type === "customized") return getCustomizedData(id);
            return (<></>);
        });
        return subqData;
    }

    const onSubmit = (e: any) => {
        e.preventDefault();
        const question = {
            description: description,
            question_tag: [...tags].filter((tag) => tag.name === params.tag)[0].id,
            title: title,
            sub_questions: getSubquestionsData()
        }
        addNewQuestion(question);
    }

    const addNewQuestion = async (questionData: object) => {
        console.log(questionData)
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions");
        const token = globalState.token;
        axios.post(url, questionData, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                onClose();
                clearState();
                setErrorMsg("");
                getTags()
                    .then((tags: tagProps[]) => getQuestionsByTag(tags))
                    .catch();
            })
            .catch((error: any) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(typeof response.error.message === "string" ? response.error.message : response.error.message[0].message);
            });
    }

    const clearState = () => {
        setTitle("");
        setDescription("");
        setType("");
        setId(0);
        setSubqList([]);
    }

    return (
        <Modal show={show} onHide={() => {onClose(); clearState(); setErrorMsg("")}} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Add New Question</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Label>Tag: {params.tag}</Form.Label>

                    <Form.Group className="mb-3">
                        <Form.Label>Title</Form.Label>
                        <Form.Control type="text" placeholder="Title" onChange={(e) => setTitle(e.target.value)} required/>
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>Description</Form.Label>
                        <HTMLEditor init={description} update={updateDescription}/>
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
                        <Button variant="primary" onClick={() => {if (type !== "") {setSubqList([...subqList, {type: type, id: id}]); setId(id + 1);}}}>Add Subquestion</Button>
                    </InputGroup>

                    <div><small className="text-danger">{errorMsg}</small></div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearState(); setErrorMsg("")}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default AddQuestionModal;
