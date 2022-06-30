import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";

const MultipleBlank = ({data, index} : {data: subQuestionDataType, index: string}) => {
    let blanks = data.choices.map((placeholder) => (
        <Form.Control type="text" placeholder={placeholder.content} />
    ));

    return (
        <QuestionLayout index={index} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
