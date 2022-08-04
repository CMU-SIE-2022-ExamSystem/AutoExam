import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, Form, ListGroup, Modal, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from 'axios';
import CollapseQuestion from './components/CollapseQuestion';
import AddQuestionModal from './components/AddQuestionModal';

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
                            // defaultValue={tag.name}
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
                Do you want to delete this tag?
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

const GraderModal = ({show, errorMessage, onClose, clearMessage}: {show: boolean, errorMessage: string, onClose: () => void, clearMessage: () => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [graders, setGraders] = useState<string[]>([]);

    const getGraders = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/list");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <ListGroup variant="flush">
                    {graders.map((grader, index) => (
                        <ListGroup.Item key={index}>
                            <Row>
                                <Col xs={10}>{grader}</Col>
                                <Col xs={2}>
                                {grader === "single_blank" || grader === "single_choice" || grader === "multiple_choice" ?
                                    <div className="text-secondary">Basic</div> :
                                    <>
                                    <i className="bi-pencil-square" style={{cursor: "pointer"}}/>
                                    <i className="bi-trash ms-2" style={{cursor: "pointer"}}/>
                                    </>
                                }
                                </Col>
                            </Row>
                        </ListGroup.Item>
                    ))}
                </ListGroup>
            </Modal.Body>
        </Modal>
    )
}

const QuestionsByTag = ({questions, deleteShow, setDeleteShow, onDelete, editShow, setEditShow, onEdit, error, setError}:
        {questions: questionDataType[] | undefined, deleteShow: boolean, setDeleteShow: any, onDelete: (id: string) => void,
        editShow: boolean, setEditShow: any, onEdit: (id: string, data: object) => void, error: string, setError: any}) => {
    return (
        <Row>
            <Col sm={10}>
                {!!questions &&
                    questions.map((question, index) => {
                        return <CollapseQuestion question={question} key={index}
                            deleteShow={deleteShow} setDeleteShow ={setDeleteShow} onDelete={onDelete}
                            editShow={editShow} setEditShow={setEditShow} onEdit={onEdit}
                            error={error} setError={setError}/>
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
        setTags(result.data.data);
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
                let response = error.response.data;
                setTagError(response.error.message[0].message);
            });
    };

    const editTag = async (id: string, name: string) => {
        console.log("edit tag " + id);
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
                setTagError(response.error.message);
            })
    }

    const deleteTag = async (id: string) => {
        console.log("delete tag " + id);
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
    const [addQuestionShow, setAddQuestionShow] = useState(false);
    const [editQuestionShow, setEditQuestionShow] = useState(false);
    const [deleteQuestionShow, setDeleteQuestionShow] = useState(false);
    const [questionError, setQuestionError] = useState("");

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

    const addNewQuestion = async (questionData: object) => {
        console.log(questionData);
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions");
        const token = globalState.token;
        // axios.post(url, questionData, {headers: {Authorization: "Bearer " + token}})
        //     .then(_ => {
        //         setAddQuestionShow(false);
        //         getTags()
        //             .then(tags => getQuestionsByTag(tags));
        //     })
        //     .catch((error: any) => {
        //         console.log(error);
        //         let response = error.response.data;
        //         setQuestionError(response.error.message);
        //     });
    }

    const editQuestion = async (id: string, questionData: object) => {
        console.log("edit question: " + id);
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions/" + id);
        const token = globalState.token;
        axios.put(url, questionData, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setEditQuestionShow(false);
                getTags()
                    .then(tags => getQuestionsByTag(tags));
            })
            .catch((error: any) => {
                console.log(error);
                let response = error.response.data;
                setQuestionError(response.error.message);
            });
    }

    const deleteQuestion = async (id: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/questions/" + id);
        const token = globalState.token;
        axios.delete(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setDeleteQuestionShow(false);
                getTags()
                    .then(tags => getQuestionsByTag(tags));
            })
            .catch((error: any) => {
                console.log(error);
                let response = error.response.data;
                setQuestionError(response.error.message);
            });
    }

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
                            {
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
                        <span style={{cursor: "pointer"}} onClick={() => setAddTagShow(true)}><i className="bi-plus-square mx-3"/> Add new tags</span>
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
                                            <QuestionsByTag questions={questionsByTag}
                                                deleteShow={deleteQuestionShow} setDeleteShow={setDeleteQuestionShow} onDelete={deleteQuestion}
                                                editShow={editQuestionShow} setEditShow={setEditQuestionShow} onEdit={editQuestion}
                                                error={questionError} setError={setQuestionError}/>
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
                    tag={([...tags].filter((tag) => tag.name === params.tag)[0] as tagProps)}
                    show={addQuestionShow}
                    errorMessage={questionError}
                    onAdd={addNewQuestion}
                    onClose={() => setAddQuestionShow(false)}
                    clearMessage={() => setQuestionError("")}
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
