import React, {useCallback, useEffect, useState} from 'react';
import {Alert, Button, Form, Modal} from "react-bootstrap";
import {Link, useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";

const BaseCourseRelationshipManageModal = ({show, toClose} : {show: boolean, toClose: () => void}) => {

    const params = useParams();
    const courseName = params.course_name;
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const [firstTime, setFirstTime] = useState(false);
    const [listOfBaseCourses, setListOfBaseCourses] = useState<string[]>([]);
    const [myCourse, setMyCourse] = useState<string>("");
    const [showErrorMessage, setShowErrorMessage] = useState(false);

    const getBaseCourses = useCallback(() => {
        const baseCourseUrl = getBackendApiUrl("/basecourses/list");
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
        } else {
            return axios.put(baseCourseUrl, {},{headers: {Authorization: "Bearer " + token}});
        }
    }, [firstTime]);

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

    let baseCourseOptions = listOfBaseCourses.map(baseCourse => (
        <option key={"baseCourse" + baseCourse} value={baseCourse}>{baseCourse}</option>
    ))

    if (listOfBaseCourses.length === 0) {
        baseCourseOptions = [(<option value="undefined" key="null">{"==No available base courses=="}</option>)]
    } else {
        baseCourseOptions.unshift((
            <option value="undefined" key="null">Select one option</option>
        ))
    }

    const [showTips, setShowTips] = useState(false);
    const baseCourseTips = (
        <Alert variant="info">
            Base course is a category that the course belongs to. Courses with same base course share question banks.
            Usually we set the course number as base course.
        </Alert>
    )

    const updateRelationship = () => {
        const selectResult = (document.getElementById('base-course-select') as HTMLInputElement).value;
        if (selectResult !== 'undefined' && listOfBaseCourses.find(item => item === selectResult)) {
            updateBaseCourse(selectResult)
                .then(response => {
                    setFirstTime(false);
                    setMyCourse(selectResult);
                    updateGlobalState({alert: {variant: "success", content: "Successfully set the base course to " + selectResult + ".", show: true}});
                    toClose();
                })
                .catch(badRequest => {
                    updateGlobalState({alert: {variant: "danger", content: "Failed on setting the base course.", show: true}});
                });
        } else {
            setShowErrorMessage(true);
        }
    }

    const deleteHandler = () => {
        deleteBaseCourse()
            .then(response => {
                setFirstTime(true);
                updateGlobalState({alert: {variant: "success", content: "Base course relationship removed", show: true}});
                navigate('/dashboard');
            })
            .catch(badRequest => {
                updateGlobalState({alert: {variant: "danger", content: "Failed on removing the base course relationship.", show: true}});
            });
    }

    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>Set base course</Modal.Title>
                <i className="bi bi-info-circle pointer-cursor" onClick={() => setShowTips(!showTips)}/>
            </Modal.Header>
            <Modal.Body>
                {showTips && baseCourseTips}
                <Form>
                    <p>Choose the base course from the following options, or <Link to={"/baseCourse"}>click here to add a new base course</Link>.</p>
                    <Form.Select id="base-course-select" value={myCourse} onChange={(e) => {setShowErrorMessage(false); setMyCourse(e.target.value)}}>
                        {baseCourseOptions}
                    </Form.Select>
                    {showErrorMessage && (<small className="text-danger">Please choose one option to proceed.</small>)}
                </Form>
                {!firstTime && (
                    <>
                        <br />
                        <span>You may also delete the base course relationship: </span>
                        <Button variant={"danger"} onClick={deleteHandler}>
                            <i className="bi bi-trash" />
                        </Button>
                    </>
                )}
            </Modal.Body>
            <Modal.Footer>
                {!firstTime && (
                    <Button variant="secondary" onClick={toClose}>Back</Button>
                )}
                <Button variant="primary" onClick={updateRelationship}>Submit</Button>
            </Modal.Footer>
        </Modal>
    );
}

export default BaseCourseRelationshipManageModal;
