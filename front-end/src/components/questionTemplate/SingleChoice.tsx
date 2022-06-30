import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";

const SingleChoice = ({data} : {data: subQuestionDataType}) => {
    const radios = data.choices.map((data: any) => (
        <Form.Check type='radio' id='default-radio' label={data.content} />
    ));

    return (
        <>
            <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
                {radios}
            </QuestionLayout>
        </>
    );
}

export default SingleChoice;
