import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const OneInMultipleBlank = ({placeholder, index, storageKey} : {placeholder: string, index: number, storageKey: string}) => {
    const {value, setValue} = usePersistState("", storageKey);
    return (
        <Form.Control type="text" placeholder={placeholder}
                      id={storageKey}
                      className="w-50 mb-2"
                      value={value}
                      onChange={(event) => {
                          const newValue = event.target.value;
                          setValue(newValue);
                      }}
        />
    )
}

const MultipleBlank = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    let blanks = data.choices.map((placeholder, index) => {
        let storageKey = `Q${headerId}_sub${data.question_id}_sub${placeholder.choice_id}`;
        return (
            <OneInMultipleBlank
                placeholder={placeholder.content}
                index={index}
                key={storageKey}
                storageKey={storageKey}
            />

        )
    });

    return (
        <QuestionLayout questionId={data.question_id.toString()} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
