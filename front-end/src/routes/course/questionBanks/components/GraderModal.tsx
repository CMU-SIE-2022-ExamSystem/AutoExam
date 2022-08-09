import React, {useCallback, useEffect, useState} from 'react';
import {Button, Col, Form, InputGroup, ListGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';
import AddGraderModal from './AddGraderModal';
import EditGraderModal from './EditGraderModal';
import DeleteGraderModal from './DeleteGraderModal';

const GraderModal = ({show, errorMessage, onClose, clearMessage}: {show: boolean, errorMessage: string, onClose: () => void, clearMessage: () => void}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [graders, setGraders] = useState<graderDataType[]>([]);
    const [grader, setGrader] = useState<graderDataType>();
    const [addGraderShow, setAddGraderShow] = useState(false);
    const [editGraderShow, setEditGraderShow] = useState(false);
    const [deleteGraderShow, setDeleteGraderShow] = useState(false);
    const [graderError, setGraderError] = useState("");

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
                                    <i className="bi-pencil-square" style={{cursor: "pointer" }} onClick={() => {setGrader(grader); onClose(); setEditGraderShow(true)}}/>
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
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />

        <DeleteGraderModal
            show={deleteGraderShow}
            onClose={() => setDeleteGraderShow(false)}
            grader={grader as graderDataType}
            getGraders={getGraders}
            errorMsg={graderError}
            setErrorMsg={setGraderError}
        />
        </>
    )
}

export default GraderModal;
