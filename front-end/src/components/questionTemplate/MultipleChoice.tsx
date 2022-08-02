import React from 'react';
import {Form} from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import {choiceDataType, subQuestionDataType} from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

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
                console.log(newValue);
                setValue(newValue);
            }} />
    )
}

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
