import React, {useCallback, useEffect, useState} from 'react';
import {Row, Col, Button} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import {getBackendApiUrl} from "../../utils/url";
import {default as axios} from "axios";
import {useGlobalState} from "../../components/GlobalStateProvider";
import moment from 'moment';

interface assessmentProps {
    name: string;
    display_name: string;
    start_at: string;
    due_at: string;
    end_at: string;
    category_name: string;
    grading_deadline?: string;
}

interface ICourseInfo {
    name: string;
    display_name: string;
    auth_level: string;
}

interface extAssessmentProps extends assessmentProps {
    permission: boolean
}

const AssessmentRow = ({name, display_name, start_at, due_at, permission}: extAssessmentProps) => {
    let startTime = moment(start_at).format("MMMM Do YYYY, h:mm:ss a");
    let dueTime = moment(due_at).format("MMMM Do YYYY, h:mm:ss a");

    let actionList = [];
    if (permission) {
        actionList.push(<Link to={"exams/" + name} className="btn btn-success me-2">Edit Exam</Link>)
        actionList.push(<Link to={"exams/" + name} className="btn btn-primary">Proctor Exam</Link>)
    } else {
        actionList.push(<Link to={"exams/" + name} className="btn btn-primary">Take Exam</Link>);
    }

    return (
        <tr>
            <th scope="row">{display_name}</th>
            <td>{startTime}</td>
            <td>{dueTime}</td>
            <td>{actionList}</td>
        </tr>
    )
}

const Table = (listOfAssessments: assessmentProps[], courseInfo?: ICourseInfo) => {

    const permission = courseInfo?.auth_level === 'instructor';

    const tableBody = listOfAssessments.map(assessment => <AssessmentRow key={assessment.name} permission={permission} {...assessment}/>)

    return (
        <table className="table text-start">
            <thead>
            <tr>
                <th scope="col">Assessment</th>
                <th scope="col">Start At</th>
                <th scope="col">Due At</th>
                <th scope="col">Actions</th>
            </tr>
            </thead>
            <tbody>
            {tableBody}
            </tbody>
        </table>
    )
}

function Assessments() {
    const params = useParams();
    const {globalState} = useGlobalState();
    const [examList, setExamList] = useState<assessmentProps[]>([]);
    const [courseInfo, setCourseInfo] = useState<ICourseInfo>();

    const getCourseInfo = useCallback(async () => {
        const infoUrl = getBackendApiUrl("/courses/" + params.course_name + "/info");
        const assessmentUrl = getBackendApiUrl("/courses/" + params.course_name + "/assessments");
        const token = globalState.token;
        const infoResult = await axios.get(infoUrl, {headers: {Authorization: "Bearer " + token}});
        setCourseInfo(infoResult.data.data);
        const assessmentResult = await axios.get(assessmentUrl, {headers: {Authorization: "Bearer " + token}});
        setExamList(assessmentResult.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getCourseInfo().catch();
    }, [getCourseInfo])


    const assessmentTable = Table(examList, courseInfo);
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={courseInfo?.display_name || params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <main>
                {courseInfo?.auth_level === "instructor" &&
                    <div className="text-end pe-5">
                        <Link to={"examConfig/new/base"}><Button variant="info" className="me-3 text-white">New Exam</Button></Link>
                        <Link to={"questionBank"}><Button variant="primary">Question Bank</Button></Link>
                    </div>
                }
                <Row>
                    <Col xs={{span: "10", offset: "1"}}>
                        <h1>Assessments</h1>
                        {assessmentTable}
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default Assessments;
