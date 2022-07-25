import React from 'react';
import {useState} from 'react';
import {Card, Collapse} from 'react-bootstrap';
import SingleChoice from './questionTemplate/SingleChoice';
import MultipleChoice from './questionTemplate/MultipleChoice';
import SingleBlank from './questionTemplate/SingleBlank';
import MultipleBlank from './questionTemplate/MultipleBlank';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";

const CollapseQuestion = ({questionData} : {questionData: questionDataType}) => {
    const [open, setOpen] = useState(false);

    const subQuestions = questionData.questions.map((subQuestionData: subQuestionDataType) => {
        const key = "Q" + questionData.headerId.toString() + "_sub" + subQuestionData.questionId.toString();
        if (subQuestionData.questionType === "single-choice") return (<SingleChoice key={key} data={subQuestionData} headerId={questionData.headerId.toString()} />);
        if (subQuestionData.questionType === "multiple-choice") return (<MultipleChoice key={key} data={subQuestionData} headerId={questionData.headerId.toString()} />);
        if (subQuestionData.questionType === "single-blank") return (<SingleBlank key={key} data={subQuestionData} headerId={questionData.headerId.toString()} />);
        if (subQuestionData.questionType === "multiple-blank") return (<MultipleBlank key={key} data={subQuestionData} headerId={questionData.headerId.toString()} />);
        return (<></>);
    });

    return (
        <>
            <Card>
                <Card.Header
                    style={{cursor: "pointer"}}
                    onClick={() => setOpen(!open)}
                    aria-controls={`Q${questionData.headerId}`}
                    aria-expanded={open}>
                    {questionData.headerId + ". Question Title"}
                </Card.Header>
                <Collapse in={open}>
                    <div id={`Q${questionData.headerId}`}>
                    <Card.Body>
                        <div dangerouslySetInnerHTML={{__html: questionData.description}}/>
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
