import React, {useState} from 'react';
import {Button, Form, InputGroup, Modal} from 'react-bootstrap';
import HTMLEditor from "../../../../components/HTMLEditor";
import AddSingleBlank from './AddSingleBlank';
import AddChoice from './AddChoice';

const AddQuestionModal = ({tag, show, onClose} : {tag: string, show: boolean, onClose: () => void}) => {
    const [title, setTitle] = useState("");
    const [description, setDescription]= useState<string>("");

    const updateDescription = (newDescription: string) => {
        setDescription(newDescription);
    }

    const [type, setType] = useState("");
    const [id, setId] = useState(0);
    const [subqList, setSubqList] = useState<any[]>([]);

    const deleteSubq = (id: number) => {
        setSubqList(subqList.filter((subq) => subq.id !== id));
    }
    
    const subquestions = (subqList).map(({type, id}) => {
        if (type === "single-blank") return (<AddSingleBlank key={id} id={id} onDelete={deleteSubq}/>);
        if (type === "single-choice") return (<AddChoice key={id} id={id} onDelete={deleteSubq}/>);
        if (type === "multiple-choice") return (<AddChoice key={id} id={id} onDelete={deleteSubq}/>);
        // if (subqType === "customized") return (<AddCustomizedQuestion/>);
        return(<></>);
    });

    return (
        <Modal show={show} onHide={onClose} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Add new Question</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Label>Tag: {tag}</Form.Label>

                    <Form.Group className="mb-3">
                        <Form.Label>Title </Form.Label>
                        <Form.Control type="text" placeholder="Title" onChange={(e) => setTitle(e.target.value)} required/>
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label>Description</Form.Label>
                            <div>
                                <HTMLEditor init={description} update={updateDescription}/>
                            </div>
                    </Form.Group>

                    <div>{subquestions}</div>

                    <InputGroup className="mb-3">
                        <Form.Select onChange={(e) => setType(e.target.value)}>
                            <option>Subquestion Type</option>
                            <option value="single-blank">Single Blank</option>
                            <option value="single-choice">Single Choice</option>
                            <option value="multiple-choice">Multiple Choice</option>
                            <option value="customized">Customized</option>
                        </Form.Select>
                        <Button variant="primary"
                            onClick={() => {if (type !== "") setSubqList([...subqList, {type: type, id: id}]); setId(id + 1);}}>
                            Add Subquestion
                        </Button>
                    </InputGroup>

                    <div className="text-end">
                        <Button variant="secondary" onClick={onClose}>Close</Button>
                        <Button variant="primary" type="submit" className="ms-2">Add</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}

export default AddQuestionModal;
