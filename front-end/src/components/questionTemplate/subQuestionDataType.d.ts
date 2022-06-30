export interface choiceDataType {
    choiceId: string;
    content: string;
}

export interface subQuestionDataType {
    questionType: "single-choice" | "multiple-choice" | "single-blank" | "multiple-blank";
    questionId: number;
    description: string;
    choices: choiceDataType[];
};
