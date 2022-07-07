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


const AssessmentRow = ({name, display_name, start_at, due_at}: assessmentProps) => {
    let startTime = moment(start_at).format("MMMM Do YYYY, h:mm:ss a");
    let dueTime = moment(due_at).format("MMMM Do YYYY, h:mm:ss a");
    return (
        <tr>
            <th scope="row">{display_name}</th>
            <td>{startTime}</td>
            <td>{dueTime}</td>
            <td><Link to={"exams/" + name} className="btn btn-primary">Take Exam</Link></td>
        </tr>
    )
}

const Table = (listOfAssessments: assessmentProps[]) => {

    const tableBody = listOfAssessments.map(assessment => <AssessmentRow key={assessment.name} {...assessment}/>)

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

    const getExams = useCallback(async () => {
        const url = getBackendApiUrl("/courses/" + params.course_name + "/assessments");
        const token = globalState.token;
        const result = await axios.get(url, {headers: {Authorization: "Bearer " + token}});
        console.log(result);
        setExamList(result.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getExams().catch();
    }, [getExams])


    const assessmentTable = Table(examList);
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <main>
                <Row className="text-end pe-5">
                    <Link to={"questionBank"}><Button variant="primary">Question Bank</Button></Link>
                </Row>
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
