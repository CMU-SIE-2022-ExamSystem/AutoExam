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
                <p>You are returning back to the assessment page.</p>
            </Modal.Body>

            <Modal.Footer>
                <Button variant="primary" onClick={onSubmit}>Confirm</Button>
                <Button variant="danger" onClick={onClose}>Close</Button>
            </Modal.Footer>
        </Modal>
    );
}

const ExamConfigExport = () => {
    let {examConfigState, setExamConfigState} = useConfigStates();
    const [publishModalShow, setPublishModalShow] = useState<boolean>(false);
    const publishedText = !!(examConfigState?.draft) ? "No" : "Yes";
    const publishOnClick = () => {

    }
    const publishOnSubmit = () => {

    }
    const removeOnClick = () => {

    }
    const publishButton = (<Button variant="primary" onClick={publishOnClick}>Publish</Button>);
    const removeButton = (<Button variant="danger" onClick={removeOnClick}>Remove Test</Button>)
    return (
        <div>
            <Row className="mt-3">
                <Col md={{span: '8', offset: '2'}}>
                    <div>
                        Published: {publishedText}
                        <span>{publishButton}</span>
                    </div>
                    <div>
                        Drop this test: <span>{removeButton}</span>
                    </div>
                </Col>
            </Row>
            <PublishModal show={publishModalShow} onSubmit={publishOnSubmit} onClose={() => setPublishModalShow(false)}/>
        </div>
    );
}

export default ExamConfigExport;