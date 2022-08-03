import React from 'react';
import {useState} from 'react';
import {Card, Collapse} from 'react-bootstrap';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";

const CollapseQuestion = ({question} : {question: questionDataType}) => {
    const [open, setOpen] = useState(false);

    const subQuestions = question.sub_questions.map((subQuestion: subQuestionDataType, index) => {
        // if (subQuestion.grader === "multiple-choice") return (<MultipleChoice key={key} data={subQuestion} headerId={question.id.toString()} />);
        // if (subQuestion.grader === "multiple-blank") return (<MultipleBlank key={key} data={subQuestion} headerId={question.id.toString()} />);
        // if (subQuestion.grader === "multiple-blank") return (<MultipleBlank key={key} data={subQuestion} headerId={question.id.toString()} />);
        return (<></>);
    });

    return (
        <>
            <Card>
                <Card.Header
                    style={{cursor: "pointer"}}
                    onClick={() => setOpen(!open)}
                    aria-expanded={open}>
                    {"Question Title"}
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
