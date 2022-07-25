import React, { useCallback, useEffect, useMemo, useState } from 'react';
import {Button, Card, Col, Collapse, Form, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from 'axios';
import CollapseQuestion from '../../../components/CollapseQuestion';

const AddTagModal = ({tags, addTag, show, onClose}: {tags: string[], addTag: any, show: boolean, onClose: () => void}) => {
    const [value, setValue] = useState("");
    const handleAdd = () => {
        if (!tags.includes(value)) {
            addTag((prevTags: string[]) => [...prevTags, value]);
            onClose();
            console.log(tags);
        }
    };
    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header closeButton>
                <Modal.Title>Add New Tag</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group>
                        <Form.Control type="text" placeholder="tag" autoFocus
                            onChange={(event) => {setValue(event.target.value)}}/>
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>Cancel</Button>
                <Button variant="primary" onClick={handleAdd}>Add</Button>
            </Modal.Footer>
        </Modal>
    );
}

const QuestionsByTag = ({questionsByTag}: {questionsByTag: questionDataType[]}) => {
    return (
        <Row>
            <Col sm={10}>
                {
                    questionsByTag.map((question) => {
                        return <CollapseQuestion questionData={question}/>
                    })
                }
            </Col>
        </Row>
    );
}

function QuestionBank () {
    const params = useParams();
    const {globalState} = useGlobalState();

    const [tags, setTags] = useState<string[]>(["Integer", "Float", "Cache", "Memory"]);
    const [addTagShow, setAddTagShow] = useState(false);

    const getTags = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setTags(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getTags().catch();
    }, [getTags]);

    const [questions, setQuestions] = useState<questionDataType[]>(require('../exams/questions_new.json').data);
    
    const getQuestions = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setQuestions(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getQuestions().catch();
    }, [getQuestions]);
    
    const updateTagQuestionMap = (questions: questionDataType[]) => {
        let tagQuestionMap = new Map<string, questionDataType[]>();
        for (let i = 0; i < questions.length; i++) {
            const tag = questions[i].questionTag;
            if (!tagQuestionMap.has(tag)) {
                tagQuestionMap.set(tag, [questions[i]]);
            } else {
                (tagQuestionMap.get(tag) as questionDataType[]).push(questions[i]);
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
            <Tab.Container id="questionBank" defaultActiveKey={params.tag || tags[0]}>
                <Row>
                    <Col xs={2} className="d-flex flex-column bg-light vh-100 sticky-top text-start">
                        <Nav variant="pills" className="my-3">
                            {tags.map((tag) => (
                                <Row className="d-flex flex-row align-items-center vw-100">
                                    <Col xs={7}>
                                        <Nav.Item>
                                            <Nav.Link eventKey={tag} href={tag}>{tag}</Nav.Link>
                                        </Nav.Item>
                                    </Col>
                                    <Col xs={5} className="text-end">
                                        <i className="bi-pencil-square" style={{cursor: "pointer"}}/>
                                        <i className="bi-trash mx-2" style={{cursor: "pointer"}}/>
                                    </Col>
                                </Row>
                            ))}
                        </Nav>
                        <i className="bi-plus-square ms-3" style={{cursor: "pointer"}} onClick={() => setAddTagShow(true)}/>
                        <AddTagModal tags={tags} addTag={setTags} show={addTagShow} onClose={() => setAddTagShow(false)}/>
                    </Col>
                    <Col xs={10}>
                        <Row className="text-end">
                            <Link to="#"><Button variant="primary" className='me-4 my-4'>Add Question</Button></Link>
                        </Row>
                        <Tab.Content className="text-start mx-4">
                            {tags.map((tag) => {
                                if (tagQuestionMap.has(tag)) return (
                                    <Tab.Pane eventKey={tag}>
                                        <QuestionsByTag questionsByTag={(tagQuestionMap.get(tag) as questionDataType[])}/>
                                    </Tab.Pane>
                                    );
                                else return (
                                    <Tab.Pane eventKey={tag}>
                                        <p>No questions in this tag!</p>
                                    </Tab.Pane>
                                    );
                            })}
                        </Tab.Content>
                    </Col>
                </Row>
            </Tab.Container>
        </AppLayout>
    );
}

export default QuestionBank;
