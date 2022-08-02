import usePersistState from "../../utils/usePersistState";
import {Form} from "react-bootstrap";
import React from "react";

const CodeInput = ({storageKey, displayIdx} : {storageKey: string, displayIdx: number}) => {
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

export default CodeInput;
