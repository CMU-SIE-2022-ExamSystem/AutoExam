import React, {useCallback, useEffect, useMemo, useRef, useState } from 'react';
import {Button, Col, Form, InputGroup, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from 'axios';
import CollapseQuestion from '../../../components/CollapseQuestion';
import HTMLEditor from "../../../components/HTMLEditor";
import AddSingleBlank from '../../../components/questionTemplate/AddSingleBlank';
import AddChoice from '../../../components/questionTemplate/AddChoice';

interface tagProps {
    id: string;
    name: string;
}

const AddTagModal = ({show, errorMessage, onClose, onSubmit, clearMessage}: {show: boolean, errorMessage: string, onClose: () => void, onSubmit: (tag: string) => void, clearMessage: () => void}) => {
    const [value, setValue] = useState("");
    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Add New Tag</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={(event) => {event.preventDefault(); onSubmit(value);}}>
                    <Form.Group className="my-4">
                        <Form.Control type="text" placeholder="Tag" required autoFocus id="new-tag-name"
                            onChange={(event) => {setValue(event.target.value); clearMessage();}}/>
                    </Form.Group>
                    <div>
                        <small className="text-danger">{errorMessage}</small>
                    </div>
                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearMessage()}}>Cancel</Button>
                        <Button variant="primary" type="submit" className="ms-2">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

const AddQuestionModal = ({tag, show, onClose} : {tag: string, show: boolean, onClose: () => void}) => {
    const editorRef = useRef<any>(null);
    const log = (): string => {
        if (editorRef.current) {
            console.log(editorRef.current.getContent());
            return editorRef.current.getContent();
        }
        return "";
    };

    const [type, setType] = useState<string>();
    const [subqList, setSubqList] = useState<string[]>([]);
    
    const subquestions = (subqList as string[]).map((subqType) => {
        if (subqType === "single-blank") return (<AddSingleBlank/>);
        if (subqType === "single-choice") return (<AddChoice/>);
        if (subqType === "multiple-choice") return (<AddChoice/>);
        // if (subqType === "customized") return (<AddCustomizedQuestion/>);
        return(<></>);
    });

    return (
        <Modal show={show} onHide={onClose} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Add new Question</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Label>Tag: {tag}</Form.Label>

                    <Form.Group className="mb-3">
                        <Form.Label>Title </Form.Label>
                        <Form.Control type="text" placeholder="Title" required/>
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>Description</Form.Label>
                            <div>
                                <HTMLEditor editorRef={editorRef}/>
                                {/* <button onClick={log}>Log editor content</button> */}
                            </div>
                    </Form.Group>

                    <div>{subquestions}</div>

                    <InputGroup className="mb-3">
                        <Form.Select value={type} onChange={(e) => setType(e.target.value)}>
                            <option>Subquestion Type</option>
                            <option value="single-blank">Single Blank</option>
                            <option value="single-choice">Single Choice</option>
                            <option value="multiple-choice">Multiple Choice</option>
                            <option value="customized">Customized</option>
                        </Form.Select>
                        <Button variant="primary"
                            onClick={() => {setSubqList((prev: string[]) => ([...prev, type] as string[]))}}>
                            Add Subquestion
                        </Button>
                    </InputGroup>

                    <div className="text-end">
                        <Button variant="secondary" onClick={onClose}>Close</Button>
                        <Button variant="primary" type="submit" className="ms-2">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

const QuestionsByTag = ({questionsByTag}: {questionsByTag: questionDataType[]}) => {
    return (
        <Row>
            <Col sm={10}>
                {
                    questionsByTag.map((question) => {
                        return <CollapseQuestion questionData={question} key={question.id}/>
                    })
                }
            </Col>
        </Row>
    );
}

function QuestionBank () {
    const params = useParams();
    const {globalState} = useGlobalState();

    const [tags, setTags] = useState<tagProps[]>([]);
    const [addTagShow, setAddTagShow] = useState(false);

    const getTags = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data);
        setTags(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getTags().catch();
    }, [getTags]);

    const [invalidAddTag, setInvalidAddTag] = useState("");

    const addNewTag = async (tag: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const data = {
            name: tag
        };
        axios.post(url, data, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setAddTagShow(false);
            })
            .catch((error: any) => {
                let response = error.response.data;
                setInvalidAddTag(response.error.message[0].message);
            });
    };

    const [questions, setQuestions] = useState<questionDataType[]>(require('../exams/questions_new.json').data);
    const [addQuestionShow, setAddQuestionShow] = useState(false);
    
    const getQuestions = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data);
        setQuestions(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getQuestions().catch();
    }, [getQuestions]);
    
    const updateTagQuestionMap = (questions: questionDataType[]) => {
        let tagQuestionMap = new Map<string, questionDataType[]>();
        for (let i = 0; i < questions.length; i++) {
            const tagName = questions[i].question_tag;
            if (!tagQuestionMap.has(tagName)) {
                tagQuestionMap.set(tagName, [questions[i]]);
            } else {
                (tagQuestionMap.get(tagName) as questionDataType[]).push(questions[i]);
            }
        }
        return tagQuestionMap;
    };
    const tagQuestionMap = useMemo(() => updateTagQuestionMap(questions), [questions]);
    
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <Tab.Container id="questionBank" defaultActiveKey={params.tag || tags[0].name}>
                <Row>
                    <Col xs={2} className="d-flex flex-column bg-light vh-100 sticky-top text-start">
                        <Nav variant="pills" className="my-3">
                            {params.tag !== "null" &&
                                tags.map((tag) => (
                                    <Row className="d-flex flex-row align-items-center vw-100">
                                        <Col xs={8}>
                                            <Nav.Item>
                                                <Nav.Link eventKey={tag.name} href={tag.name}>{tag.name}</Nav.Link>
                                            </Nav.Item>
                                        </Col>
                                        <Col xs={4} className="text-end">
                                            <i className="bi-pencil-square" style={{cursor: "pointer"}}/>
                                            <i className="bi-trash mx-2" style={{cursor: "pointer"}}/>
                                        </Col>
                                    </Row>
                                ))
                            }
                        </Nav>
                        <i className="bi-plus-square ms-3" style={{cursor: "pointer"}} onClick={() => setAddTagShow(true)}/>
                    </Col>
                    {params.tag !== "null" &&
                        <Col xs={10}>
                            <Row className="text-end">
                                <Link to="#"><Button variant="primary" className='me-4 my-4' onClick={() => setAddQuestionShow(true)}>Add Question</Button></Link>
                            </Row>
                            <Tab.Content className="text-start mx-4">
                                {tags.map((tag) => {
                                    if (tagQuestionMap.has(tag.name)) return (
                                        <Tab.Pane eventKey={tag.name}>
                                            <QuestionsByTag questionsByTag={(tagQuestionMap.get(tag.name) as questionDataType[])} key={tag.name}/>
                                        </Tab.Pane>
                                        );
                                    else return (
                                        <Tab.Pane eventKey={tag.name}>
                                            <p>No questions in this tag!</p>
                                        </Tab.Pane>
                                        );
                                })}
                            </Tab.Content>
                        </Col>
                    }
                </Row>
                <AddTagModal show={addTagShow}
                    errorMessage={invalidAddTag}
                    onClose={() => setAddTagShow(false)}
                    onSubmit={(tagName) => {addNewTag(tagName)}}
                    clearMessage={() => setInvalidAddTag("")}/>
                <AddQuestionModal tag={(params.tag as string)}
                    show={addQuestionShow}
                    onClose={() => setAddQuestionShow(false)}/>
            </Tab.Container>
        </AppLayout>
    );
}

export default QuestionBank;
