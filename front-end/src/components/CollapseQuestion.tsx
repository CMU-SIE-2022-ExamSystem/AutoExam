import React from 'react';
import {useState} from 'react';
import {Card, Collapse} from 'react-bootstrap';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";
import BlankWithSolution from './questionTemplate/BlankWithSolution';
import ChoiceWithSolution from './questionTemplate/ChoiceWithSolution';

const CollapseQuestion = ({question} : {question: questionDataType}) => {
    const [open, setOpen] = useState(false);

    const subQuestions = question.sub_questions.map((subQuestion: subQuestionDataType, index) => {
        if (subQuestion.grader === "single_blank") return (<BlankWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        if (subQuestion.grader === "single_choice" || subQuestion.grader === "multiple_choice") return (<ChoiceWithSolution key={index + 1} index = {index + 1} subQuestion={subQuestion}/>);
        return (<></>);
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
