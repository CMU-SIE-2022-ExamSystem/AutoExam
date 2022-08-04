import React, {useCallback, useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Card, Col, Nav, Row, Tab, Table} from "react-bootstrap";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import moment from "moment";

interface LooseObject extends Object {
    [key: string]: any
}

interface examResultsType {
    version: number,
    filename: string;
    created_at: string;
    scores: LooseObject,
    max_score: LooseObject,
    total_score: number,
}

const ResultPage = ({answers} : {answers: examResultsType}) => {
    const date = moment(answers.created_at).format("MM/DD/YYYY, HH:mm:ss Z");

    const scoreKeys = Object.keys(answers.scores);

    let totalMaxScore = 0;

    const questionTr = scoreKeys.map((key) => {
        let myScore = answers.scores[key];
        let maxScore = answers.max_score[key];
        totalMaxScore += maxScore;
        return <tr key={key}><td>{key}</td><td>{myScore} / {maxScore}</td></tr>
    });

    const studentScoreTbody = (
        <tbody>
            {questionTr}
        </tbody>
    )

    const studentSubmission = (
        <Card className="mt-2 text-start">
            <Card.Header>Score details</Card.Header>
            <Card.Body>
                <Row>
                    <Col md={{span: '8', offset: '2'}}>
                        <Table striped bordered>
                            <thead>
                            <tr>
                                <th scope="col">Question</th>
                                <th scope="col">Score</th>
                            </tr>
                            </thead>
                            {studentScoreTbody}
                        </Table>
                    </Col>
                </Row>
            </Card.Body>
        </Card>
    )

    const answerMetaData = (
        <Card className="mt-2 text-start">
            <Card.Header>Metadata</Card.Header>
            <Card.Body>
                <Card.Text>Version: {answers.version}</Card.Text>
                <Card.Text>Submitted at: {date}</Card.Text>
                <Card.Text>Total Score: {answers.total_score} / {totalMaxScore}</Card.Text>
            </Card.Body>
        </Card>
    );

    return (
        <div>
            {answerMetaData}
            {studentSubmission}
        </div>
    );
}

const ExamResultQuestions = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    const [examResults, setExamResults] = useState<examResultsType[]>([]);

    const getSubmission = useCallback(() => {
        const submissionUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/submissions`);
        const token = globalState.token;
        return axios.get(submissionUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getSubmission()
            .then(result => {
                const examResults: examResultsType[] = result.data.data;
                setExamResults(examResults);
            })
            .catch(badExam => {
                console.error(badExam);
            });
    }, [getSubmission]);

    const activeResult = examResults ? examResults.at(0) : undefined;

    let layout = (
        <div>
            <p>We cannot find any of your submissions.</p>
            <p>If you have just submitted your work, it may take time to finish handling and display here.</p>
            <p>Please try again later or contact your instructors.</p>
        </div>
    );

    const navItemList = examResults.map((examResult, index) => {
        const date = moment(examResult.created_at).format("M/D/YY");
        return (
            <Nav.Item key={examResult.filename + "nav"}>
                <Nav.Link eventKey={"attempt_" + index} href="#">
                    Attempt {examResult.version} ({date})
                </Nav.Link>
            </Nav.Item>
        )
    });

    const tabPanes = examResults.map((examResult, index) => {
        return (
            <Tab.Pane eventKey={"attempt_" + index} key={examResult.filename + "pane"}>
                <ResultPage answers={examResult} />
            </Tab.Pane>
        )
    });

    if (activeResult) {
        layout = (
            <Tab.Container defaultActiveKey={"attempt_0"}>
                <Row>
                    <Col xs={{span: "2"}}>
                        <Nav variant="pills" className="flex-column p-3">
                            {navItemList}
                        </Nav>
                    </Col>
                    <Col xs={{span: "10"}}>
                        <Tab.Content>
                            {tabPanes}
                        </Tab.Content>
                    </Col>
                </Row>
            </Tab.Container>
        );
    }

    return (
        <div>
            {layout}
        </div>
    )
}

export default ExamResultQuestions;
