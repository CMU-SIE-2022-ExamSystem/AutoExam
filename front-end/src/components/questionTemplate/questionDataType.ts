import { subQuestionDataType } from "./subQuestionDataType";

export default interface questionDataType {
    description: string;
    id: string;
    question_tag: string;
    sub_questions: subQuestionDataType[];
};
