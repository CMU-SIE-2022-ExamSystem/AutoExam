import {Form} from "react-bootstrap";
import React from "react";

const BlankReadOnly = ({storageKey, value} : {storageKey: string, value: string}) => {
    return (
        <div>
            <Form.Control type="text"
                          id={storageKey}
                          className="w-50 mb-2"
                          value={value}
                          readOnly
            />
        </div>
    )
}

export default BlankReadOnly;
