import usePersistState from "../../utils/usePersistState";
import {Form} from "react-bootstrap";
import React from "react";

/**
 * Display a blank that does not contain the feature of local storage.
 * @param storageKey The id of the blank, in case you need to manipulate answers afterwards.
 */
const BlankInput = ({storageKey, displayIdx} : {storageKey: string, displayIdx: number}) => {
    const {value, setValue} = usePersistState("", storageKey);
    return (
        <div>
            {/*<Form.Label>({displayIdx}). </Form.Label>*/}
            <Form.Control type="text"
                          id={storageKey}
                          className="w-50 mb-2"
                          value={value}
                          onChange={(event) => {
                              const newValue = event.target.value;
                              setValue(newValue);
                          }}
            />
        </div>
    )
}

export default BlankInput;
