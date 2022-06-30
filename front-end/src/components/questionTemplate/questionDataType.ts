import { subQuestionDataType } from "./subQuestionDataType";

export default interface questionDataType {
    headerId: number;
    questionTag: string;
    description: string;
    questions: subQuestionDataType[];
};
