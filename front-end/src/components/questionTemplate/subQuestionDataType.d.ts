export interface choiceDataType {
    choice_id: string;
    content: string;
}

export interface blankDataType {
    type: 'string' | 'code';
    multiple: boolean;
}

export interface subQuestionDataType {
    blanks: blankDataType[];
    description: string;
    choices: (choiceDataType[] | null)[];
    score: number;
    grader: string;
    solutions: (string[])[];
}
