import {subQuestionDataType} from "../subQuestionDataType";
import QuestionLayout from "../QuestionLayout";
import React from "react";
import MultipleChoiceReadOnly from "./MultipleChoiceReadOnly";
import SingleChoiceReadOnly from "./SingleChoiceReadOnly";
import BlankReadOnly from "./BlankReadOnly";
import CodeReadOnly from "./CodeReadOnly";

const SubQuestionReadOnly = ({data, headerId, displayIdx, value} : {data: subQuestionDataType, headerId: string, displayIdx: number, value: string[]}) => {

    let blanks = data.blanks.map((blank, index) => {
        let choices = data.choices[index];
        let storageKey = `${headerId}_sub${index + 1}`;
        // Detect type
        if (choices !== null) {
            // Choices type
            if (blank.multiple) {
                return <MultipleChoiceReadOnly key={storageKey} data={choices} storageKey={storageKey} value={value[index]} />
            } else {
                return <SingleChoiceReadOnly key={storageKey} data={choices} storageKey={storageKey} value={value[index]} />
            }
        } else {
            // Blanks type, check blank string
            if (blank.type === 'string') {
                return <BlankReadOnly key={storageKey} storageKey={storageKey} value={value[index]}/>
            } else if (blank.type === 'code') {
                return <CodeReadOnly key={storageKey} storageKey={storageKey} value={value[index]}/>
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

export default SubQuestionReadOnly;
