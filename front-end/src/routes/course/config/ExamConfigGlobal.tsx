import React, {useCallback, useEffect, useState} from 'react';
import {useConfigStates} from "./ExamConfigStates";
import {Col, Form, Row} from "react-bootstrap";
import {DateTimePicker} from "../../../components/DateTimePicker";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../../components/GlobalStateProvider";

const ExamConfigGlobal = () => {
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
                        <DateTimePicker setUpdate={newStart => {updateGeneral({start_at: newStart});}} pickerId="start_at" />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>End At</Form.Label>
                        <DateTimePicker setUpdate={newEnd => {updateGeneral({end_at: newEnd});}} pickerId="end_at" />
                    </Form.Group>
                    <Form.Group className="mb-3">
                        <Form.Label>Grading Deadline</Form.Label>
                        <DateTimePicker setUpdate={newGradingDeadline => {updateGeneral({grading_deadline: newGradingDeadline});}} pickerId="grading_deadline"/>
                    </Form.Group>
                </Form>
            </Col>
        </Row>

    )
}

export default ExamConfigGlobal;