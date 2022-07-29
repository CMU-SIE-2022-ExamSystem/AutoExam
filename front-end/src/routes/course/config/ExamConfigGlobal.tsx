import React, {useCallback, useEffect, useState} from 'react';
import {useConfigStates} from "./ExamConfigStates";
import {Col, Form, InputGroup, Row} from "react-bootstrap";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {DateTime, Namespace, TempusDominus} from "@eonasdan/tempus-dominus";
import {ChangeEvent} from "@eonasdan/tempus-dominus/types/utilities/event-types";

const DateTimePicker = ({pickerId}: { pickerId: string }) => {
    return (
        <InputGroup id={pickerId} data-td-target-input={'nearest'} data-td-target-toggle={'nearest'}>
            <Form.Control type="text" id={pickerId + 'Input'} data-td-target={"#" + pickerId}/>
            <InputGroup.Text data-td-target={"#" + pickerId} data-td-toggle={'datetimepicker'}><i
                className="bi bi-calendar"/></InputGroup.Text>
        </InputGroup>
    )
}

const toDateTime = (dateString: string) => {
    return DateTime.convert(new Date(dateString));
}

const bindPicker = (pickerElement: HTMLElement, restrictions: any, initValue: DateTime) => new TempusDominus(pickerElement, {
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
        },
    },
    restrictions: restrictions,
    useCurrent: false,
    defaultDate: initValue,
});

const ExamConfigGlobal = () => {
    let {examConfigState, setExamConfigState} = useConfigStates();

    const {globalState} = useGlobalState();

    const [categoryList, setCategoryList] = useState<string[]>([]);
    const [apiReady, setApiReady] = useState<boolean>(false);

    const getCategoryList = useCallback(async () => {
        const categoryListUrl = getBackendApiUrl("/courses/assessments/config/categories");
        const token = globalState.token;
        const categoryListResult = await axios.get(categoryListUrl, {headers: {Authorization: "Bearer " + token}});
        setCategoryList(categoryListResult.data.data.categories);
        setApiReady(true);
    }, [globalState.token])

    useEffect(() => {
        getCategoryList().catch();
    }, [getCategoryList]);

    const updateState = (updateTerm: any) => {
        const newState = Object.assign({}, examConfigState, updateTerm)
        setExamConfigState(newState);
    }

    const updateGeneral = (updateTerm: any) => {
        updateState({general: Object.assign({}, examConfigState?.general, updateTerm)})
    }

    const formCheck = categoryList.map(category => (
        <Form.Check
            inline
            label={category}
            name={"category_name"}
            type="radio"
            key={`category-${category}`}
            id={`category-${category}`}
            checked={examConfigState?.general.category_name === category}
            value={category}
            onChange={(e) => {
                updateGeneral({category_name: e.target.value})
            }}
        />
    ))

    useEffect(() => {
        if (!apiReady) return;
        if (!examConfigState) return;
        let {start_at, end_at, grading_deadline} = examConfigState.general;
        const startAtDateTime = toDateTime(start_at);
        const endAtDateTime = toDateTime(end_at);
        const gradingDeadlineDateTime = toDateTime(grading_deadline)
        const pickers = [{pickerId: 'start_at', value: startAtDateTime, restrictions: {maxDate: endAtDateTime}},
            {pickerId: 'end_at', value: endAtDateTime, restrictions: {minDate: startAtDateTime, maxDate: gradingDeadlineDateTime}},
            {pickerId: 'grading_deadline', value: gradingDeadlineDateTime, restrictions: {minDate: endAtDateTime}}].map(({
                                                                                                      pickerId,
                                                                                                      value,
                                                                                                      restrictions
                                                                                                  }) => {
            const myElement = document.getElementById(pickerId) as HTMLElement;
            const inputElement = document.getElementById(pickerId + 'Input') as HTMLInputElement;
            console.log(value.format({timeStyle: "short", dateStyle: "short"}));
            inputElement.value = value.format({timeStyle: "short", dateStyle: "short"});
            return bindPicker(myElement, restrictions, value);
        })
        pickers[0].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newStart = e.date;
            updateGeneral({start_at: newStart?.toISOString()});
            pickers[1].updateOptions({
                restrictions: {
                    minDate: newStart
                }
            })
        });
        pickers[1].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newEnd = e.date;
            updateGeneral({end_at: newEnd?.toISOString()});
            pickers[0].updateOptions({
                restrictions: {
                    maxDate: newEnd
                }
            });
            pickers[2].updateOptions({
                restrictions: {
                    minDate: newEnd
                }
            })
        });
        pickers[2].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newDDL = e.date;
            updateGeneral({grading_deadline: newDDL?.toISOString()});
            pickers[1].updateOptions({
                restrictions: {
                    maxDate: newDDL
                }
            })
        });
    }, [apiReady]);

    return (
        <Row className="text-start">
            <Col className="mb-3" xs={{span: '12'}}>
                <h1>Exam Global Settings</h1>
            </Col>
            <Col sm={{span: '8', offset: '2'}}>
                <Form>
                    <Form.Group className="mb-3" controlId="name">
                        <Form.Label>Name</Form.Label>
                        <Form.Control type="text" value={examConfigState?.general.name || ""} onChange={(e) => {
                            updateGeneral({name: e.target.value})
                        }}/>
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="category_name">
                        <div>
                            <Form.Label>Category</Form.Label>
                        </div>
                        {formCheck}
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>Start At</Form.Label>
                        <DateTimePicker pickerId="start_at"/>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>End At</Form.Label>
                        <DateTimePicker pickerId="end_at"/>
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>Grading Deadline</Form.Label>
                        <DateTimePicker pickerId="grading_deadline"/>
                    </Form.Group>
                </Form>
            </Col>
        </Row>

    )
}

export default ExamConfigGlobal;