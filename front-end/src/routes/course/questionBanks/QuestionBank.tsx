import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, Form, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from 'axios';
import CollapseQuestion from './components/CollapseQuestion';
import AddQuestionModal from './components/AddQuestionModal';
import EditQuestionModal from './components/EditQuestionModal';
import GraderModal from './components/GraderModal';

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

const EditTagModal = ({show, tag, errorMessage, onClose, onSubmit, clearMessage}: {show: boolean, tag: tagProps, errorMessage: string, onClose: () => void, onSubmit: (id: string, name: string) => void, clearMessage: () => void}) => {
    const [value, setValue] = useState("");
    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Edit Tag</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={(event) => {event.preventDefault(); onSubmit(tag.id, value);}}>
                    <Form.Group className="my-4">
                        <Form.Control type="text" placeholder="New Tag Name" required autoFocus id="edit-tag-name"
                            defaultValue={tag !== undefined ? tag.name : ""}
                            onChange={(event) => {setValue(event.target.value); clearMessage();}}/>
                    </Form.Group>
                    <div>
                        <small className="text-danger">{errorMessage}</small>
                    </div>
                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearMessage()}}>Cancel</Button>
                        <Button variant="primary" type="submit" className="ms-2">Confirm</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

const DeleteTagModal = ({show, tag, errorMessage, onClose, onSubmit, clearMessage}: {show: boolean, tag: tagProps, errorMessage: string, onClose: () => void, onSubmit: (id: string) => void, clearMessage: () => void}) => {
    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Delete Tag</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {"Do you want to delete tag \"" + (tag !== undefined? tag.name : "") + "\"?"}
                <div>
                    <small className="text-danger">{errorMessage}</small>
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => {onClose(); clearMessage()}}>Cancel</Button>
                <Button variant="primary" type="submit" className="ms-2" onClick={() => onSubmit(tag.id)}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const DeleteQuestionModal = ({show, onClose, tags, getTags, getQuestionsByTag, question, clearQuestion, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, tags: tagProps[], getTags: () => any, getQuestionsByTag: (tags: tagProps[]) => void, question: questionDataType, clearQuestion: () => void, errorMsg: string,  setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const deleteQuestion = async (id: string, hard: boolean) => {
        let url: string = "";
        if (hard) {
            url = getBackendApiUrl("/courses/" + params.course_name + "/questions/" + id + "?hard=true");
        } else {
            url = getBackendApiUrl("/courses/" + params.course_name + "/questions/" + id);
        }
        const token = globalState.token;
        axios.delete(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                onClose();
                clearQuestion();
                setErrorMsg("");
                getTags()
                    .then((tags: tagProps[]) => getQuestionsByTag(tags))
                    .catch();
            })
            .catch((error: any) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(response.error.message || response.error.message[0].message);
            });
    }

    return (
        <Modal show={show} onHide={() => {onClose(); clearQuestion(); setErrorMsg("")}}>
            <Modal.Header closeButton>
                <Modal.Title>Delete Queston</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {
                    question !== undefined &&
                        (question.hidden? "Do you want to delete this question permanently?" : "Do you want to delete this question?")
                }
                <div>
                    <small className="text-danger">{errorMsg}</small>
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => {onClose(); clearQuestion(); setErrorMsg("")}}>Cancel</Button>
                <Button variant="primary" type="submit" className="ms-2" onClick={() => deleteQuestion(question.id, question.hidden? true : false)}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const QuestionsByTag = ({questions, setQuestion, setDeleteShow, setEditShow}: {questions: questionDataType[] | undefined, setQuestion: any, setDeleteShow: any, setEditShow: any}) => {
    return (
        <Row>
            <Col sm={10}>
                {!!questions &&
                    questions.map((question) => {
                        return <CollapseQuestion question={question} key={question.id}
                            setQuestion={setQuestion} setDeleteShow ={setDeleteShow} setEditShow={setEditShow}/>
                    })
                }
                {!questions &&
                    "There are no questions under this tag."
                }
            </Col>
        </Row>
    );
}

function QuestionBank () {
    const params = useParams();
    const {globalState} = useGlobalState();

    const [tags, setTags] = useState<tagProps[]>([]);
    const [tag, setTag] = useState<tagProps>();
    const [addTagShow, setAddTagShow] = useState(false);
    const [editTagShow, setEditTagShow] = useState(false);
    const [deleteTagShow, setDeleteTagShow] = useState(false);
    const [tagError, setTagError] = useState("");

    const getTags = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data);
        if (result.data.data) setTags(result.data.data);
        return result.data.data;
    }, [globalState.token, params.course_name]);

    const addNewTag = async (name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags");
        const token = globalState.token;
        const data = {
            name: name
        };
        axios.post(url, data, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setAddTagShow(false);
                getTags();
            })
            .catch((error: any) => {
                console.log(error)
                let response = error.response.data;
                setTagError(response.error.message[0].message);
            });
    };

    const editTag = async (id: string, name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags/" + id);
        const token = globalState.token;
        const data = {
            name: name,
        };
        axios.put(url, data, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setEditTagShow(false);
                getTags();
            })
            .catch((error: any) => {
                let response = error.response.data;
                console.log(error);
                setTagError(response.error.message[0].message);
            })
    }

    const deleteTag = async (id: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/tags/" + id);
        const token = globalState.token;
        axios.delete(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setDeleteTagShow(false);
                getTags();
            })
            .catch((error: any) => {
                let response = error.response.data;
                console.log(error);
                setTagError(response.error.message);
            })
    }

    const [questionsByTag, setQuestionsByTag] = useState<questionDataType[]>();
    const [question, setQuestion] = useState<questionDataType>();
    const [addQuestionShow, setAddQuestionShow] = useState(false);
    const [editQuestionShow, setEditQuestionShow] = useState(false);
    const [deleteQuestionShow, setDeleteQuestionShow] = useState(false);
    const [questionError, setQuestionError] = useState("");

    const emptyQuestion: questionDataType = {
        description: "",
        question_tag: "",
        sub_questions: [],
        sub_question_number: -1,
        title: "",
        score: -1,
        id: "",
        hidden: false
    }

    const getQuestionsByTag = useCallback(async (tags: tagProps[]) => {
        const id = [...tags].filter((tag) => tag.name === params.tag)[0].id;
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions?tag_id=" + id + "&hidden=true");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result.data.data);
        setQuestionsByTag(result.data.data);
    }, [globalState.token, params.course_name, params.tag])

    useEffect(() => {
        getTags()
            .then(tags => getQuestionsByTag(tags))
            .catch();
    }, [getTags, getQuestionsByTag]);

    const [graderShow, setGraderShow] = useState(false);
    const [graderError, setGraderError] = useState("");

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <Tab.Container id="questionBank" defaultActiveKey={params.tag || tags[0].name}>
                <Row>
                    <Col xs={2} className="d-flex flex-column bg-light vh-100 sticky-top text-start p-3">
                        <Nav variant="pills" className="my-3">
                            {tags !== null &&
                                tags.map((tag) => (
                                    <Row className="d-flex flex-row align-items-center vw-100" key={tag.id}>
                                        <Col xs={8} className="text-start p-1">
                                            <Nav.Item onClick={() => setTag(tag)}>
                                                <Nav.Link eventKey={tag.name} href={tag.name}>{tag.name}</Nav.Link>
                                            </Nav.Item>
                                        </Col>
                                        <Col xs={4} className="text-end p-1">
                                            <i className="bi-pencil-square" style={{cursor: "pointer"}} onClick={() => {setTag(tag); setEditTagShow(true);}}/>
                                            <i className="bi-trash mx-2" style={{cursor: "pointer"}} onClick={() => {setTag(tag); setDeleteTagShow(true);}}/>
                                        </Col>
                                    </Row>
                                ))
                            }
                        </Nav>
                        <span style={{cursor: "pointer"}} onClick={() => setAddTagShow(true)}><i className="bi-plus-square mx-3"/>Add New Tag</span>
                    </Col>
                    {params.tag !== "null" &&
                        <Col xs={10}>
                            <Row className="text-end">
                                <Link to="#"><Button variant="primary" className='me-4 mt-4' onClick={() => setGraderShow(true)}>Grader</Button></Link>
                                <Link to="#"><Button variant="primary" className='me-4 my-4' onClick={() => setAddQuestionShow(true)}>Add Question</Button></Link>
                            </Row>
                            <Tab.Content className="text-start mx-4">
                                {tags.map((tag) => {
                                    if (tag.name === params.tag)
                                    return (
                                        <Tab.Pane eventKey={tag.name} key={tag.id}>
                                            <QuestionsByTag questions={questionsByTag} setQuestion={setQuestion}
                                                setDeleteShow={setDeleteQuestionShow} setEditShow={setEditQuestionShow}/>
                                        </Tab.Pane>);
                                })}
                            </Tab.Content>
                        </Col>
                    }
                </Row>
                
                <AddTagModal show={addTagShow}
                    errorMessage={tagError}
                    onClose={() => setAddTagShow(false)}
                    onSubmit={(tagName) => addNewTag(tagName)}
                    clearMessage={() => setTagError("")}
                />

                <EditTagModal show={editTagShow}
                    tag={(tag as tagProps)}
                    errorMessage={tagError}
                    onClose={() => setEditTagShow(false)}
                    onSubmit={(id, name) => editTag(id, name)}
                    clearMessage={() => setTagError("")}
                />

                <DeleteTagModal show={deleteTagShow}
                    tag={(tag as tagProps)}
                    errorMessage={tagError}
                    onClose={() => setDeleteTagShow(false)}
                    onSubmit={(id) => deleteTag(id)}
                    clearMessage={() => setTagError("")}
                />
                
                <AddQuestionModal
                    show={addQuestionShow}
                    onClose={() => setAddQuestionShow(false)}
                    tags={tags}
                    getTags={getTags}
                    getQuestionsByTag={getQuestionsByTag}
                    errorMsg={questionError}
                    setErrorMsg={setQuestionError}
                />

                <EditQuestionModal
                    show={editQuestionShow}
                    onClose={() => setEditQuestionShow(false)}
                    tags={tags}
                    getTags={getTags}
                    getQuestionsByTag={getQuestionsByTag}
                    question={question as questionDataType}
                    clearQuestion={() => setQuestion(emptyQuestion)}
                    errorMsg={questionError}
                    setErrorMsg={setQuestionError}
                />

                <DeleteQuestionModal
                    show={deleteQuestionShow}
                    onClose={() => setDeleteQuestionShow(false)}
                    tags={tags}
                    getTags={getTags}
                    getQuestionsByTag={getQuestionsByTag}
                    question={question as questionDataType}
                    clearQuestion={() => setQuestion(emptyQuestion)}
                    errorMsg={questionError}
                    setErrorMsg={setQuestionError}
                />

                <GraderModal
                    show={graderShow}
                    errorMessage={graderError}
                    onClose={() => setGraderShow(false)}
                    clearMessage={() => setGraderError("")}
                />
            </Tab.Container>
        </AppLayout>
    );
}

export default QuestionBank;
