import React, {useState} from 'react';
import {Button, Card, Col, Collapse, Row} from 'react-bootstrap';
import {subQuestionDataType} from "../../../../components/questionTemplate/subQuestionDataType";
import questionDataType from "../../../../components/questionTemplate/questionDataType";
import BlankWithSolution from './BlankWithSolution';
import ChoiceWithSolution from './ChoiceWithSolution';
import CustomizedWithSolution from './CustomizedWithSolution';

const CollapseQuestion = ({question, setQuestion, setDeleteShow, setEditShow} : {question: questionDataType, setQuestion: any, setDeleteShow: any, setEditShow: any}) => {
    const [open, setOpen] = useState(false);

    const subQuestions = question.sub_questions.map((subQuestion: subQuestionDataType, index) => {
        if (subQuestion.grader === "single_blank") return (<BlankWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        if (subQuestion.grader === "single_choice" || subQuestion.grader === "multiple_choice") return (<ChoiceWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        return (<CustomizedWithSolution key={index + 1} index={index + 1} subQuestion={subQuestion}/>);
    });

    return (
        <>
            <Card>
                <Card.Header style={{cursor: "pointer"}} onClick={() => setOpen(!open)} aria-expanded={open}>
                    <Row>
                        <Col>{question.title}</Col>
                        <Col>{question.hidden && <div className="text-end">(soft deleted)</div>}</Col>
                    </Row>
                </Card.Header>
                
                <Collapse in={open}>
                    <div>
                        <div className="text-end my-3 me-3">
                            <Button variant="success" onClick={() => {setQuestion(question); setEditShow(true)}}>Edit</Button>
                            <Button variant="secondary" className="ms-2" onClick={() => {setQuestion(question); setDeleteShow(true)}}>Delete</Button>
                        </div>
                        <Card.Body>
                            <div dangerouslySetInnerHTML={{__html: question.description}}/>
                            {subQuestions}
                        </Card.Body>
                    </div>
                </Collapse>
            </Card>
            <br/>
        </>
    );
}

export default CollapseQuestion;
