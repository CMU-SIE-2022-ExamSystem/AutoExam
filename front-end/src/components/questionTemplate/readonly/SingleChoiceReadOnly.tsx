import { Form } from 'react-bootstrap';
import { choiceDataType } from "../subQuestionDataType";
import React from "react";

const OneInSingleChoice = ({choice, storageKey, value} : {choice: choiceDataType, storageKey: string, value: string}) => {
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='radio'
                    name={storageKey}
                    id={id}
                    label={choice.content}
                    defaultChecked={value.includes(choice.choice_id)}
                    readOnly />
    )
}

const SingleChoiceReadOnly = ({data, storageKey, value} : {data: choiceDataType[], storageKey: string, value: string}) => {
    const radios = data.map((choice) => {
        if (!choice) return (<>Bad Question</>);
        let key = `${storageKey}_choice${choice.choice_id}`;
        return (
            <OneInSingleChoice
                choice={choice}
                key={key}
                storageKey={storageKey}
                value={value}
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
