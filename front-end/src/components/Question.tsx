import React from 'react';
import {Card} from 'react-bootstrap';
import SingleChoice from './questionTemplate/SingleChoice';
import MultipleChoice from './questionTemplate/MultipleChoice';
import SingleBlank from './questionTemplate/SingleBlank';
import MultipleBlank from './questionTemplate/MultipleBlank';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";

const Question = ({questionData} : {questionData: questionDataType}) => {
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
            <br/>
            <Card className="text-start">
                <Card.Header>{questionData.headerId + ". " + questionData.questionTag}</Card.Header>
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
