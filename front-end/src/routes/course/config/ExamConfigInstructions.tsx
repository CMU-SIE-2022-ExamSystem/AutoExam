import React, {useRef} from 'react';
import HTMLEditor from "../../../components/HTMLEditor";

const ExamConfigInstructions = () => {
    const editorRef = useRef<any>(null);
    const log = (): string => {
        if (editorRef.current) {
            console.log(editorRef.current.getContent());
            return editorRef.current.getContent();
        }
        return "";
    };
    return (
        <div>
            <h1 className={"text-start"}>Instructions</h1>
            <HTMLEditor editorRef={editorRef}/>
        </div>
    )
}

export default ExamConfigInstructions;