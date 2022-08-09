import React, {useCallback, useEffect, useState} from 'react';
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import AppLayout from "../../../components/AppLayout";
import TopNavbar from "../../../components/TopNavbar";
import {Alert, Button, Col, Container, Form, Modal, Row, Table} from "react-bootstrap";

const AddModal = ({show, onSubmit, onCancel}: {show: boolean, onSubmit: (newBaseName: string) => void, onCancel: () => void}) => {
    const validate = () => {
        const baseCourseName = (document.getElementById('new-base-course') as HTMLInputElement).value;
        (document.getElementById('new-base-course') as HTMLInputElement).value = "";
        onSubmit(baseCourseName);
    }
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Add base course</Modal.Title></Modal.Header>
            <Modal.Body>
                <p>Please type the new base course name:</p>
                <Form>
                    <Form.Group className="pb-4">
                        <Form.Control type="text"
                                      className="mb-2"
                                      required
                                      id="new-base-course"
                        />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={validate}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const EditModal = ({show, oldBaseName, onSubmit, onCancel}: {show: boolean, oldBaseName: string, onSubmit: (oldName: string, newBaseName: string) => void, onCancel: () => void}) => {
    const validate = () => {
        const baseCourseName = (document.getElementById('new-base-course-edit') as HTMLInputElement).value;
        (document.getElementById('new-base-course-edit') as HTMLInputElement).value = "";
        onSubmit(oldBaseName, baseCourseName);
    }
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Change base course name</Modal.Title></Modal.Header>
            <Modal.Body>
                <p>Changing '{oldBaseName}' to the following:</p>
                <Form>
                    <Form.Group className="pb-4">
                        <Form.Control placeholder={oldBaseName}
                                      type="text"
                                      className="mb-2"
                                      required
                                      id="new-base-course-edit"
                        />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={validate}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const DeleteModal = ({show, oldBaseName, onSubmit, onCancel}: {show: boolean, oldBaseName: string, onSubmit: (oldBaseName: string) => void, onCancel: () => void}) => {
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Delete base course</Modal.Title></Modal.Header>
            <Modal.Body>
                <p>Removing '{oldBaseName}' base course.</p>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={() => onSubmit(oldBaseName)}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const BaseCourseManage = () => {

    const {globalState, updateGlobalState} = useGlobalState();

    const [listOfBaseCourses, setListOfBaseCourses] = useState<string[]>([]);
    const [baseCourseInInterest, setBaseCourseInInterest] = useState<string>("");

    const [addModalShow, setAddModalShow] = useState(false);
    const [editModalShow, setEditModalShow] = useState(false);
    const [deleteModalShow, setDeleteModalShow] = useState(false);

    const getBaseCourses = useCallback(async () => {
        const baseCourseUrl = getBackendApiUrl("/basecourses/list");
        const token = globalState.token;
        const response = await axios.get(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
        const listOfCourses : { id: number, name: string}[] = response.data.data;
        setListOfBaseCourses(listOfCourses.map((item) => item.name));
    }, []);

    useEffect(() => {
        getBaseCourses()
            .catch();
    }, [])

    const addBaseCourseApi = useCallback((newBaseCourse: string) => {
        const baseCourseUrl = getBackendApiUrl("/basecourses/create");
        const token = globalState.token;
        const data = {name: newBaseCourse};
        return axios.post(baseCourseUrl, data,{headers: {Authorization: "Bearer " + token}});
    }, []);

    const addHandler = (newBaseCourse: string) => {
        setAddModalShow(false);
        addBaseCourseApi(newBaseCourse)
            .then(async () => await getBaseCourses())
            .then(() => updateGlobalState({alert: {show: true, content: "New base course " + newBaseCourse + " created", variant: "success"}}))
            .catch(error => {
                updateGlobalState({alert: {show: true, content: "Error: " + error.toString(), variant: "danger"}})
            })
    }

    const editBaseCourseApi = useCallback((oldBaseCourse: string, newBaseCourse: string) => {
        const baseCourseUrl = getBackendApiUrl("/basecourses/" + oldBaseCourse);
        const token = globalState.token;
        const data = {name: newBaseCourse};
        return axios.put(baseCourseUrl, data,{headers: {Authorization: "Bearer " + token}});
    }, []);

    const editHandler = (oldBaseCourse: string, newBaseCourse: string) => {
        setEditModalShow(false);
        editBaseCourseApi(oldBaseCourse, newBaseCourse)
            .then(async () => await getBaseCourses())
            .then(() => updateGlobalState({alert: {show: true, content: "Successfully changed from " + oldBaseCourse + " to " + newBaseCourse, variant: "success"}}))
            .catch(error => {updateGlobalState({alert: {show: true, content: "Error: " + error.toString(), variant: "danger"}})})
    }

    const deleteBaseCourseApi = useCallback((oldBaseCourse: string) => {
        const baseCourseUrl = getBackendApiUrl("/basecourses/" + oldBaseCourse);
        const token = globalState.token;
        return axios.delete(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    const deleteHandler = (oldBaseCourse: string) => {
        setDeleteModalShow(false);
        deleteBaseCourseApi(oldBaseCourse)
            .then(async () => await getBaseCourses())
            .then(() => updateGlobalState({alert: {show: true, content: "Base course " + oldBaseCourse + " deleted", variant: "success"}}))
            .catch(error => {updateGlobalState({alert: {show: true, content: "Error: " + error.toString(), variant: "danger"}})})
    }

    const actionList = (baseCourse: string) => (
        <div>
            <Button className="me-1" variant="warning" onClick={() => {setBaseCourseInInterest(baseCourse); setEditModalShow(true);}}><i className="bi bi-pencil-square me-1"/>Edit</Button>
            <Button className="me-1" variant="danger" onClick={() => {setBaseCourseInInterest(baseCourse); setDeleteModalShow(true);}}><i className="bi bi-trash me-1"/>Delete</Button>
        </div>
    )

    const tbodyElement = (
        <tbody>
            {
                listOfBaseCourses.map(baseCourse => {
                    return (<tr key={"baseCourse_" + baseCourse}><td>{baseCourse}</td><td>{actionList(baseCourse)}</td></tr>)
                })
            }
        </tbody>
    )

    const baseCourseTable = (
        <Table striped>
            <thead>
                <tr>
                    <th scope="col">Course name</th>
                    <th scope="col">Actions</th>
                </tr>
            </thead>
            {tbodyElement}
        </Table>
    )

    return (
        <AppLayout>
            <Row>
                <TopNavbar brandLink="/dashboard" />
            </Row>
            <Row className="mt-2">
                <Col md={{span: '8', offset: '2'}}>
                    <h1 className="mb-2">Base course management</h1>
                    <Container fluid className="text-end mb-3">
                        <Button variant="success" onClick={() => setAddModalShow(true)}>Add base course</Button>
                    </Container>
                    {baseCourseTable}
                    <Alert variant="info" className="mb-3 text-start" dismissible>
                        Base course is a category that the course belongs to.<br />
                        Courses with same base course share question banks with each other.<br />
                        By convention, we set course numbers as base course names.
                    </Alert>
                </Col>
            </Row>
            <AddModal show={addModalShow} onSubmit={addHandler} onCancel={() => {setAddModalShow(false);}} />
            <EditModal show={editModalShow} oldBaseName={baseCourseInInterest} onSubmit={editHandler} onCancel={() => {setEditModalShow(false);}} />
            <DeleteModal show={deleteModalShow} oldBaseName={baseCourseInInterest} onSubmit={deleteHandler} onCancel={() => {setDeleteModalShow(false);}} />
        </AppLayout>
    )
}

export default BaseCourseManage;