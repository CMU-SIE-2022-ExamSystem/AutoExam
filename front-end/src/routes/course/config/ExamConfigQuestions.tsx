import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {ExamConfigSettingsType, useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {Button, Card, Form, Modal, Row, Table, Col} from "react-bootstrap";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import Question from "../../../components/questionTemplate/readonly/QuestionReadOnly";

interface tagProps {
    id: string;
    name: string;
    course: string;
}

const SettingToQuestion = ({setting, qIndex, tags, editWrapper, deleteWrapper}:{setting: ExamConfigSettingsType, qIndex: number, tags: tagProps[], editWrapper: (arg0: number) => void, deleteWrapper: (arg0: number) => void}) => {
    const myTag = tags.find(tag => tag.id === setting.tag);
    const idLength = setting.id.length;
    const appointed = idLength > 0 ? ("Chosen from " + idLength + " questions") : "Random pick according to tag";
    const subquestionTitles = Array.from({length: setting.sub_question_number}, (value, index) => "sub" + (index + 1).toString());
    const subquestionScores = (
        <Table bordered>
            <thead>
                <tr><th>Questions</th>{subquestionTitles.map((title, index) => <th scope="col" key={"question_" + qIndex + "_subtitle_" + index}>{title}</th>)}</tr>
            </thead>
            <tbody>
                <tr><td>Scores</td>{setting.scores.map((score, index) => <td key={"question_" + qIndex + "_subscore_" + index}>{score}</td>)}</tr>
            </tbody>
        </Table>
    )
    return (
        <Card className="my-3 text-start">
            <Card.Header>
                {qIndex + 1}. {myTag ? myTag.name : setting.title}
            </Card.Header>
            <Card.Body>
                <Card.Text>Question title: {setting.title}</Card.Text>
                <Card.Text>Score: {setting.max_score}</Card.Text>
                <Card.Text># of sub questions: {setting.sub_question_number}</Card.Text>
                {subquestionScores}
                <Card.Text>Generation method: {appointed}</Card.Text>
            </Card.Body>
            <Card.Footer>
                <Button variant="warning" size="sm" className="me-1" onClick={() => editWrapper(qIndex)}><i className="bi bi-pencil-square me-1"/>Edit</Button>
                <Button variant="danger" size="sm" onClick={() => deleteWrapper(qIndex)}><i className="bi bi-trash me-1"/>Delete</Button>
            </Card.Footer>
        </Card>
    );
};

const tagListOptions = (tagList: tagProps[]) => tagList.map(tag => (
    <option key={tag.id} value={tag.id}>{tag.name}</option>
));

const templateExamConfigSetting : ExamConfigSettingsType = {
    id: [], max_score: 0, scores: [], sub_question_number: 0, tag: "", title: ""
};

const QuestionDisplayModal = ({show, toClose, question} : {show: boolean, toClose: () => void, question: questionDataType}) => {
    return (
        <Modal show={show} onHide={toClose} fullscreen>
            <Modal.Header closeButton>
                Question view
            </Modal.Header>
            <Modal.Body>
                <Question questionData={question} questionId={1} />
            </Modal.Body>
        </Modal>
    )
}

const AddModal = ({show, question, setQuestion, onSubmit, onCancel, tagList, pickList}:
                      {show: boolean, question: ExamConfigSettingsType, setQuestion: React.Dispatch<React.SetStateAction<ExamConfigSettingsType>>,
                          onSubmit: () => void, onCancel: () => void, tagList: tagProps[], pickList: questionDataType[]}) => {
    const updateQuestion = (updateTerm: any) => {
        let emptyArray = {};
        if (updateTerm.sub_question_number) {
            emptyArray = {scores: Array.from({length: updateTerm.sub_question_number}, _ => 0)}
        }
        const newQuestion = Object.assign({}, question, updateTerm, emptyArray);
        setQuestion(newQuestion);
        return newQuestion;
    }
    const [badSubQuestionsNumber, setBadSubQuestionsNumber] = useState<boolean>(false);
    const [displayQuestion, setDisplayQuestion] = useState<questionDataType | null>(null);
    const [showDisplayModal, setShowDisplayModal] = useState<boolean>(false);
    const filteredQuestions = ((arr: questionDataType[]) => {
        if (question.sub_question_number > 0) {
            return arr.filter(q => q.sub_question_number === question.sub_question_number);
        } else {
            return arr;
        }
    })(pickList);
    const pickListChecks = filteredQuestions.map(filtered => {
        const label = (
            <span>{filtered.title} ({filtered.sub_question_number}) <i className="bi bi-zoom-in pointer-cursor" onClick={() => {
                setDisplayQuestion(filtered);
                setShowDisplayModal(true);
            }}/></span>
        )
        return (
            <Form.Check name="picked-questions"
                        value={filtered.id}
                        key={filtered.id}
                        label={label}
                        defaultChecked={question.id.includes(filtered.id)}
                        onChange={(e) => {
                            let prevIdList = question.id;
                            let idx = question.id.findIndex(id => id === filtered.id);
                            if (idx >= 0 && !e.target.checked) {
                                prevIdList.splice(idx, 1);
                                updateQuestion({id: prevIdList});
                            }
                            if (idx < 0 && e.target.checked) {
                                prevIdList.push(filtered.id);
                                if (question.sub_question_number === 0) {
                                    console.log(updateQuestion({sub_question_number: filtered.sub_question_number, id: prevIdList}));
                                } else {
                                    updateQuestion({id: prevIdList});
                                }
                            }
                        }}
                        />
        )
    })
    const availableValues = Array.from(new Set(pickList.map(item => item.sub_question_number)).values());

    const pointsFormControls = Array.from({length: question.sub_question_number}, (value, index) => {
        return (
            <React.Fragment key={"SubquestionPoint_" + (index+1)}>
                <Form.Label column xs={3}>Q{index + 1}</Form.Label>
                <Col xs={3}><Form.Control type="number" min="0" step="0.01" value={question.scores[index]}
                                          onChange={(e) => {
                                              let newArray = question.scores;
                                              newArray[index] = parseFloat(e.target.value) || 0;
                                              const totalPoints = newArray.reduce((prev, next) => prev + next, 0);
                                              updateQuestion({scores: newArray, max_score: totalPoints})
                                          }}/></Col>
            </React.Fragment>

        )
    })

    const validate = () => {
        if (!availableValues.includes(question.sub_question_number)) {
            setBadSubQuestionsNumber(true);
            return;
        }
        console.log(question);
        onSubmit();
    }

    let availableValuesText = (<small className={badSubQuestionsNumber ? "text-danger" : ""}>Following values are available: {availableValues.join(",")}.</small>);
    return (
        <>
            <Modal show={show}>
                <Modal.Header><Modal.Title>Add question</Modal.Title></Modal.Header>
                <Form onSubmit={(e) => {e.preventDefault(); validate();}}>
                    <Modal.Body>
                        <Form.Group className="mb-2">
                            <Form.Label>Tag</Form.Label>
                            <Form.Select required id="new-question-add-tag" value={question.tag} onChange={(e) => updateQuestion({tag: e.target.value})}>
                                <option value={""}>Select a tag</option>
                                {tagListOptions(tagList)}
                            </Form.Select>
                        </Form.Group>
                        <Form.Group className="mb-2">
                            <Form.Label>Title</Form.Label>
                            <Form.Control type="text"
                                          className="mb-2"
                                          required
                                          id="new-question-add-title"
                                          value={question.title}
                                          onChange={(e) => {updateQuestion({title: e.target.value})}}
                            />
                        </Form.Group>
                        <Form.Group className="mb-2">
                            <Form.Label>Number of Subquestions</Form.Label>
                            <Form.Control type="number" min="0" step="1" value={question.sub_question_number}
                                          onChange={(e) => {
                                              setBadSubQuestionsNumber(false);
                                              updateQuestion({sub_question_number: parseInt(e.target.value) || 0, id: []});
                                          }}
                                          isInvalid={badSubQuestionsNumber}
                            />
                            {question.tag && availableValuesText}
                        </Form.Group>
                        <Form.Group className="mb-2">
                            <Form.Label>Appointed Questions</Form.Label>
                            {pickListChecks}
                            {pickListChecks.length === 0 && <div className="text-secondary fst-italic"><small>No available questions, please change your tag or number of subquestions.</small></div>}
                        </Form.Group>
                        {
                            question.sub_question_number !== 0 &&
                            <Form.Group className="mb-2" as={Row}>
                                <Form.Label column xs={"12"}>Point Distributions</Form.Label>
                                {pointsFormControls}
                                <small>Total Points: {question.max_score}</small>
                            </Form.Group>
                        }
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={onCancel}>Back</Button>
                        <Button variant="primary" type="submit">Submit</Button>
                    </Modal.Footer>
                </Form>
            </Modal>
            {
                displayQuestion && <QuestionDisplayModal show={showDisplayModal} toClose={() => setShowDisplayModal(false)} question={displayQuestion} />
            }
        </>
    );
}

const DeleteModal = ({show, oldQuestion, onSubmit, onCancel}: {show: boolean, oldQuestion: ExamConfigSettingsType, onSubmit: () => void, onCancel: () => void}) => {
    return (
        <Modal show={show}>
            <Modal.Header><Modal.Title>Delete question</Modal.Title></Modal.Header>
            <Modal.Body>
                <p>Removing '{oldQuestion.title}' from the exam.</p>
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
    const {globalState} = useGlobalState();
    let {examConfigState, setExamConfigState}  = useConfigStates();

    const updateState = (updateTerm: any) => {
        const newState = Object.assign({}, examConfigState, updateTerm)
        setExamConfigState(newState);
    }

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
    const [qIIindex, setQIIindex] = useState<number>(0);
    const [questionInInterest, setQuestionInInterest] = useState<ExamConfigSettingsType>(templateExamConfigSetting);

    const addQuestionWrapper = () => {
        setQuestionInInterest(templateExamConfigSetting);
        setQIIindex(-1);
        setAddQuestionModalShow(true);
    }

    const editQuestionWrapper = (index: number) => {
        if (index >= 0 && index <= settingLength) {
            if (examConfigState && examConfigState.settings) {
                setQuestionInInterest(examConfigState.settings[index]);
                setQIIindex(index);
            }
            setEditQuestionModalShow(true);
        }
    }
    const deleteQuestionWrapper = (index: number) => {
        if (index >= 0 && index <= settingLength) {
            if (examConfigState && examConfigState.settings) {
                setQuestionInInterest(examConfigState.settings[index]);
                setQIIindex(index);
            }
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


    const addFormHandler = () => {
        let currentSettings = examConfigState?.settings;
        if (!currentSettings) return;
        currentSettings.push(questionInInterest);
        updateState({settings: currentSettings});
        addFormCleanUp();
    }

    const editFormHandler = () => {
        let currentSettings = examConfigState?.settings;
        if (!currentSettings) return;
        currentSettings.splice(qIIindex, 1, questionInInterest);
        updateState({settings: currentSettings});
        editFormCleanUp();
    }

    const deleteFormHandler = () => {
        let currentSettings = examConfigState?.settings;
        if (!currentSettings) return;
        currentSettings.splice(qIIindex, 1);
        updateState({settings: currentSettings});
        deleteFormCleanUp();
    }

    const [queryQuestionList, setQueryQuestionList] = useState<questionDataType[]>([]);

    const getQuestionList = useCallback(async () => {
        if (questionInInterest.tag.length > 0) {
            const url = getBackendApiUrl("/courses/" + courseName + "/questions?tag_id=" + questionInInterest.tag);
            const token = globalState.token;
            const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
            setQueryQuestionList(result.data.data);
        }
    }, [questionInInterest.tag])
    useEffect(() => {
        getQuestionList().catch();
    }, [questionInInterest.tag])

    const settingsToQuestion = examConfigState?.settings?.map((setting, index) => (
        <SettingToQuestion key={"exam_config_question_" + index} setting={setting} qIndex={index} tags={tags}  editWrapper={editQuestionWrapper} deleteWrapper={deleteQuestionWrapper}/>
    ));
    return (
        <div className="mb-3">
            <div className="text-end mb-2">
                <Button variant="success" className="me-1" onClick={addQuestionWrapper}><i className="bi bi-plus-square me-1"/>Add Question</Button>
                {(settingLength >= 2) && (<Button variant="warning"><i className="bi bi-list me-1"/>Change Order</Button>)}
            </div>
            {settingsToQuestion}
            <AddModal show={addQuestionModalShow} question={questionInInterest} setQuestion={setQuestionInInterest} onSubmit={addFormHandler} onCancel={addFormCleanUp} tagList={tags} pickList={queryQuestionList}/>
            <AddModal show={editQuestionModalShow} question={questionInInterest} setQuestion={setQuestionInInterest} onSubmit={editFormHandler} onCancel={editFormCleanUp} tagList={tags} pickList={queryQuestionList}/>
            <DeleteModal show={deleteQuestionModalShow} oldQuestion={questionInInterest} onSubmit={deleteFormHandler} onCancel={deleteFormCleanUp} />
        </div>
    )
}

export default ExamConfigQuestions;