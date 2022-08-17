import {Form} from "react-bootstrap";
import React from "react";

/**
 * Display a blank that does not contain the feature of local storage.
 * @param storageKey The id of the blank, in case you need to manipulate answers afterwards.
 */
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
