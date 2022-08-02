import {subQuestionDataType} from "./subQuestionDataType";
import QuestionLayout from "./QuestionLayout";
import React from "react";
import MultipleChoice from "./MultipleChoice";
import SingleChoice from "./SingleChoice";
import BlankInput from "./BlankInput";
import CodeInput from "./CodeInput";

const SubQuestion = ({data, headerId, displayIdx} : {data: subQuestionDataType, headerId: string, displayIdx: number}) => {

    let blanks = data.blanks.map((blank, index) => {
        let choices = data.choices[index];
        let storageKey = `${headerId}_sub${index + 1}`;
        // Detect type
        if (choices !== null) {
            // Choices type
            if (blank.multiple) {
                return <MultipleChoice key={storageKey} data={choices} storageKey={storageKey} displayIdx={index + 1} />
            } else {
                return <SingleChoice key={storageKey} data={choices} storageKey={storageKey} displayIdx={index + 1} />
            }
        } else {
            // Blanks type, check blank string
            if (blank.type === 'string') {
                return <BlankInput key={storageKey} storageKey={storageKey} displayIdx={index + 1}/>
            } else if (blank.type === 'code') {
                return <CodeInput key={storageKey} storageKey={storageKey} displayIdx={index + 1}/>
            }
        }
        return (
            <p key={storageKey}>Bad blank info</p>
        )
    });

    return (
        <QuestionLayout displayIdx={displayIdx} questionId={headerId} description={data.description} score={data.score}>
            {blanks}
        </QuestionLayout>
    );
}

export default SubQuestion;
