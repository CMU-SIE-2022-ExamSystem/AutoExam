import {subQuestionDataType} from "../subQuestionDataType";
import QuestionLayout from "../QuestionLayout";
import React from "react";
import MultipleChoiceReadOnly from "./MultipleChoiceReadOnly";
import SingleChoiceReadOnly from "./SingleChoiceReadOnly";
import BlankReadOnly from "./BlankReadOnly";
import CodeReadOnly from "./CodeReadOnly";

const SubQuestionReadOnly = ({data, headerId, displayIdx} : {data: subQuestionDataType, headerId: string, displayIdx: number}) => {

    let blanks = data.blanks.map((blank, index) => {
        let choices = data.choices[index];
        let storageKey = `${headerId}_sub${index + 1}`;
        // Detect type
        if (choices !== null) {
            // Choices type
            if (blank.multiple) {
                return <MultipleChoiceReadOnly key={storageKey} data={choices} storageKey={storageKey}/>
            } else {
                return <SingleChoiceReadOnly key={storageKey} data={choices} storageKey={storageKey}/>
            }
        } else {
            // Blanks type, check blank string
            if (blank.type === 'string') {
                return <BlankReadOnly key={storageKey} storageKey={storageKey}/>
            } else if (blank.type === 'code') {
                return <CodeReadOnly key={storageKey} storageKey={storageKey}/>
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
