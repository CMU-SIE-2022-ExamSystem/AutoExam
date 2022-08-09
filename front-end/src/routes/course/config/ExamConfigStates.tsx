import React, {createContext, useState} from "react";

export interface ExamConfigGeneralProperties {
    category_name: string;
    description: string;
    end_at: string;
    grading_deadline: string;
    max_submissions: number;
    name: string;
    start_at: string;
}

export interface ExamConfigSettingsType {
    id: string[];
    max_score: number;
    scores: number[];
    sub_question_number: number;
    tag: string;
    title: string;
}

export interface ExamConfigStates {
    course: string;
    draft: boolean;
    general: ExamConfigGeneralProperties;
    id: string;
    settings: ExamConfigSettingsType[];
}

export interface ExamConfigStatesContextType {
    examConfigState: ExamConfigStates | null;
    setExamConfigState: React.Dispatch<React.SetStateAction<ExamConfigStates | null>>;
}

const ExamConfigContext = createContext<ExamConfigStatesContextType>(null!);

export const ExamConfigStateProvider = ({children} : {children: React.ReactNode}) => {
    let [examConfigState, setExamConfigState] = useState<ExamConfigStates | null>(null);
    let value = {examConfigState, setExamConfigState};
    return <ExamConfigContext.Provider value={value}>{children}</ExamConfigContext.Provider>;
}

export const useConfigStates = () => React.useContext(ExamConfigContext);
