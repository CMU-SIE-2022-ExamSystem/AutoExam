import React, {useCallback, useEffect, useState} from 'react';
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import AppLayout from "../../../components/AppLayout";
import TopNavbar from "../../../components/TopNavbar";
import {Alert, Button, Col, Container, Modal, Row, Table} from "react-bootstrap";

const AddModal = ({show, onSubmit, onCancel}: {show: boolean, onSubmit: (newBaseName: string) => void, onCancel: () => void}) => {
    const [newBaseName, setNewBaseName] = useState("");
    return (
        <Modal>
            <Modal.Header><Modal.Title>Add base course</Modal.Title></Modal.Header>
            <Modal.Body>

            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => onCancel()}>Back</Button>
                <Button variant="primary" onClick={() => onSubmit(newBaseName)}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const EditModal = ({show, oldBaseName, onSubmit, onCancel}: {show: boolean, oldBaseName: string, onSubmit: (oldName, newBaseName: string) => void, onCancel: () => void}) => {
    const [newBaseName, setNewBaseName] = useState(oldBaseName);
    return (
        <Modal>
            <Modal.Header><Modal.Title>Change base course name</Modal.Title></Modal.Header>
            <Modal.Body>

            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => onCancel()}>Back</Button>
                <Button variant="primary" onClick={() => onSubmit(oldBaseName, newBaseName)}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const DeleteModal = ({show, oldBaseName, onSubmit, onCancel}: {show: boolean, oldBaseName: string, onSubmit: (oldBaseName: string) => void, onCancel: () => void}) => {
    return (
        <Modal>
            <Modal.Header><Modal.Title>Delete base course</Modal.Title></Modal.Header>
            <Modal.Body>

            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => onCancel()}>Back</Button>
                <Button variant="primary" onClick={() => onSubmit(oldBaseName)}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const BaseCourseManage = () => {

    const params = useParams();
    const {globalState} = useGlobalState();

    const [listOfBaseCourses, setListOfBaseCourses] = useState<string[]>([]);
    const [baseCourseInInterest, setBaseCourseInInterest] = useState<string>("");

    const [addModalShow, setAddModalShow] = useState(false);
    const [editModalShow, setEditModalShow] = useState(false);
    const [deleteModalShow, setDeleteModalShow] = useState(false);

    const getBaseCourses = useCallback(() => {
        const baseCourseUrl = getBackendApiUrl("/basecourse/list");
        const token = globalState.token;
        return axios.get(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getBaseCourses()
            .then(response => {
                const listOfCourses : { id: number, name: string}[] = response.data.data;
                setListOfBaseCourses(listOfCourses.map((item) => item.name));
            })
    })

    const actionList = (baseCourse: string) => (
        <div>
            <Button variant="warning" onClick={() => {setBaseCourseInInterest(baseCourse); setEditModalShow(true);}}><i className="bi bi-pencil-square me-1"/>Edit</Button>
            <Button variant="danger" onClick={() => {setBaseCourseInInterest(baseCourse); setDeleteModalShow(true);}}><i className="bi bi-trash me-1"/>Delete</Button>
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
                <th scope="col">Course name</th>
                <th scope="col">Actions</th>
            </thead>
            {tbodyElement}
        </Table>
    )

    return (
        <AppLayout>
            <Row>
                <TopNavbar brandLink="/dashboard" />
            </Row>
            <Row>
                <Col md={{span: '8', offset: '2'}}>
                    <h1 className="mb-2">Base course management</h1>
                    <Container fluid className="text-end mb-3">
                        <Button variant="success" onClick={() => setAddModalShow(true)}>Add base course</Button>
                    </Container>
                    <Alert variant="info" className="mb-3">
                        Base course is a category that the course belongs to. Courses with same base course share question banks.
                        Usually we set the course number as base course.
                    </Alert>
                    {baseCourseTable}
                </Col>
            </Row>
        </AppLayout>
    )
}

export default BaseCourseManage;