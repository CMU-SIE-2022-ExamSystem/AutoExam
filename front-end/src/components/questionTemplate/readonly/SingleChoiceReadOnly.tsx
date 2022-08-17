import { Form } from 'react-bootstrap';
import { choiceDataType } from "../subQuestionDataType";
import React from "react";

// A single radio checkbox
const OneInSingleChoice = ({choice, storageKey} : {choice: choiceDataType, storageKey: string}) => {
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='radio'
                    name={storageKey}
                    id={id}
                    label={choice.content}
                    //defaultChecked={value.includes(choice.choice_id)}
                    readOnly />
    )
}

// MultipleChoice questions with a single answer (radio)
const SingleChoiceReadOnly = ({data, storageKey} : {data: choiceDataType[], storageKey: string}) => {
    const radios = data.map((choice) => {
        if (!choice) return (<>Bad Question</>);
        let key = `${storageKey}_choice${choice.choice_id}`;
        return (
            <OneInSingleChoice
                choice={choice}
                key={key}
                storageKey={storageKey}
            />
        )
    });

    return (
        <div id={storageKey}>
            {radios}
        </div>
    );
}

export default SingleChoiceReadOnly;
