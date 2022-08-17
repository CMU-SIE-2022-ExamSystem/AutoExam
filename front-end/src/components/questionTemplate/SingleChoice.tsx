import { Form } from 'react-bootstrap';
import { choiceDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";
import React from "react";

/**
 * Display a single radio checkbox.
 * @param choice The choice contents
 * @param storageKey ID of the component
 */
const OneInSingleChoice = ({choice, storageKey} : {choice: choiceDataType, storageKey: string}) => {
    const {value, setValue} = usePersistState("", storageKey);
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='radio'
            name={storageKey}
            id={id}
            label={choice.content}
            defaultChecked={value.includes(choice.choice_id)}
            onChange={(event) => {
                const newValue = choice.choice_id;
                event.target.checked ? setValue(newValue) : setValue("");
            }} />
    )
}

/**
 * A multiple choice question component (with only one answer), with storage.
 * @param data
 * @param storageKey The ID of the question
 * @param displayIdx The index of the question (not used)
 */
const SingleChoice = ({data, storageKey, displayIdx} : {data: choiceDataType[], storageKey: string, displayIdx: number}) => {
    const radios = data.map((choice) => {
        if (!choice) return (<>Bad Question</>);
        let key = `${storageKey}_choice${choice.choice_id}`;
        return (
            <OneInSingleChoice
                choice={choice}
                key={key}
                storageKey={storageKey} />
        )
    });

    return (
        <div id={storageKey}>
            {/*<Form.Label>({displayIdx}).</Form.Label>*/}
            {radios}
        </div>
    );
}

export default SingleChoice;
