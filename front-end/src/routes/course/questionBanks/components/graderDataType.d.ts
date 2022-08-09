export interface blankDataType {
    is_choice: boolean;
    multiple: boolean;
    type: 'string' | 'code';
}

export default interface graderDataType {
    name: string;
    blanks: blankDataType[];
    modules: string[];
    valid: boolean;
    uploaded: boolean;
}
