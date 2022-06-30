import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const MultipleChoice = ({data} : {data: subQuestionDataType}) => {
    const checkboxes = data.choices.map((choice: any) => (
        <Form.Check type='checkbox' id={choice.choiceId} label={choice.content} />
    ));

    return (
        <QuestionLayout questionId={data.questionId} description={data.description}>
            {checkboxes}
        </QuestionLayout>
    );
}

export default MultipleChoice;
