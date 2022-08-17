import React, {useState, useEffect} from 'react';
import {Button, Col, Form, InputGroup, Row} from 'react-bootstrap';
import {subQuestionDataType} from '../../../../components/questionTemplate/subQuestionDataType';

interface choiceProps {
    choice_idx: number;
    choice_content: string;
    choice_checked: boolean;
}

const EditChoice = ({type, id, displayIdx, subQuestion, onDelete}: {type: string, id: number, displayIdx: number, subQuestion: subQuestionDataType | undefined | null, onDelete: (id: number) => void}) => {
    const [description, setDescription] = useState("");

    const [choiceIdx, setChoiceIdx] = useState(0);
    const [choiceList, setChoiceList] = useState<choiceProps[]>([]);

    const clearState = () => {
        setChoiceIdx(0);
        setChoiceList([]);
    }

    useEffect(() => {
        clearState();
        subQuestion !== undefined && subQuestion!== null &&
            setDescription(subQuestion.description);
        subQuestion !== undefined && subQuestion!== null && subQuestion.choices[0] !== null &&
            setChoiceIdx(subQuestion.choices[0]?.length);
        subQuestion !== undefined && subQuestion!== null &&
            subQuestion.choices[0]?.map((choice, index) =>
                setChoiceList((prevState) => ([
                    ...prevState,
                    {
                        choice_idx: index,
                        choice_content: choice.content,
                        choice_checked: type === "single_choice" ? subQuestion.solutions[0].includes(choice.choice_id) : subQuestion.solutions[0][0].includes(choice.choice_id)
                    }
                ]))
            );
    }, [subQuestion, type])

    const deleteChoice = (idx: number) => {
        setChoiceList(choiceList.filter((choice) => choice.choice_idx !== idx));
    }

    const choices = choiceList.map((choice, index) => {
        return (
            <Row className="d-flex flex-row align-items-center" key={choice.choice_idx}>
                <Col>
                    <InputGroup className="my-2">
                        <InputGroup.Checkbox name={"sub" + id + "_choices"} defaultChecked={choice.choice_checked}/>
                        <Form.Control id={"sub" + id + "_choice" + index} defaultValue={choice.choice_content}/>
                    </InputGroup>
                </Col>
                <Col xs={1}>
                    <i className="bi-trash" style={{cursor: "pointer"}} onClick={() => deleteChoice(choice.choice_idx)}/>
                </Col>
            </Row>
        );
    });

    return (
        <>
        <Form.Group>
            <Form.Label><h5>{displayIdx + (type === "single_choice" ? ". Single Choice" : ". Multiple Choice")}</h5></Form.Label>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Description</Form.Label>
            <Form.Control id={"sub" + id + "_description"} defaultValue={subQuestion?.description} onChange={(e) => setDescription(e.target.value)}/>
        </Form.Group>

        <Form.Group className="mb-3">
            <Form.Label>Choices</Form.Label><br/>
            <Form.Text>{type === "single_choice" ? "choose all possible answers" : "choose all correct answers"}</Form.Text>
        </Form.Group>

        <div>{choices}</div>

        <div className="mb-3 text-end">
            <Button variant="primary" onClick={() => {setChoiceList([...choiceList, {choice_idx: choiceIdx, choice_content: "", choice_checked: false}]); setChoiceIdx(choiceIdx + 1);}}>Add Choice</Button>
            <Button variant="secondary" className="ms-2" onClick={() => onDelete(id)}>Delete Subquestion</Button>
        </div>
        <hr/>
        </>
    );
}

export default EditChoice;
