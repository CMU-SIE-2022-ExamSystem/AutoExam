import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { choiceDataType, subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const OneInMultipleChoice = ({choice, name, storageKey} : {choice: choiceDataType, name: string, storageKey: string}) => {
    const {value, setValue} = usePersistState("", storageKey);
    return (
        <Form.Check type='checkbox'
            name={name}
            key={storageKey}
            id={storageKey}
            label={choice.content}
            value={value}
            defaultChecked={false}
            onChange={(event) => {
                const newValue = choice.choiceId;
                event.target.checked ? setValue(newValue) : setValue("");
            }} />
    )
}

const MultipleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const checkboxes = data.choices.map((choice: any) => {
        let storageKey = `Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`;
        return (
            <OneInMultipleChoice
                choice={choice}
                name={`Q${headerId}_sub${data.questionId}`}
                key={storageKey}
                storageKey={storageKey} />
        )
    });

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {checkboxes}
        </QuestionLayout>
    );
}

export default MultipleChoice;
