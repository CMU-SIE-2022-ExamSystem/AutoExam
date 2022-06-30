import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import {choiceDataType, subQuestionDataType} from "./subQuestionDataType";

const SingleBlank = ({data} : {data: subQuestionDataType}) => {
    let placeholder: string = "";
    if (data.choices.length > 0) {
        placeholder = data.choices[0].content;
    }
    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            <Form.Control type="text" placeholder={placeholder} className="w-50 mb-2"/>
        </QuestionLayout>
    );
}

export default SingleBlank;
