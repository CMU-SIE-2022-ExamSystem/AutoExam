import React from 'react';
import {Editor} from "@tinymce/tinymce-react";
import {getFrontendUrl} from "../utils/url";

const HTMLEditor = ({init, update} : {init: string, update: any}) => {
    return (
        <Editor
            tinymceScriptSrc={getFrontendUrl('') + '/tinymce/tinymce.min.js'}
            onChange={(evt, editor) => update(editor.getContent())}
            initialValue={init}
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
                    'bullist numlist outdent indent | removeformat | image table help',
                content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }'
            }}
        />
    )
}

export default HTMLEditor;
