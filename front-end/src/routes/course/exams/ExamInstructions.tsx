import React, {useCallback, useEffect, useState} from 'react';
import {Alert, Button, Col, Row} from 'react-bootstrap';
import {Link, useNavigate, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import moment from "moment";

function ExamInstructions() {
    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();
    const courseName = params.course_name, examId = params.exam_id;

    const [description, setDescription] = useState<string>("");

    const getTestGeneralInfo = useCallback(() => {
        const examUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}`);
        const token = globalState.token;
        return axios.get(examUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getTestGeneralInfo()
            .then(result => {
                const testInfo: any = result.data.data;
                let nowMoment = moment();
                if (moment(testInfo.start_at).diff(nowMoment) > 0 || nowMoment.diff(moment(testInfo.end_at)) > 0) {
                    updateGlobalState({alert: {show: true, content: `'${examId}' is not ready.`, variant: 'danger'}});
                    navigate('/courses/' + courseName);
                    return;
                }
                if (!testInfo.can_submit) {
                    updateGlobalState({alert: {show: true, content: `You have used all attempts of '${examId}'.`, variant: 'danger'}});
                    navigate('/courses/' + courseName);
                    return;
                }
                setDescription(testInfo.description);
            })
    }, [getTestGeneralInfo]);

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name}/>
            </Row>
            <main>
                <Row>
                    <Col xs={{span: "8", offset: "2"}}>
                        <div>
                            <h1 className="my-3">{params.exam_id}</h1>
                            <h2 className="text-start my-4"><strong>Instructions</strong></h2>
                            <div dangerouslySetInnerHTML={{__html: description}}/>
                            <Alert key="primary" variant="primary" className="text-start my-4">Please turn on your camera to
                                start the exam.</Alert>
                            <Link to="questions">
                                <Button type="button" className="btn btn-lg btn-primary w-50">Start</Button>
                            </Link>
                        </div>
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default ExamInstructions;
