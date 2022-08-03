import React, {useCallback, useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Col, Nav, Row, Tab} from "react-bootstrap";
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

    const getStats = useCallback(() => {
        const statsUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/statistic`);
        const token = globalState.token;
        return axios.get(statsUrl, {headers: {Authorization: "Bearer " + token}});
    }, [globalState.token]);

    useEffect(() => {
        getStats()
            .then(result => {
                let data = result.data.data;
                if (data) setStats(data);
            })
    }, []);

    const courseName = params.course_name;
    const examId = params.exam_id;

    let tbody = ((s) => {
        if (s === null) return (<tbody />);
        return (<tbody>
            <tr><td># of students</td><td>{s.number}</td></tr>
            <tr><td>Highest score</td><td>{s.highest}</td></tr>
            <tr><td>Lowest score</td><td>{s.lowest}</td></tr>
            <tr><td>Mean score</td><td>{s.mean}</td></tr>
        </tbody>
        )
    })(stats)

    return (
        <div>
            <h1>Class Statistics</h1>
            <Row>
                <Col sm={{span: '6', offset: '3'}}>
                    <table className="table text-start">
                        <thead>
                        <tr>
                            <th scope="col">Criteria</th>
                            <th scope="col">Score</th>
                        </tr>
                        </thead>
                        {tbody}
                    </table>
                </Col>
            </Row>

        </div>
    )
}

export default ExamResultStats;
