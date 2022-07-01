import React from 'react';
import {Card} from 'react-bootstrap';
import SingleChoice from './questionTemplate/SingleChoice';
import MultipleChoice from './questionTemplate/MultipleChoice';
import SingleBlank from './questionTemplate/SingleBlank';
import MultipleBlank from './questionTemplate/MultipleBlank';

const Question = ({questionData} : {questionData: any}) => {
    const subQuestions = questionData.questions.map((subQuestionData: any) => {
        if (subQuestionData.questionType === "single-choice") return (<SingleChoice data={subQuestionData} headerId={questionData.headerId} />);
        if (subQuestionData.questionType === "multiple-choice") return (<MultipleChoice data={subQuestionData} headerId={questionData.headerId} />);
        if (subQuestionData.questionType === "single-blank") return (<SingleBlank data={subQuestionData} headerId={questionData.headerId} />);
        if (subQuestionData.questionType === "multiple-blank") return (<MultipleBlank data={subQuestionData} headerId={questionData.headerId} />);
    });

    return (
        <>
            <br/>
            <Card className="text-start">
                <Card.Header>{questionData.headerId + ". " + questionData.questionTag}</Card.Header>
                <Card.Body className="d-flex flex-column">
                    {/* <Card.Title className="fs-4 fw-bold">{questionData.questionTag}</Card.Title> */}
                    <Card.Text><div dangerouslySetInnerHTML={{__html: questionData.description}}/></Card.Text>
                    <Card.Text>{subQuestions}</Card.Text>
                </Card.Body>
            </Card>
        </>
    );
}

export default Question;
