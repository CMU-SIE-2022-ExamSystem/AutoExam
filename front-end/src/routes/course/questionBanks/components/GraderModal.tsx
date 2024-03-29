import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, ListGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';
import AddGraderModal from './AddGraderModal';
import EditGraderModal from './EditGraderModal';
import DeleteGraderModal from './DeleteGraderModal';
import TestGraderModal from './TestGraderModal';

const GraderModal = ({show, errorMessage, onClose, clearMessage}: {show: boolean, errorMessage: string, onClose: () => void, clearMessage: () => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [graders, setGraders] = useState<graderDataType[]>([]);
    const [grader, setGrader] = useState<graderDataType>();
    const [addGraderShow, setAddGraderShow] = useState(false);
    const [editGraderShow, setEditGraderShow] = useState(false);
    const [deleteGraderShow, setDeleteGraderShow] = useState(false);
    const [testGraderShow, setTestGraderShow] = useState(false);
    const [graderError, setGraderError] = useState("");

    const emptyGrader: graderDataType = {
        name: "",
        blanks: [],
        modules: [],
        valid: false,
        uploaded: false
    }

    const getGraders = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        setGraders(result.data.data);
    }, [globalState.token, params.course_name])

    useEffect(() => {
        getGraders().catch();
    }, [getGraders])

    return (
        <>
        <Modal show={show} onHide={() => {onClose(); clearMessage()}} size="lg">
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
                                <Col xs={7}>{grader.name}</Col>
                                <Col xs={5} className="text-end">
                                    {
                                        grader.valid && <Button size="sm" variant="outline-success" disabled>Valid</Button>
                                    }
                                    <Button size="sm" variant="outline-primary" className="ms-2" onClick={() => {setGrader(grader); onClose(); setTestGraderShow(true)}}>Test</Button>
                                    <i className="bi-pencil-square ms-3" style={{cursor: "pointer" }} onClick={() => {setGrader(grader); onClose(); setEditGraderShow(true)}}/>
                                    <i className="bi-trash ms-2" style={{cursor: "pointer"}} onClick={() => {setGrader(grader); onClose(); setDeleteGraderShow(true)}}/>
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
            getGraders={getGraders}
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />

        <EditGraderModal
            show={editGraderShow}
            onClose={() => setEditGraderShow(false)}
            grader={grader as graderDataType}
            getGraders={getGraders}
            clearGrader={() => setGrader(emptyGrader)}
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />

        <DeleteGraderModal
            show={deleteGraderShow}
            onClose={() => setDeleteGraderShow(false)}
            grader={grader as graderDataType}
            getGraders={getGraders}
            clearGrader={() => setGrader(emptyGrader)}
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />

        <TestGraderModal
            show={testGraderShow}
            setTestGraderShow={setTestGraderShow}
            onClose={() => setTestGraderShow(false)}
            grader={grader as graderDataType}
            getGraders={getGraders}
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />
        </>
    )
}

export default GraderModal;
