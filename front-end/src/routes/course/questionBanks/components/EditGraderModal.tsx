import React, {useEffect, useState} from 'react';
import {Button, Col, Form, InputGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';
import graderDataType from './graderDataType';

interface inputProps {
    input_idx: number;
    type: string;
}

interface moduleProps {
    module_idx: number;
    name: string;
}

const EditGraderModal = ({show, onClose, grader, getGraders, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, grader: graderDataType, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();

    const [name, setName] = useState("");

    const [type, setType] = useState("");
    const [inputIdx, setInputIdx] = useState(0);
    const [inputList, setInputList] = useState<inputProps[]>([]);

    const [moduleIdx, setModuleIdx] = useState(0);
    const [moduleList, setModuleList] = useState<moduleProps[]>([]);

    useEffect(() => {
        clearState();
        grader !== undefined &&
            setName(grader.name);
        grader !== undefined &&
            setInputIdx(grader.blanks.length);
        grader !== undefined &&
            grader.blanks.forEach((blank, index) => {
                setInputList((prevState) => ([
                    ...prevState,
                    {
                        input_idx: index,
                        type: blank.is_choice ? (blank.multiple? "multiple_choice" : "single_choice") : (blank.type === "string" ? "blank" : "code")
                    }
                ]))
            })
        grader !== undefined && grader.modules !== null &&
            setModuleIdx(grader.modules.length);
        grader !== undefined && grader.modules !== null &&
            grader.modules.forEach((module, index) => {
                setModuleList((prevState) => ([
                    ...prevState,
                    {
                        module_idx: index,
                        name: module
                    }
                ]))
            })
    }, [grader])

    const deleteInput = (idx: number) => {
        setInputList(inputList.filter((input) => input.input_idx !== idx));
    }

    const inputs = inputList.map((input, index) => {
        let capitalizedType: string[] = []
        input.type.split("_").forEach((word) => {
            capitalizedType.push(word.charAt(0).toUpperCase() + word.slice(1))
        })
        const formattedType = capitalizedType.join(" ")
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={input.input_idx}>
                <Col>
                    <Form.Label>{"(" + (index + 1) + ") " + formattedType}</Form.Label>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteInput(input.input_idx)}/>
                </Col>
            </Row>
        );
    })

    const deleteModule = (idx: number) => {
        setModuleList(moduleList.filter((module) => module.module_idx !== idx));
    }

    const modules = moduleList.map((module, index) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={module.module_idx}>
                <Col>
                    <Form.Control id={"module" + module.module_idx} defaultValue={module.name}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteModule(module.module_idx)}/>
                </Col>
            </Row>
        )
    })

    const [file, setFile] = useState<any>();
    const [fileName, setFileName] = useState("");

    const saveFile = (e: any) => {
        setFile(e.target.files[0]);
        setFileName(e.target.files[0].name)
    }

    const onSubmit = (e: any) => {
        e.preventDefault();

        if (fileName !== "") {
            const formData = new FormData();
            formData.append("file", file);
            formData.append("fileName", fileName);
            uploadGraderFile(name, formData);
        }

        const getBlanks = () => {
            const blanksData = inputList.map((input) => {
                let blank: object;
                if (input.type === "blank" || input.type === "code") {
                    blank = {
                        is_choice: false,
                        multiple: false,
                        type: input.type === "blank" ? "string" : "code"
                    }
                } else {
                    blank = {
                        is_choice: true,
                        multiple: input.type === "single_choice" ? false : true,
                        type: "string"
                    }
                }
                return blank;
            });
            return blanksData;
        }
    
        const getModules = () => {
            if (moduleList.length === 0) return null;
    
            let modulesData: string[] = []
            moduleList.forEach((module) => {
                const moduleName = (document.getElementById("module" + module.module_idx) as HTMLInputElement).value;
                modulesData.push(moduleName);
            });
            return modulesData;
        }

        const graderData = {
            name: name,
            blanks: getBlanks(),
            modules: getModules()
        }
        editGrader(name, graderData);
    }

    const uploadGraderFile = async (name: string, file: FormData) => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name + "/upload");
        const token = globalState.token;
        axios.put(url, file, {headers: {Authorization: "Bearer " + token}})
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

    const editGrader = async (name: string, graderData: object) => {
        console.log(graderData)
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders/" + name);
        const token = globalState.token;
        axios.put(url, graderData, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                onClose();
                clearState();
                setErrorMsg("");
                getGraders();
            })
            .catch((error) => {
                console.log(error);
                let response = error.response.data;
                setErrorMsg(typeof response.error.message === "string" ? response.error.message : response.error.message[0].message);
            });
    }

    const clearState = () => {
        setInputIdx(0);
        setInputList([]);
        setModuleIdx(0);
        setModuleList([]);
    }
    
    return (
        <Modal show={show} onHide={() => {onClose(); setErrorMsg("");}} backdrop="static" size="lg">
            <Modal.Header>
                <Modal.Title>Edit Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    {grader !== undefined &&
                    <>
                        <Form.Label>{"Name: " + grader.name}</Form.Label>

                        <Form.Group className="mb-3">
                            <Form.Label>Input</Form.Label>
                            {inputs}
                        </Form.Group>

                        <InputGroup className="mb-3">
                            <Form.Select onChange={(e) => setType(e.target.value)}>
                                <option>Input Type</option>
                                <option value="blank">Blank</option>
                                <option value="code">Code</option>
                                <option value="single_choice">Single Choice</option>
                                <option value="multiple_choice">Multiple Choice</option>
                            </Form.Select>
                            <Button variant="primary" onClick={() => {if (type !== "") {setInputList([...inputList, {input_idx: inputIdx, type: type}]); setInputIdx(inputIdx + 1);}}}>Add</Button>
                        </InputGroup>

                        <Form.Group className="mb-3">
                            <Form.Label>Module</Form.Label>
                            {modules}
                            <div className='text-end'>
                                <Button variant="primary" onClick={() => {setModuleList([...moduleList, {module_idx: moduleIdx, name: ""}]); setModuleIdx(moduleIdx + 1)}}>Add Module</Button>
                            </div>
                        </Form.Group>
                    </>
                    }

                    <Form.Group className="mb-3">
                        <Form.Label>File</Form.Label><br/>
                        <Form.Text>optional, update grader file with .py extension here</Form.Text>
                        <Form.Control type="file" onChange={saveFile}/>
                    </Form.Group>

                    <div><small className="text-danger">{errorMsg}</small></div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); setErrorMsg("");}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Confirm</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    )
}

export default EditGraderModal;
