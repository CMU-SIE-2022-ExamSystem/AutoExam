import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import subQuestionDataType from "./subQuestionDataType";

const MultipleBlank = ({data} : {data: subQuestionDataType}) => {
    let blanks = data.choices.map(() => (
        <><Form.Control type="text" /><br/></>
    ));

    return (
        <QuestionLayout questionId={data.questionId} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
