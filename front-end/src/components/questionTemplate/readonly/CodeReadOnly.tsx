import {Form} from "react-bootstrap";
import React, {useState} from "react";
import CodeEditor from "@uiw/react-textarea-code-editor";

// List of languages supported, see documentation of react-textarea-code-editor
const listOfLanguages = ["c", "cpp", "java", "javascript", "plaintext", "python"]

/**
 * Display a code editor with highlight that does not contain the feature of local storage.
 * Language select feature is provided.
 * @param storageKey The id of the blank, in case you need to manipulate answers afterwards.
 */
const CodeReadOnly = ({storageKey} : {storageKey: string}) => {
    const [language, setLanguage] = useState("c");
    return (
        <div>
            <CodeEditor
                id={storageKey}
                language={language}
                className="mb-2"
                padding={10}
                style={{
                    height: "200px",
                    fontSize: 12,
                    backgroundColor: "#f5f5f5",
                    fontFamily: 'ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace',
                }}
                readOnly
            />
            <Form.Group>
                <Form.Label>Choose the language:</Form.Label>
                <Form.Select value={language} onChange={(evt) => setLanguage(evt.target.value)}>
                    {listOfLanguages.map(lang => {
                        return (
                            <option value={lang}>
                                {lang}
                            </option>
                        );
                    })}
                </Form.Select>
            </Form.Group>
        </div>
    )
}

export default CodeReadOnly;
