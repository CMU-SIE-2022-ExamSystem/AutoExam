import React from 'react';
import {} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

type AssessmentProps = {
    name: string
}
const AssessmentRow = ({name}: AssessmentProps) => {
    return (
        <tr>
            <th scope="row">{name}</th>
            <td></td>
            <td></td>
            <td><Link to={"exams/" + name} className="btn btn-primary">Take Exam</Link></td>
        </tr>
    )
}

const Table = () => {
    const listOfAssessments = [{
        exam_id: 1,
        name: 'Exam 1'
    }, {
        exam_id: 2,
        name: 'Exam 2'
    }, {
        exam_id: 3,
        name: 'Final Exam'
    }];
    const tableBody = listOfAssessments.map(assessment => <AssessmentRow key={assessment.exam_id} {...assessment}/>)
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
    const assessmentTable = Table();
    return (
        <AppLayout>
            <TopNavbar brand={params.course_name}/>
            <main>
                <h1>Assessment</h1>
                {assessmentTable}
            </main>
        </AppLayout>
    );
}

export default Assessments;
