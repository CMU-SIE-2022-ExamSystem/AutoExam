import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const MultipleBlank = ({data} : {data: subQuestionDataType}) => {
    let blanks = data.choices.map((placeholder) => (
        <Form.Control type="text" placeholder={placeholder} />
    ));

    return (
        <QuestionLayout index={data.index} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
