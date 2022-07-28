import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { choiceDataType, subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

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

const SingleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const radios = data.choices.map((choice) => {
        let key = `Q${headerId}_sub${data.question_id}_choice${choice.choice_id}`;
        let storageKey = `Q${headerId}_sub${data.question_id}`;
        return (
            <OneInSingleChoice
                choice={choice}
                key={key}
                storageKey={storageKey} />
        )
    });

    return (
        <QuestionLayout questionId={data.question_id.toString()} description={data.description}>
            {radios}
        </QuestionLayout>
    );
}

export default SingleChoice;
