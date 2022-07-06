import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const SingleBlank = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const {value, setValue} = usePersistState("", `Q${headerId}_sub${data.questionId}`)
    let placeholder: string = "";
    if (data.choices.length > 0) {
        placeholder = data.choices[0].content;
    }
    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            <Form.Control type="text"
                          placeholder={placeholder}
                          id={`Q${headerId}_sub${data.questionId}`}
                          className="w-50 mb-2"
                          value={value}
                          onChange={(event) => {
                              const newValue = event.target.value;
                              setValue(newValue);
                          }}
            />
        </QuestionLayout>
    );
}

export default SingleBlank;
