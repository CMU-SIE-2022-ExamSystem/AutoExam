import React, {useState} from 'react';
import {Alert, Button, Col, Container, Modal, Row} from "react-bootstrap";
import {useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {RightBottomAlert} from "../../../components/RightBottomAlert";
import {downloadBlob} from "../../../utils/downloadFile";

const PublishModal = ({show, onSubmit, onClose} :{ show: boolean, onSubmit: () => void, onClose: () => void }) => {
    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Warning</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>By pressing down "Confirm", You are sure that you have uploaded the exam zip to AutoLab, and you are ready for publishing this exam.</p>
                <p>Your Students will be able to see this exam after the exam is published.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>Back</Button>
                <Button variant="primary" onClick={onSubmit}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const RemoveModal = ({show, onSubmit, onClose} :{ show: boolean, onSubmit: () => void, onClose: () => void }) => {
    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Warning</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>You are permanently removing this exam. This operation cannot be undone.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>Back</Button>
                <Button variant="danger" onClick={onSubmit}>Delete</Button>
            </Modal.Footer>
        </Modal>
    );
}

const ExamConfigExport = () => {
    let {examConfigState, setExamConfigState} = useConfigStates();
    const [publishModalShow, setPublishModalShow] = useState<boolean>(false);
    const publishedText = !!(examConfigState?.draft) ? "No" : "Yes";

    const [alertVariant, setAlertVariant] = useState<string>("success");
    const [alertShow, setAlertShow] = useState<boolean>(false);
    const [alertContent, setAlertContent] = useState<string>("");


    let params = useParams();
    const courseName = params.course_name;
    const examId = params.exam_id;
    const {globalState} = useGlobalState();
    const navigate = useNavigate();

    const publishOnClick = () => {
        setPublishModalShow(true);
    }
    const publishOnSubmit = () => {
        setPublishModalShow(false);
        const url = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId + "/draft");
        const token = globalState.token;
        axios.put(url, {draft: false}, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                const newState = Object.assign({}, examConfigState, {draft: false});
                setExamConfigState(newState);
            })
            .catch(onFailure => {
                setAlertVariant("danger");
                setAlertContent(onFailure.toString().substring(0, 50))
                console.error(onFailure);
                setAlertShow(true);
            })
    }

    const generateExams = () => {
        const url = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId + "/generate");
        const token = globalState.token;
        axios.get(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                setAlertVariant("success");
                setAlertContent("Successfully generated.")
                setAlertShow(true);
            })
            .catch(onFailure => {
                setAlertVariant("danger");
                setAlertContent(onFailure.toString().substring(0, 50))
                console.error(onFailure);
                setAlertShow(true);
            })
    }
    const [removeModalShow, setRemoveModalShow] = useState<boolean>(false);
    const removeOnClick = () => {
        setRemoveModalShow(true);
    }
    const removeOnSubmit = async () => {
        setRemoveModalShow(false);
        const url = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId);
        const token = globalState.token;
        axios.delete(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                navigate("/courses/" + courseName);
            })
            .catch(onFailure => {
                setAlertVariant("danger");
                setAlertContent(onFailure.toString().substring(0, 50))
                console.error(onFailure);
                setAlertShow(true);
            });

    }
    const downloadExamPackage = () => {
        const downloadUrl = getBackendApiUrl("/courses/" + courseName + "/assessments/" + examId + "/download");
        const token = globalState.token;
        axios.get(downloadUrl, {headers: {Authorization: "Bearer " + token}})
            .then(response => {
                const url = window.URL.createObjectURL(new Blob([response.data]));
                downloadBlob(examId+".tar", url);
            })
    }
    const generateButton = (<Button variant="info" onClick={generateExams}>Generate</Button>);
    const publishButton = (<Button variant="success" onClick={publishOnClick}>Publish</Button>);
    const removeButton = (<Button variant="danger" onClick={removeOnClick}>Remove Test</Button>)

    return (
        <div>
            <Row className="mt-3 text-start">
                <Col md={{span: '8', offset: '2'}}>
                    <div className="mb-4 text-center">
                        <Button variant="primary" onClick={downloadExamPackage}>Download Package for AutoLab</Button>
                    </div>
                    <hr></hr>
                    <div className="mb-3">
                        Generate exams: <span className="ms-3">{generateButton}</span>
                    </div>
                    <div className="mb-3">
                        Published: {publishedText}
                        <span className="ms-3">{publishButton}</span>
                    </div>
                    <div>
                        Drop this test: <span className="ms-3">{removeButton}</span>
                    </div>
                </Col>
            </Row>
            <Container className="text-start mt-3">
                <Alert variant="secondary">
                    <i className="bi bi-info-circle"/> First Time Reminder: Steps to publish an exam
                    <ol>
                        <li>Download the package for Autolab using the blue button.</li>
                        <li>Upload the package to Autolab.</li>
                        <li>Generate the exams using the cyan button.</li>
                        <li>Publish the exam to students using the green button.</li>
                    </ol>
                </Alert>
            </Container>
            <PublishModal show={publishModalShow} onSubmit={publishOnSubmit} onClose={() => setPublishModalShow(false)}/>
            <RemoveModal show={removeModalShow} onSubmit={removeOnSubmit} onClose={() => setRemoveModalShow(false)} />
            <RightBottomAlert variant={alertVariant} content={alertContent} show={alertShow} onClose={() => {setAlertShow(false)}} />
        </div>
    );
}

export default ExamConfigExport;