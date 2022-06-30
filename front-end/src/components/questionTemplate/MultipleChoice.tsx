import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const MultipleChoice = ({data} : {data: subQuestionDataType}) => {
    const checkboxes = data.choices.map((data: any) => (
        <Form.Check type='checkbox' id='default-checkbox' label={data.content} />
    ));

    return (
        <>
            <QuestionLayout questionId={data.questionId} description={data.description}>
                {checkboxes}
            </QuestionLayout>
        </>
    );
}

export default MultipleChoice;
