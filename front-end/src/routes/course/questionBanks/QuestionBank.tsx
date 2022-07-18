import React, { useCallback, useEffect, useState } from 'react';
import {Button, Col, Form, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import Question from "../../../components/Question";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from 'axios';

const AddTagModal = ({tags, addTag, show, onClose}: {tags: string[], addTag: any, show: boolean, onClose: () => void}) => {
    const [value, setValue] = useState("");
    const handleAdd = () => {
        if (!tags.includes(value)) {
            addTag((prevTags: string[]) => [...prevTags, value]);
            onClose();
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

const QuestionsByTag = ({tag}: {tag: string}) => {
    let questionList: questionDataType[];
    questionList = require('../exams/questions_new.json').data;

    let tagQuestionMap = new Map<string, questionDataType[]>();
    for (let i = 0; i < questionList.length; i++) {
        const tag = questionList[i].questionTag;
        if (!tagQuestionMap.has(tag)) {
            tagQuestionMap.set(tag, [questionList[i]]);
        } else {
            (tagQuestionMap.get(tag) as questionDataType[]).push(questionList[i]);
        }
    }

    if (!tagQuestionMap.has(tag)) return (<p>No questions in this tag!</p>);

    const questionsByTag = tagQuestionMap.get(tag);
    return (
        <Row>
            <Col sm={10}>
                {
                    (questionsByTag as questionDataType[]).map((question) => (
                        <Question key={`Q${question.headerId}`} questionData={question} />
                    ))
                }
            </Col>
        </Row>
    );
}

function QuestionBank () {
    let params = useParams();

    const [tags, setTags] = useState<string[]>(["Integer", "Float", "Cache", "Memory"]);

    const [addTagShow, setAddTagShow] = useState(false);

    const {globalState} = useGlobalState();

    const getTags = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result);
        setTags(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getTags().catch();
    }, [getTags]);

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <Tab.Container id="questionBankSidebar" defaultActiveKey={tags[0]}>
                <Row>
                    <Col sm={2} className="bg-light vh-100 sticky-top">
                        <Nav className="flex-column my-3">
                            {tags.map((tag) => (
                                <Nav.Item>
                                    <Nav.Link eventKey={tag} href="#">
                                        {tag}
                                    </Nav.Link>
                                </Nav.Item>
                            ))}
                        </Nav>
                        <Link to="#"><Button variant="primary" onClick={() => setAddTagShow(true)}>Add Tag</Button></Link>
                        <AddTagModal tags={tags} addTag={setTags} show={addTagShow} onClose={() => setAddTagShow(false)}/>
                    </Col>
                    <Col sm={10}>
                        <Row className="text-end">
                            <Link to="#"><Button variant="primary" className='me-4 my-4'>Add Question</Button></Link>                                
                        </Row>
                        <Tab.Content className="text-start mx-4">
                            {tags.map((tag) => (
                                <Tab.Pane eventKey={tag}>
                                    <QuestionsByTag tag={tag}/>
                                </Tab.Pane>                           
                            ))}
                        </Tab.Content>
                    </Col>
                </Row>
            </Tab.Container>
        </AppLayout>
    );
}

export default QuestionBank;
