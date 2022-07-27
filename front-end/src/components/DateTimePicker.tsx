import React, {useEffect} from 'react';
import {TempusDominus} from "@eonasdan/tempus-dominus";
import { Form, InputGroup } from 'react-bootstrap';

export const DateTimePicker = ({setUpdate, pickerId} : {setUpdate : (timeDate: string) => void, pickerId: string}) => {

    useEffect(() =>{
        const picker = new TempusDominus(document.getElementById(pickerId) as HTMLElement, {
            display: {
                sideBySide: true,
                icons: {
                    time: "bi bi-clock-fill",
                    date: "bi bi-calendar-fill",
                    up: "bi bi-arrow-up",
                    down: "bi bi-arrow-down",
                    next: "bi bi-chevron-right",
                    previous: "bi bi-chevron-left",
                    today: "bi bi-calendar-check-fill",
                    clear: "bi bi-trash-fill",
                    close: "bi bi-x",
                    type: "icons"
                }
            }
        });
    }, []);

    return (
        <InputGroup id={pickerId} data-td-target-input={'nearest'} data-td-target-toggle={'nearest'}>
            <Form.Control type="text" id={pickerId + 'Input'} data-td-target={"#" + pickerId}/>
            <InputGroup.Text data-td-target={"#" + pickerId} data-td-toggle={'datetimepicker'}><i className="bi bi-calendar" /></InputGroup.Text>
        </InputGroup>
        )
}