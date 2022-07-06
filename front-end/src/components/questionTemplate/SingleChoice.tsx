import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { choiceDataType, subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const OneInSingleChoice = ({choice, name, storageKey} : {choice: choiceDataType, name: string, storageKey: string}) => {
    const {value, setValue} = usePersistState("", storageKey);
    return (
        <Form.Check type='radio'
            name={name}
            key={storageKey}
            id={storageKey}
            label={choice.content}
            value = {value}
            defaultChecked={false}
            onChange={(event) => {
                const newValue = choice.choiceId;
                event.target.checked ? setValue(newValue) : setValue("");
            }} />
    )
}

const SingleChoice = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const radios = data.choices.map((choice) => {
        let storageKey = `Q${headerId}_sub${data.questionId}_choice${choice.choiceId}`;
        return (
            <OneInSingleChoice
                choice={choice}
                name={`Q${headerId}_sub${data.questionId}`}
                key={storageKey}
                storageKey={storageKey} />
        )
    });

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {radios}
        </QuestionLayout>
    );
}

export default SingleChoice;
