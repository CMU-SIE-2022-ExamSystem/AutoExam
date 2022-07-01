import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";

const MultipleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const checkboxes = data.choices.map((choice: any) => (
        <Form.Check type='checkbox'
            name={`Q${headerId}_sub${data.questionId}`}
            id={`Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`}
            label={choice.content} />
    ));

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {checkboxes}
        </QuestionLayout>
    );
}

export default MultipleChoice;
