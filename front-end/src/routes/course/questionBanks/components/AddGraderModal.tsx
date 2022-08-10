import React, {useState} from 'react';
import {Button, Col, Form, InputGroup, Modal, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../../utils/url";
import axios from 'axios';

interface inputProps {
    id: number;
    type: string;
}

const UploadFileModal = ({show, onClose, name, getGraders, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, name: string, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [file, setFile] = useState<any>();
    const [fileName, setFileName] = useState("");

    const saveFile = (e: any) => {
        setFile(e.target.files[0]);
        setFileName(e.target.files[0].name)
    }

    const onSubmit = (e: any) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append("file", file);
        formData.append("fileName", fileName);
        uploadGraderFile(name, formData)
    }

    const uploadGraderFile = async (name: string, file: FormData) => {
        console.log("upload grader file: " + name)
        console.log(file)
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
    
    return (
        <Modal show={show} onHide={() => {onClose(); setErrorMsg("");}}>
            <Modal.Header>
                <Modal.Title>Add New Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Group className="mb-3">
                        <Form.Label>File</Form.Label><br/>
                        <Form.Text>upload the grader file with .py extension</Form.Text>
                        <Form.Control type="file" onChange={saveFile}/>
                    </Form.Group>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); setErrorMsg("");}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Upload</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    )
}

const AddGraderModal = ({show, onClose, getGraders, errorMsg, setErrorMsg}: {show: boolean, onClose: () => void, getGraders: () => void, errorMsg: string, setErrorMsg: any}) => {
    const params = useParams();
    const {globalState} = useGlobalState();
    
    const [name, setName] = useState("");

    const [type, setType] = useState("");
    const [inputIdx, setInputIdx] = useState(0);
    const [inputList, setInputList] = useState<inputProps[]>([]);

    const deleteInput = (idx: number) => {
        setInputList(inputList.filter((input) => input.id !== idx));
    }

    const inputs = inputList.map((input, index) => {
        let capitalizedType: string[] = []
        input.type.split("_").forEach((word) => {
            capitalizedType.push(word.charAt(0).toUpperCase() + word.slice(1))
        })
        const formattedType = capitalizedType.join(" ")
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={input.id}>
                <Col>
                    <Form.Label>{"(" + (index + 1) + ") " + formattedType}</Form.Label>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteInput(input.id)}/>
                </Col>
            </Row>
        );
    })

    const [moduleIdx, setModuleIdx] = useState(0);
    const [moduleList, setModuleList] = useState<number[]>([]);

    const deleteModule = (idx: number) => {
        setModuleList(moduleList.filter((module) => module !== idx));
    }

    const modules = moduleList.map((module) => {
        return (
            <Row className="d-flex flex-row align-items-center my-2" key={module}>
                <Col>
                    <Form.Control id={"module" + module}/>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteModule(module)}/>
                </Col>
            </Row>
        )
    })

    const [uploadFileShow, setUploadFileShow] = useState(false);

    const onSubmit = (e: any) => {
        e.preventDefault();

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
                const moduleName = (document.getElementById("module" + module) as HTMLInputElement).value;
                modulesData.push(moduleName);
            });
            return modulesData;
        }

        const graderData = {
            name: name,
            blanks: getBlanks(),
            modules: getModules()
        }
        addGrader(graderData);
    }

    const addGrader = async (graderData: object) => {
        console.log(graderData)
        const url = getBackendApiUrl("/courses/" + params.course_name + "/graders");
        const token = globalState.token;
        axios.post(url, graderData, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                onClose();
                setUploadFileShow(true);
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
        <>
        <Modal show={show} onHide={() => {onClose(); clearState(); setErrorMsg("");}} size="lg">
            <Modal.Header>
                <Modal.Title>Add New Grader</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <Form onSubmit={onSubmit}>
                    <Form.Group className="mb-3">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" placeholder="Name" onChange={(e) => setName(e.target.value)} required/>
                    </Form.Group>

                    <Form.Label>Input</Form.Label>
                    <div>{inputs}</div>

                    <InputGroup className="mb-3">
                        <Form.Select onChange={(e) => setType(e.target.value)}>
                            <option>Input Type</option>
                            <option value="blank">Blank</option>
                            <option value="code">Code</option>
                            <option value="single_choice">Single Choice</option>
                            <option value="multiple_choice">Multiple Choice</option>
                        </Form.Select>
                        <Button variant="primary" onClick={() => {if (type !== "") {setInputList([...inputList, {id: inputIdx, type: type}]); setInputIdx(inputIdx + 1);}}}>Add</Button>
                    </InputGroup>

                    <Form.Group className="mb-3">
                        <Form.Label>Modules</Form.Label>
                        {modules}
                        <div className='text-end'>
                            <Button variant="primary" onClick={() => {setModuleList([...moduleList, moduleIdx]); setModuleIdx(moduleIdx + 1)}}>Add Module</Button>
                        </div>
                    </Form.Group>

                    <div><small className="text-danger">{errorMsg}</small></div>

                    <div className="text-end">
                        <Button variant="secondary" onClick={() => {onClose(); clearState(); setErrorMsg("");}}>Close</Button>
                        <Button variant="primary" className="ms-2" type="submit">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>

        <UploadFileModal
            show={uploadFileShow}
            onClose={() => setUploadFileShow(false)}
            name={name}
            getGraders={getGraders}
            errorMsg={errorMsg}
            setErrorMsg={setErrorMsg}
        />
        </>
    )
}

export default AddGraderModal;
