import React from 'react';
import {Card} from 'react-bootstrap';
import SingleChoice from './questionTemplate/SingleChoice';
import MultipleChoice from './questionTemplate/MultipleChoice';
import SingleBlank from './questionTemplate/SingleBlank';
import MultipleBlank from './questionTemplate/MultipleBlank';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";

const Question = ({questionData} : {questionData: questionDataType}) => {
    const subQuestions = questionData.sub_questions.map((subQuestionData: subQuestionDataType) => {
        const key = "Q" + questionData.id.toString() + "_sub" + subQuestionData.question_id.toString();
        if (subQuestionData.question_type === "single-choice") return (<SingleChoice key={key} data={subQuestionData} headerId={questionData.id.toString()} />);
        if (subQuestionData.question_type === "multiple-choice") return (<MultipleChoice key={key} data={subQuestionData} headerId={questionData.id.toString()} />);
        if (subQuestionData.question_type === "single-blank") return (<SingleBlank key={key} data={subQuestionData} headerId={questionData.id.toString()} />);
        if (subQuestionData.question_type === "multiple-blank") return (<MultipleBlank key={key} data={subQuestionData} headerId={questionData.id.toString()} />);
        return (<></>);
    });

    return (
        <>
            <br/>
            <Card className="text-start">
                <Card.Header>{questionData.id + ". " + questionData.question_tag}</Card.Header>
                <Card.Body className="d-flex flex-column">
                    {/* <Card.Title className="fs-4 fw-bold">{questionData.questionTag}</Card.Title> */}
                    <div dangerouslySetInnerHTML={{__html: questionData.description}}/>
                    {subQuestions}
                </Card.Body>
            </Card>
        </>
    );
}

export default Question;
