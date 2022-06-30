import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";

const MultipleBlank = ({data} : {data: subQuestionDataType}) => {
    let blanks = data.choices.map((placeholder) => (
        <Form.Control type="text" placeholder={placeholder.content} key={data.questionId + "-" + placeholder.choiceId} className="w-50 mb-2" />
    ));

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
