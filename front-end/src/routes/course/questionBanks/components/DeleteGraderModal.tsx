import React from 'react';
import {Button, Modal} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';

const DeleteGraderModal = ({show, onClose, grader, getGraders, clearGrader, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, grader: graderDataType, getGraders: () => void, clearGrader: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const deleteGrader = (name: string) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name);
        const token = globalState.token;
        axios.delete(url, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                onClose();
                setErrorMsg("");
                getGraders();
            })
            .catch((error) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(typeof response.error.message === "string" ? response.error.message : response.error.message[0].message);
            });
    }
    
    return (
        <Modal show={show} onHide={() => {onClose(); clearGrader(); setErrorMsg("")}}>
            <Modal.Header closeButton>
                <Modal.Title>Delete Grader</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {grader !== undefined &&
                    "Do you want to delete grader \"" + grader.name + "\"?"
                }
                <div>
                    <small className="text-danger">{errorMsg}</small>
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => {onClose(); clearGrader(); setErrorMsg("")}}>Cancel</Button>
                <Button variant="primary" className="ms-2" onClick={() => deleteGrader(grader.name)}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    )
}

export default DeleteGraderModal;
