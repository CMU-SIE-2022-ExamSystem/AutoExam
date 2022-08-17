import React from 'react';
import {} from 'react-bootstrap';
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
            <td></td>
            <td></td>
        </tr>
    )
}

const Table = () => {
    const listOfAssessments = [{
        name: 'Exam 1'
    },{
        name: 'Exam 2'
    },{
        name: 'Final Exam'
    }];
    const tableBody = listOfAssessments.map(assessment => <AssessmentRow {...assessment}/>)
    return (
        <table className="table text-start">
            <thead>
                <tr>
                    <th scope="col">Assessment</th>
                    <th scope="col">Start At</th>
                    <th scope="col">Due At</th>
                    <th scope="col">End At</th>
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
    const assessmentTable = Table();
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1>Assessment</h1>
                    {assessmentTable}
                </>
            </AppLayout>
        </div>
    );
}

export default Assessments;
