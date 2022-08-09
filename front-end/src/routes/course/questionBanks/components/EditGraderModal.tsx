import React, {useState} from 'react';
import {Button, Col, Form, InputGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';

const EditGraderModal = ({show, onClose, grader, getGraders, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, grader: graderDataType, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    return (
        <Modal show={show} onHide={() => {onClose(); setErrorMsg("");}} size="lg">
            <Modal.Header>
                <Modal.Title>Edit Grader</Modal.Title>
            </Modal.Header>
        </Modal>
    )
}

export default EditGraderModal;
