import {Form} from "react-bootstrap";
import React from "react";

const BlankReadOnly = ({storageKey} : {storageKey: string}) => {
    return (
        <div>
            <Form.Control type="text"
                          id={storageKey}
                          className="w-50 mb-2"
                          readOnly
            />
        </div>
    )
}

export default BlankReadOnly;
