import React, {useCallback, useEffect, useState} from 'react';
import {Button, Form, InputGroup, Modal} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {subQuestionDataType} from '../../../../components/questionTemplate/subQuestionDataType';
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import HTMLEditor from "../../../../components/HTMLEditor";
import questionDataType from '../../../../components/questionTemplate/questionDataType';
import EditSingleBlank from './EditSingleBlank';
import EditChoice from './EditChoice';
import EditCustomized from './EditCustomized';

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
    content: subQuestionDataType | null;
}

interface tagProps {
    id: string;
    name: string;
}

const EditQuestionModal = ({show, tag, question, errorMessage, onEdit, onClose, clearQuestion, clearMessage}: {show: boolean, tag: tagProps, question: questionDataType, errorMessage: string, onEdit: (id: string, data: object) => void, onClose: () => void, clearQuestion: () => void, clearMessage: () => void}) => {
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
    
    useEffect(() => {
        question !== undefined &&
            setId(question.sub_questions.length);
        question !== undefined &&
            question.sub_questions.map((subQuestion, index) =>
                setSubqList((prevState) => ([
                    ...prevState,
                    {
                        id: index,
                        type: subQuestion.grader === "single_blank" || subQuestion.grader === "single_choice" || subQuestion.grader === "multiple_choice" ? subQuestion.grader : "customized",
                        content: subQuestion
                    }
                ]))
            );
    }, [question]);

    const deleteSubq = (id: number) => {
        setSubqList(subqList.filter((subq) => subq.id !== id));
    }
    
    const subquestions = subqList.map(({type, id, content}) => {
        if (type === "single_blank") return (<EditSingleBlank key={id} id={id} subQuestion={content} onDelete={deleteSubq}/>);
        if (type === "single_choice" || type === "multiple_choice") return (<EditChoice key={id} type={type} id={id} subQuestion={content} onDelete={deleteSubq}/>);
        if (type === "customized") return (<EditCustomized key={id} id={id} subQuestion={content} onDelete={deleteSubq}/>);
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

        function getCustomizedData(type: string, id: number) {
            const description = (document.getElementById("sub" + id + "_description") as HTMLInputElement).value;
            const graderName = (document.getElementById("sub" + id + "_grader") as HTMLInputElement).value;
            const grader = graders.filter((grader) => grader.name === graderName)[0];
            let choices: (object[] | null)[] = [];
            grader.blanks.forEach((blank: blankProps, index) => {
                if (blank.multiple) {
                    choices[index] = []
                    const choiceNodeList = document.getElementsByName("sub" + id + "_sub" + index + "_choices");
                    choiceNodeList.forEach((choice, choiceIdx) => {
                        const choiceId = "sub" + id + "_sub" + index + "_choice" + choiceIdx;
                        const choiceContent = (document.getElementById(choiceId) as HTMLInputElement).value;
                        (choices[index] as object[]).push({choice_id: String.fromCharCode(choiceIdx + 65), content: choiceContent})
                    })
                } else {
                    choices[index] = null
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
            if (type === "customized") return getCustomizedData(type, id);
            return (<></>);
        });
        return subqData;
    }

    const onSubmit = (e: any) => {
        e.preventDefault();
        const questionData = {
            description: description,
            question_tag: tag.id,
            title: title,
            sub_questions: getSubquestionsData()
        }
        onEdit(question.id, questionData);
    }

    const clearState = () => {
        setTitle("");
        setDescription("");
        setType("");
        setId(0);
        setSubqList([]);
    }

    return (
        <Modal show={show} onHide={() => {onClose(); clearState(); clearQuestion(); clearMessage()}} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Edit Queston</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Label>Tag: {params.tag}</Form.Label>

                    {question !== undefined &&
                    <>
                        <Form.Group className="mb-3">
                            <Form.Label>Title </Form.Label>
                            <Form.Control type="text" required defaultValue={question.title} onChange={(e) => setTitle(e.target.value)}/>
                        </Form.Group>

                        <Form.Group className="mb-3">
                            <Form.Label>Description</Form.Label>
                            <HTMLEditor init={question.description} update={updateDescription}/>
                        </Form.Group>

                        <div>{subquestions}</div>
                    </>
                    }

                    <InputGroup className="mb-3">
                        <Form.Select onChange={(e) => setType(e.target.value)}>
                            <option>Subquestion Type</option>
                            <option value="single_blank">Single Blank</option>
                            <option value="single_choice">Single Choice</option>
                            <option value="multiple_choice">Multiple Choice</option>
                            <option value="customized">Customized</option>
                        </Form.Select>
                        <Button variant="primary" onClick={() => {if (type !== "") setSubqList([...subqList, {type: type, id : id as number, content: null}]); setId((id as number) + 1)}}>Add Subquestion</Button>
                    </InputGroup>

                    <div>
                        <small className="text-danger">{errorMessage}</small>
                    </div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearQuestion(); clearMessage()}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Confirm</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default EditQuestionModal;
