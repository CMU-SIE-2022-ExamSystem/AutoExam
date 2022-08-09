import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, Form, ListGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';

const AddGraderModal = ({show, onClose}: {show: boolean, onClose: () => void}) => {
    const [name, setName] = useState("");
    const [file, setFile] = useState();
    const [fileName, setFileName] = useState("");

    const saveFile = (e: any) => {
        setFile(e.target.files[0]);
        setFileName(e.target.files[0].name)
    }
    return (
        <Modal show={show} onHide={() => {onClose();}} size="lg">
            <Modal.Header>
                <Modal.Title>Add New Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form>
                    <Form.Group className="mb-3">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" placeholder="Name" onChange={(e) => setName(e.target.value)} required/>
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>File</Form.Label><br/>
                        <Form.Text>optional, upload a file with .py extension</Form.Text>
                        <Form.Control type="file" onChange={saveFile}/>
                    </Form.Group>

                </Form>
            </Modal.Body>
        </Modal>
    )
}

const GraderModal = ({show, errorMessage, onClose, clearMessage}: {show: boolean, errorMessage: string, onClose: () => void, clearMessage: () => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [graders, setGraders] = useState<graderDataType[]>([]);

    const getGraders = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    const [addGraderShow, setAddGraderShow] = useState(false);
    const [editGraderShow, setEditGraderShow] = useState(false);
    const [deleteGraderShow, setDeleteGraderShow] = useState(false);
    const [graderError, setGeaderError] = useState("");

    return (
        <>
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <div className="text-end">
                    <Button variant="primary" className='me-2 my-2' onClick={() => {onClose(); setAddGraderShow(true)}}>Add Grader</Button>
                </div>
                <ListGroup variant="flush">
                    {graders.map((grader, index) => (
                        <ListGroup.Item key={index}>
                            <Row>
                                <Col xs={10}>{grader.name}</Col>
                                <Col xs={2}>
                                {grader.name === "single_blank" || grader.name === "single_choice" || grader.name === "multiple_choice" ?
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
        
        <AddGraderModal
            show={addGraderShow}
            onClose={() => setAddGraderShow(false)}
        
        />
        </>
    )
}

export default GraderModal;
