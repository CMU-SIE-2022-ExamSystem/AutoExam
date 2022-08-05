import React, {useCallback, useEffect, useState} from 'react';
import {Alert, Modal, OverlayTrigger, Tooltip} from "react-bootstrap";
import {useParams} from "react-router-dom";
import {useGlobalState} from "./GlobalStateProvider";
import {getBackendApiUrl} from "../utils/url";
import axios from "axios";

const BaseCourseModal = ({show, toClose} : {show: boolean, toClose: () => void}) => {

    const params = useParams();
    const courseName = params.course_name;
    const {globalState} = useGlobalState();

    const [firstTime, setFirstTime] = useState(false);
    const [listOfBaseCourses, setListOfBaseCourses] = useState<string[]>([]);
    const [myCourse, setMyCourse] = useState<string>("");

    const getBaseCourses = useCallback(() => {
        const baseCourseUrl = getBackendApiUrl("/basecourse/list");
        const token = globalState.token;
        return axios.get(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    const getMyBaseCourse = useCallback(() => {
        const baseCourseUrl = getBackendApiUrl("/courses/" + courseName + "/base");
        const token = globalState.token;
        return axios.get(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    const updateBaseCourse = useCallback((newBase: string) => {
        const baseCourseUrl = getBackendApiUrl("/courses/" + courseName + "/base/" + newBase);
        const token = globalState.token;
        if (firstTime) {
            return axios.post(baseCourseUrl, {},{headers: {Authorization: "Bearer " + token}});
        }
        return axios.put(baseCourseUrl, {},{headers: {Authorization: "Bearer " + token}});
    }, []);

    const deleteBaseCourse = useCallback(() => {
        const baseCourseUrl = getBackendApiUrl("/courses/" + courseName + "/base");
        const token = globalState.token;
        return axios.delete(baseCourseUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getBaseCourses()
            .then(response => {
                const listOfCourses : { id: number, name: string}[] = response.data.data;
                setListOfBaseCourses(listOfCourses.map((item) => item.name));
            })
            .then(() => {
                getMyBaseCourse()
                    .then(response => {
                        setMyCourse(response.data.data);
                    })
                    .catch(badRequest => {
                        setFirstTime(true);
                    })
            })
    }, [])

    const [showTips, setShowTips] = useState(false);
    const baseCourseTips = (
        <Alert variant="info">
            Base course is a category that the course belongs to. Courses with same base course share question banks.
            Usually we set the course number as base course.
        </Alert>
    )

    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Set base course</Modal.Title>
                <i className="bi bi-info-circle" onClick={() => setShowTips(!showTips)}/>
            </Modal.Header>
            <Modal.Body>
                {showTips && baseCourseTips}
            </Modal.Body>
        </Modal>
    );
}

export default BaseCourseModal;
