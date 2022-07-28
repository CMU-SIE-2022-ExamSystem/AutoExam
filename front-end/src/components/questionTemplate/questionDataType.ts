import { subQuestionDataType } from "./subQuestionDataType";

export default interface questionDataType {
    id: number;
    question_tag: string;
    description: string;
    questions: subQuestionDataType[];
};
