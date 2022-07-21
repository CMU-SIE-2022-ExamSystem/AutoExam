import React from 'react';
import {Editor} from "@tinymce/tinymce-react";
import {getFrontendUrl} from "../utils/url";

const HTMLEditor = ({editorRef} : {editorRef: any}) => {
    return (
        <Editor
            tinymceScriptSrc={getFrontendUrl('') + '/tinymce/tinymce.min.js'}
            onInit={(evt, editor) => editorRef.current = editor}
            init={{
                height: 500,
                menubar: false,
                plugins: [
                    'advlist','autolink',
                    'lists','link','image','charmap','preview','anchor','searchreplace','visualblocks',
                    'fullscreen','insertdatetime','media','table','help','wordcount'
                ],
                toolbar: 'undo redo | casechange blocks | bold italic backcolor | ' +
                    'alignleft aligncenter alignright alignjustify | ' +
                    'bullist numlist checklist outdent indent | removeformat | a11ycheck code table help',
                content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }'
            }}
        />
    )
}

export default HTMLEditor;
