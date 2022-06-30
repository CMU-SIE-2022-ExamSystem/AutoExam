import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const SingleBlank = ({data} : {data: subQuestionDataType}) => {
    return (
        <QuestionLayout questionId={data.questionId} description={data.description}>
            <Form.Control type="text" />
        </QuestionLayout>
    );
}

export default SingleBlank;
