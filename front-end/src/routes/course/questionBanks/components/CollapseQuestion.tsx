import React, {useState} from 'react';
import {Button, Card, Collapse, Modal} from 'react-bootstrap';
import {subQuestionDataType} from "../../../../components/questionTemplate/subQuestionDataType";
import questionDataType from "../../../../components/questionTemplate/questionDataType";
import BlankWithSolution from './BlankWithSolution';
import ChoiceWithSolution from './ChoiceWithSolution';
import CustomizedWithSolution from './CustomizedWithSolution';

const EditQuestionModal = ({show, question, errorMessage, onEdit, onClose, clearMessage}: {show: boolean, question: questionDataType, errorMessage: string, onEdit: (id: string, data: object) => void, onClose: () => void, clearMessage: () => void}) => {
    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}} size="lg">
            <Modal.Header closeButton>
                <Modal.Title>Edit Queston</Modal.Title>
            </Modal.Header>
        </Modal>
    );
}

const DeleteQuestionModal = ({show, questionId, errorMessage, onDelete, onClose, clearMessage}: {show: boolean, questionId: string, errorMessage: string, onDelete: (id: string) => void, onClose: () => void, clearMessage: () => void}) => {
    return (
        <Modal show={show} onHide={() => {onClose(); clearMessage()}}>
            <Modal.Header closeButton>
                <Modal.Title>Delete Queston</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                Do you want to delete this question?
                <div>
                    <small className="text-danger">{errorMessage}</small>
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={() => {onClose(); clearMessage()}}>Cancel</Button>
                <Button variant="primary" type="submit" className="ms-2" onClick={() => onDelete(questionId)}>Confirm</Button>
            </Modal.Footer>
        </Modal>
    );
}

const CollapseQuestion = ({question, deleteShow, setDeleteShow, onDelete, editShow, setEditShow, onEdit, error, setError} :
        {question: questionDataType, deleteShow: boolean, setDeleteShow: any, onDelete: (id: string) => void,
        editShow: boolean, setEditShow: any, onEdit: (id: string, data: object) => void, error: string, setError: any}) => {
    const [open, setOpen] = useState(false);

    const subQuestions = question.sub_questions.map((subQuestion: subQuestionDataType, index) => {
        if (subQuestion.grader === "single_blank") return (<BlankWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        if (subQuestion.grader === "single_choice" || subQuestion.grader === "multiple_choice") return (<ChoiceWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        return (<CustomizedWithSolution key={index + 1} index={index + 1} subQuestion={subQuestion}/>);
    });

    return (
        <>
            <Card>
                <Card.Header
                    style={{cursor: "pointer"}}
                    onClick={() => setOpen(!open)}
                    aria-expanded={open}>
                    {question.title}
                </Card.Header>
                <Collapse in={open}>
                    <div>
                    <div className="text-end my-3 me-3">
                        <Button variant="success" onClick={() => setEditShow(true)}>Edit</Button>
                        <Button variant="secondary" className="ms-2" onClick={() => setDeleteShow(true)}>Delete</Button>
                    </div>
                    <Card.Body>
                        <div dangerouslySetInnerHTML={{__html: question.description}}/>
                        {subQuestions}
                    </Card.Body>
                    </div>
                </Collapse>
            </Card>
            <br/>

            <EditQuestionModal
                show={editShow}
                question={question}
                errorMessage={error}
                onEdit={onEdit}
                onClose={() => setEditShow(false)}
                clearMessage={() => setError("")}
            />

            <DeleteQuestionModal
                show={deleteShow}
                questionId={question.id}
                errorMessage={error}
                onDelete={onDelete}
                onClose={() => setDeleteShow(false)}
                clearMessage={() => setError("")}
            />
        </>
    );
}

export default CollapseQuestion;
