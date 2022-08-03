import React from 'react';
import {Form} from 'react-bootstrap';
import {choiceDataType} from "../subQuestionDataType";

const OneInMultipleChoice = ({choice, storageKey, value} : {choice: choiceDataType, storageKey: string, value: string}) => {
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='checkbox'
            name={storageKey}
            id={id}
            label={choice.content}
            defaultChecked={value.includes(choice.choice_id)}
            readOnly
        />
    )
}

const MultipleChoiceReadOnly = ({data, storageKey, value} : {data: choiceDataType[], storageKey: string, value: string}) => {
    if (!data) return (<>Bad choices field</>);
    const checkboxes = data.map((choice: any) => {
        let key = `${storageKey}_choice${choice.choiceId}`;
        return (
            <OneInMultipleChoice
                choice={choice}
                key={key}
                storageKey={storageKey}
                value={value}
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
