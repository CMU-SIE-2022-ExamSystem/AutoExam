import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";

const SingleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const radios = data.choices.map((choice: any) => (
        <Form.Check type='radio'
            name={`Q${headerId}_sub${data.questionId}`}
            key={`Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`}
            id={`Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`}
            label={choice.content} />
    ));

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {radios}
        </QuestionLayout>
    );
}

export default SingleChoice;
