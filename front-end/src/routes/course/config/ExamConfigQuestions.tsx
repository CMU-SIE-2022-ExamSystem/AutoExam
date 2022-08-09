import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {ExamConfigSettingsType, useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {Button, Card, Form, Modal, Table} from "react-bootstrap";

interface tagProps {
    id: string;
    name: string;
    course: string;
}

const settingToQuestion = (setting: ExamConfigSettingsType, qIndex: number, tags: tagProps[], editWrapper: (arg0: number) => void, deleteWrapper: (arg0: number) => void) => {
    const myTag = tags.find(tag => tag.id === setting.tag);
    const idLength = setting.id.length;
    const appointed = idLength > 0 ? ("Chosen from " + idLength + "questions") : "Random pick according to tag";
    const subquestionTitles = Array.from({length: setting.sub_question_number}, (value, index) => "sub" + (index + 1).toString());
    const subquestionScores = (
        <Table bordered>
            <thead>
                <tr><th>Questions</th>{subquestionTitles.map(title => <th scope="col" key={title}>{title}</th>)}</tr>
            </thead>
            <tbody>
                <tr><td>Scores</td>{setting.scores.map(score => <td>{score}</td>)}</tr>
            </tbody>
        </Table>
    )
    return (
        <Card className="my-3 text-start" key={"question_" + qIndex}>
            <Card.Header>
                {qIndex + 1}. {myTag ? myTag.name : setting.title}
            </Card.Header>
            <Card.Body>
                <Card.Text>Question title: {setting.title}</Card.Text>
                <Card.Text>Score: {setting.max_score}</Card.Text>
                <Card.Text># of sub questions: {setting.sub_question_number}</Card.Text>
                {subquestionScores}
                <Card.Text>Generation source: {appointed}</Card.Text>
            </Card.Body>
            <Card.Footer>
                <Button variant="warning" size="sm" className="me-1" onClick={() => editWrapper(qIndex)}><i className="bi bi-pencil-square me-1"/>Edit</Button>
                <Button variant="danger" size="sm" onClick={() => deleteWrapper(qIndex)}><i className="bi bi-trash me-1"/>Delete</Button>
            </Card.Footer>
        </Card>
    );
};

const AddModal = ({show, onSubmit, onCancel}: {show: boolean, onSubmit: () => void, onCancel: () => void}) => {
    const validate = () => {
        const baseCourseName = (document.getElementById('new-base-course-edit') as HTMLInputElement).value;
        (document.getElementById('new-question-add-title') as HTMLInputElement).value = "";
    }
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Add question</Modal.Title></Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group className="pb-4">
                        <Form.Control type="text"
                                      className="mb-2"
                                      required
                                      id="new-question-add-title"
                        />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={validate}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const EditModal = ({show, onSubmit, onCancel}: {show: boolean, onSubmit: () => void, onCancel: () => void}) => {
    const validate = () => {
        const newTitle = (document.getElementById('new-question-edit-title') as HTMLInputElement).value;
        (document.getElementById('new-question-edit-title') as HTMLInputElement).value = "";
    }
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Edit question</Modal.Title></Modal.Header>
            <Modal.Body>
                <Form>
                    <Form.Group className="pb-4">
                        <Form.Control type="text"
                                      className="mb-2"
                                      required
                                      id="new-base-course-edit"
                        />
                    </Form.Group>
                </Form>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={validate}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const DeleteModal = ({show, onSubmit, onCancel}: {show: boolean, onSubmit: () => void, onCancel: () => void}) => {
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Delete question</Modal.Title></Modal.Header>
            <Modal.Body>
                <p>Removing from the exam.</p>
            </Modal.Body>
            <Modal.Footer>
                <Button variant="secondary" onClick={onCancel}>Back</Button>
                <Button variant="primary" onClick={() => onSubmit()}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

const ExamConfigQuestions = () => {
    let params = useParams();
    const courseName = params.course_name;
    const examId = params.exam_id;
    const {globalState} = useGlobalState();
    let {examConfigState, setExamConfigState}  = useConfigStates();

    const settingLength = examConfigState?.settings ? examConfigState.settings.length : -1;

    const [tags, setTags] = useState<tagProps[]>([]);
    const getTags = useCallback(() => {
        const url = getBackendApiUrl("/courses/" + courseName + "/tags");
        const token = globalState.token;
        return axios.get(url, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getTags()
            .then(response => {
                const data : tagProps[] = response.data.data;
                setTags(data);
            })
    }, [])

    const [addQuestionModalShow, setAddQuestionModalShow] = useState(false);
    const [editQuestionModalShow, setEditQuestionModalShow] = useState(false);
    const [deleteQuestionModalShow, setDeleteQuestionModalShow] = useState(false);
    const [questionInInterest, setQuestionInInterest] = useState<ExamConfigSettingsType>();

    const editQuestionWrapper = (index: number) => {
        if (index >= 0 && index <= settingLength) {
            setQuestionInInterest(examConfigState?.settings[index]);
            setEditQuestionModalShow(true);
        }
    }
    const deleteQuestionWrapper = (index: number) => {
        if (index >= 0 && index <= settingLength) {
            setQuestionInInterest(examConfigState?.settings[index]);
            setDeleteQuestionModalShow(true);
        }
    }

    const addFormCleanUp = () => {
        setAddQuestionModalShow(false);
    }
    const editFormCleanUp = () => {
        setEditQuestionModalShow(false);
    }
    const deleteFormCleanUp = () => {
        setDeleteQuestionModalShow(false);
    }

    const settingsToQuestion = examConfigState?.settings?.map((setting, index) => settingToQuestion(setting, index, tags, editQuestionWrapper, deleteQuestionWrapper));
    return (
        <div>
            <div className="text-end mb-2">
                <Button variant="success" className="me-1" onClick={() => setAddQuestionModalShow(true)}><i className="bi bi-plus-square me-1"/>Add Question</Button>
                {(settingLength >= 2) && (<Button variant="warning"><i className="bi bi-list me-1"/>Change Order</Button>)}
            </div>
            {settingsToQuestion}
            <AddModal show={addQuestionModalShow} onSubmit={() => {}} onCancel={addFormCleanUp} />
            <EditModal show={editQuestionModalShow} onSubmit={() => {}} onCancel={editFormCleanUp} />
            <DeleteModal show={deleteQuestionModalShow} onSubmit={() => {}} onCancel={deleteFormCleanUp} />
        </div>
    )
}

export default ExamConfigQuestions;