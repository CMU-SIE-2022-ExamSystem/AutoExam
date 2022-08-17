import React from 'react';
import {Form} from 'react-bootstrap';
import {choiceDataType} from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

/**
 * Display a single checkbox.
 * @param choice The choice contents
 * @param storageKey ID of the component
 */
const OneInMultipleChoice = ({choice, storageKey} : {choice: choiceDataType, storageKey: string}) => {
    const {value, setValue, getValue} = usePersistState('', storageKey);
    const id = storageKey + "_choice" + choice.choice_id;
    return (
        <Form.Check type='checkbox'
            name={storageKey}
            id={id}
            label={choice.content}
            defaultChecked={value.includes(choice.choice_id)}
            onChange={(event) => {
                let newValue = "";
                const prevValue = getValue();
                if (prevValue.includes(choice.choice_id) && !event.target.checked) {
                    newValue = prevValue.replace(choice.choice_id, "");
                }
                if (!prevValue.includes(choice.choice_id) && event.target.checked) {
                    newValue = prevValue.concat(choice.choice_id);
                }
                setValue(newValue);
            }} />
    )
}

/**
 * A multiple choice question component (with multiple answers), with storage.
 * @param data
 * @param storageKey The ID of the question
 * @param displayIdx The index of the question (not used)
 */
const MultipleChoice = ({data, storageKey, displayIdx} : {data: choiceDataType[], storageKey: string, displayIdx: number}) => {
    if (!data) return (<>Bad choices field</>);
    const checkboxes = data.map((choice: any) => {
        let key = `${storageKey}_choice${choice.choiceId}`;
        return (
            <OneInMultipleChoice
                choice={choice}
                key={key}
                storageKey={storageKey} />
        )
    });

    return (
        <div id={storageKey}>
            {/*<Form.Label>({displayIdx}).</Form.Label>*/}
            {checkboxes}
        </div>
    );
}

export default MultipleChoice;
