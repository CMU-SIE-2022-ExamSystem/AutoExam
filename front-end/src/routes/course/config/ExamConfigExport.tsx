import React, {useState} from 'react';
import {Button, Col, Modal, Row} from "react-bootstrap";
import {useConfigStates} from "./ExamConfigStates";

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
                <Button variant="secondary" onClick={onClose}>Close</Button>
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
                <p>You are deleting this exam.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={onClose}>Close</Button>
                <Button variant="danger" onClick={onSubmit}>Delete</Button>
            </Modal.Footer>
        </Modal>
    );
}

const ExamConfigExport = () => {
    let {examConfigState} = useConfigStates();
    const [publishModalShow, setPublishModalShow] = useState<boolean>(false);
    const publishedText = !!(examConfigState?.draft) ? "No" : "Yes";
    const publishOnClick = () => {
        setPublishModalShow(true);
    }
    const publishOnSubmit = () => {
        setPublishModalShow(false);
    }
    const [removeModalShow, setRemoveModalShow] = useState<boolean>(false);
    const removeOnClick = () => {
        setRemoveModalShow(true);
    }
    const removeOnSubmit = () => {
        setRemoveModalShow(false);
    }
    const downloadExamPackage = () => {
        console.log("downloadExam");
    }
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
                        Published: {publishedText}
                        <span className="ms-3">{publishButton}</span>
                    </div>
                    <div>
                        Drop this test: <span className="ms-3">{removeButton}</span>
                    </div>
                </Col>
            </Row>
            <PublishModal show={publishModalShow} onSubmit={publishOnSubmit} onClose={() => setPublishModalShow(false)}/>
            <RemoveModal show={removeModalShow} onSubmit={removeOnSubmit} onClose={() => setRemoveModalShow(false)} />
        </div>
    );
}

export default ExamConfigExport;