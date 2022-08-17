import React, {useCallback, useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Alert, Col, Row, Table} from "react-bootstrap";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";

interface statsProperties {
    best: boolean;
    number: number;
    highest: number;
    lowest: number;
    mean: number;
}

const ExamResultStats = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const [stats, setStats] = useState<statsProperties | null>(null);

    const courseName = params.course_name;
    const examId = params.exam_id;

    const getStats = useCallback(() => {
        const statsUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/statistic`);
        const token = globalState.token;
        return axios.get(statsUrl, {headers: {Authorization: "Bearer " + token}});
    }, [globalState.token, courseName, examId]);

    useEffect(() => {
        getStats()
            .then(result => {
                let data = result.data.data;
                if (data) setStats(data);
            })
    }, []);


    let tbody = ((s) => {
        if (s === null) return (<tbody />);
        return (<tbody>
            <tr><td># of students</td><td>{s.number}</td></tr>
            <tr><td>Highest score</td><td>{s.highest}</td></tr>
            <tr><td>Lowest score</td><td>{s.lowest}</td></tr>
            <tr><td>Mean score</td><td>{s.mean}</td></tr>
        </tbody>
        )
    })(stats);

    let tableElement = (
        <Table className="table text-start" striped bordered>
            <thead>
            <tr>
                <th scope="col">Criteria</th>
                <th scope="col">Score</th>
            </tr>
            </thead>
            {tbody}
        </Table>
    );

    if (stats === null || stats.number === 0) {
        tableElement = (
            <Alert variant="warning">
                <i className="bi-pencil-square"/><span> The statistics are not published right now.</span>
            </Alert>
        )
    }

    return (
        <div>
            <h1>Class Statistics</h1>
            <Row>
                <Col sm={{span: '6', offset: '3'}}>
                    {tableElement}
                </Col>
            </Row>

        </div>
    )
}

export default ExamResultStats;
