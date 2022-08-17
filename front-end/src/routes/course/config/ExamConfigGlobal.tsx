import React, {useCallback, useEffect, useState} from 'react';
import {useConfigStates} from "./ExamConfigStates";
import {Col, Form, InputGroup, Row} from "react-bootstrap";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {DateTime, Namespace, TempusDominus} from "@eonasdan/tempus-dominus";
import {ChangeEvent} from "@eonasdan/tempus-dominus/types/utilities/event-types";
import moment from "moment";

const DateTimePicker = ({pickerId}: { pickerId: string }) => {
    return (
        <InputGroup id={pickerId} data-td-target-input={'nearest'} data-td-target-toggle={'nearest'}>
            <Form.Control type="text" id={pickerId + '_input'} data-td-target={"#" + pickerId}/>
            <InputGroup.Text data-td-target={"#" + pickerId} data-td-toggle={'datetimepicker'} className="pointer-cursor"><i
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

const ExamConfigGlobal = ({dataReady} : {dataReady: boolean}) => {
    let {examConfigState, setExamConfigState} = useConfigStates();

    const {globalState} = useGlobalState();

    const [categoryList, setCategoryList] = useState<string[]>([]);

    const getCategoryList = useCallback(async () => {
        const categoryListUrl = getBackendApiUrl("/courses/assessments/config/categories");
        const token = globalState.token;
        const categoryListResult = await axios.get(categoryListUrl, {headers: {Authorization: "Bearer " + token}});
        setCategoryList(categoryListResult.data.data.categories);
    }, [globalState.token])

    useEffect(() => {
        getCategoryList().catch();
    }, [getCategoryList]);

    const updateState = (updateTerm: any) => {
        const newState = Object.assign({}, examConfigState, updateTerm)
        setExamConfigState(newState);
    };

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
            const inputElement = document.getElementById(pickerId + '_input') as HTMLInputElement;
            inputElement.value = moment(value.toISOString()).format("MM/DD/YYYY, h:mm A");
            return bindPicker(myElement, restrictions, value);
        })
        const updateTime = () => {
            // console.log({
            //     start_at: moment(pickers[0].viewDate.toISOString()).toISOString(true),
            //     start_at_input: (document.getElementById('start_at_input') as HTMLInputElement).value,
            //     end_at: moment(pickers[1].viewDate.toISOString()).toISOString(true),
            //     end_at_value: (document.getElementById('end_at_input') as HTMLInputElement).value,
            //     grading_deadline: moment(pickers[2].viewDate.toISOString()).toISOString(true),
            //     grading_deadline_value: (document.getElementById('grading_deadline_input') as HTMLInputElement).value,
            // })
        };
        const p1 = pickers[0].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newStart = e.date;
            if (moment(newStart).diff(moment(startAtDateTime)) === 0) {
                let picker1Date = pickers[1].viewDate;
                pickers[1].updateOptions({
                    restrictions: {
                        minDate: newStart
                    }
                })
                pickers[1].dates.setValue(picker1Date);
            }
            updateTime();
            updateGeneral({
                start_at: moment(pickers[0].viewDate.toISOString()).toISOString(true),
            })
        });
        const p2 = pickers[1].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newEnd = e.date;
            if (moment(newEnd).diff(moment(endAtDateTime)) === 0) {
                let picker0Date = pickers[0].viewDate;
                let picker2Date = pickers[2].viewDate;
                pickers[0].updateOptions({
                    restrictions: {
                        maxDate: newEnd
                    },
                    useCurrent: false,
                    defaultDate: picker0Date
                });
                pickers[2].updateOptions({
                    restrictions: {
                        minDate: newEnd
                    },
                    useCurrent: false,
                    defaultDate: picker2Date
                })
                pickers[0].dates.setValue(picker0Date);
                pickers[2].dates.setValue(picker2Date);
            }
            updateTime();
            updateGeneral({
                end_at: moment(pickers[1].viewDate.toISOString()).toISOString(true),
            })
        });
        const p3 = pickers[2].subscribe(Namespace.events.change, (e: ChangeEvent) => {
            const newDDL = e.date;
            if (moment(newDDL).diff(moment(gradingDeadlineDateTime)) === 0) {
                let picker1Date = pickers[1].viewDate;
                pickers[1].updateOptions({
                    restrictions: {
                        maxDate: newDDL
                    },
                    useCurrent: false,
                    defaultDate: picker1Date
                })
                pickers[1].dates.setValue(picker1Date);
            }
            updateTime();
            updateGeneral({
                grading_deadline: moment(pickers[2].viewDate.toISOString()).toISOString(true)
            })
        });
    }, [dataReady, examConfigState]);

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
                    <Form.Group className="mb-3">
                        <Form.Label>Zoom Link</Form.Label>
                        <Form.Control type="text" value={examConfigState?.general.zoom || ""} onChange={(e) => {updateGeneral({zoom: e.target.value})}} />
                    </Form.Group>
                </Form>
            </Col>
        </Row>

    )
}

export default ExamConfigGlobal;