import React from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { choiceDataType, subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const OneInMultipleChoice = ({choice, storageKey} : {choice: choiceDataType, storageKey: string}) => {
    const {value, setValue, getValue} = usePersistState('', storageKey);
    const id = storageKey + "_choice" + choice.choiceId;
    return (
        <Form.Check type='checkbox'
            name={storageKey}
            id={id}
            label={choice.content}
            defaultChecked={value.includes(choice.choiceId)}
            onChange={(event) => {
                let newValue = "";
                const prevValue = getValue();
                if (prevValue.includes(choice.choiceId) && !event.target.checked) {
                    newValue = prevValue.replace(choice.choiceId, "");
                }
                if (!prevValue.includes(choice.choiceId) && event.target.checked) {
                    newValue = prevValue.concat(choice.choiceId);
                }
                console.log(newValue);
                setValue(newValue);
            }} />
    )
}

const MultipleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const checkboxes = data.choices.map((choice: any) => {
        let key = `Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`;
        let storageKey = `Q${headerId}_sub${data.questionId}`;
        return (
            <OneInMultipleChoice
                choice={choice}
                key={key}
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