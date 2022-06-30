import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const SingleChoice = ({data} : {data: subQuestionDataType}) => {
    return (
        <>
            <QuestionLayout questionId={data.questionId} description={data.description}>
                <Form.Control type="text" />
            </QuestionLayout>
        </>
    );
}

export default SingleChoice;
