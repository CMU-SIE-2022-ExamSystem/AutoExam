import React from 'react';
import {Badge, Card} from 'react-bootstrap';
import {subQuestionDataType} from "./questionTemplate/subQuestionDataType";
import questionDataType from "./questionTemplate/questionDataType";
import SubQuestion from "./questionTemplate/SubQuestion";

/**
 * The layout of a question
 * @param questionData question information
 * @param questionId the index order of that question
 */
const Question = ({questionData, questionId} : {questionData: questionDataType, questionId: number}) => {
    const subQuestions = questionData.sub_questions.map((subQuestionData: subQuestionDataType, index) => {
        let idx = (index + 1).toString();
        const key = "Q" + questionId + "_sub" + idx;
        return <SubQuestion key={key} data={subQuestionData} headerId={key} displayIdx={index + 1} />
    });

    const questionNumberBadge = (<Badge bg="primary ms-1">{questionData.sub_question_number} questions</Badge>)
    const scoreBadge = (<Badge bg="success ms-1">{questionData.score} points</Badge>);

    return (
        <>
            <br/>
            <Card className="text-start">
                <Card.Header>{questionId + ". " + questionData.question_tag}{questionNumberBadge}{scoreBadge}</Card.Header>
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
