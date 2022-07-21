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
            <HTMLEditor editorRef={editorRef}/>
            <button onClick={log}>Log editor content</button>
        </div>
    )
}

export default ExamConfigInstructions;