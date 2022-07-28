export interface choiceDataType {
    choice_id: string;
    content: string;
}

export interface subQuestionDataType {
    question_type: "single-choice" | "multiple-choice" | "single-blank" | "multiple-blank";
    question_id: number;
    description: string;
    choices: choiceDataType[];
};
