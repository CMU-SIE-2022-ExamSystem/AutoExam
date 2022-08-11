import React, {useEffect, useState} from 'react';
import HTMLEditor from "../../../components/HTMLEditor";
import {useConfigStates} from "./ExamConfigStates";

const ExamConfigInstructions = () => {
    let {examConfigState, setExamConfigState} = useConfigStates();
    const [instructions, setInstructions]= useState<string>(examConfigState?.general.description || "");

    useEffect(() => {
        setInstructions(examConfigState?.general.description || "");
    }, [examConfigState?.general.description])

    const updateState = (updateTerm: any) => {
        const newState = Object.assign({}, examConfigState, updateTerm)
        setExamConfigState(newState);
    }

    const updateGeneral = (updateTerm: any) => {
        updateState({general: Object.assign({}, examConfigState?.general, updateTerm)})
    }

    const updateInstructions = (newInstruction: string) => {
        setInstructions(newInstruction);
        updateGeneral({description: newInstruction});
    }

    return (
        <div className={"mb-3"}>
            <h1 className={"text-start"}>Instructions</h1>
            <HTMLEditor init={instructions} update={updateInstructions}/>
        </div>
    )
}

export default ExamConfigInstructions;