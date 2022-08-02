import { subQuestionDataType } from "./subQuestionDataType";

export default interface questionDataType {
    description: string;
    question_tag: string;
    sub_questions: subQuestionDataType[];
    sub_question_number: number;
    title: string;
    score: number;
};
