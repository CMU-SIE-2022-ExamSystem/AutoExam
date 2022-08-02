import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { choiceDataType, subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";
import React from "react";

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
                console.log(newValue);
                event.target.checked ? setValue(newValue) : setValue("");
            }} />
    )
}

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
