import React from 'react';
import {Form} from 'react-bootstrap';
import {choiceDataType} from "../subQuestionDataType";

// A single checkbox
const OneInMultipleChoice = ({choice, storageKey} : {choice: choiceDataType, storageKey: string}) => {
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='checkbox'
            name={storageKey}
            id={id}
            label={choice.content}
            readOnly
        />
    )
}

// MultipleChoice questions with multiple answers (checkbox)
const MultipleChoiceReadOnly = ({data, storageKey} : {data: choiceDataType[], storageKey: string}) => {
    if (!data) return (<>Bad choices field</>);
    const checkboxes = data.map((choice: any) => {
        let key = `${storageKey}_choice${choice.choiceId}`;
        return (
            <OneInMultipleChoice
                choice={choice}
                key={key}
                storageKey={storageKey}
            />
        )
    });

    return (
        <div id={storageKey}>
            {checkboxes}
        </div>
    );
}

export default MultipleChoiceReadOnly;
