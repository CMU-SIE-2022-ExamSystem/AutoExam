import React, {useEffect} from 'react';
import { Form } from 'react-bootstrap';
import QuestionLayout from "./QuestionLayout";
import { subQuestionDataType } from "./subQuestionDataType";
import usePersistState from "../../utils/usePersistState";

const MultipleBlank = ({data, headerId} : {data: subQuestionDataType, headerId: string}) => {
    const {value, setValue} = usePersistState(new Array<string>(data.choices.length).fill(""), `Q${headerId}_sub${data.questionId}`)
    let blanks = data.choices.map((placeholder, index) => {

        return (
            <Form.Control type="text" placeholder={placeholder.content}
                key={`Q${headerId}_sub${data.questionId}_sub${placeholder.choiceId}`}
                id={`Q${headerId}_sub${data.questionId}_sub${placeholder.choiceId}`}
                className="w-50 mb-2"
                value={(value as Array<string>)[index]}
                onChange={(event) => {
                    //console.log([...value], index);
                    const newValue = [...value];
                    newValue[index] = event.target.value;
                    //console.log(newValue);
                    setValue(newValue);
                }}
            />
        )
    });

    return (
        <QuestionLayout questionId={data.questionId.toString()} description={data.description}>
            {blanks}
        </QuestionLayout>
    );
}

export default MultipleBlank;
