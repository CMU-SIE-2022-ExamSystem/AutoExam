import React, {useEffect} from 'react';
import {Namespace, TempusDominus} from "@eonasdan/tempus-dominus";
import { Form, InputGroup } from 'react-bootstrap';
import {ChangeEvent} from "@eonasdan/tempus-dominus/types/utilities/event-types";

export const DateTimePicker = ({setUpdate, pickerId} : {setUpdate : (timeDate: string) => void, pickerId: string}) => {

    useEffect(() =>{
        const myElement = document.getElementById(pickerId) as HTMLElement;
        const picker = new TempusDominus(myElement, {
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
        myElement.addEventListener(Namespace.events.change, (e: Event) => {
            console.log(e);
            setUpdate(new Date((e as CustomEvent).detail.date!).toISOString());
        })
    }, []);

    return (
        <InputGroup id={pickerId} data-td-target-input={'nearest'} data-td-target-toggle={'nearest'}>
            <Form.Control type="text" id={pickerId + 'Input'} data-td-target={"#" + pickerId}/>
            <InputGroup.Text data-td-target={"#" + pickerId} data-td-toggle={'datetimepicker'}><i className="bi bi-calendar" /></InputGroup.Text>
        </InputGroup>
        )
}