import React, {useCallback, useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {Card, Col, Nav, Row, Tab} from "react-bootstrap";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import moment from "moment";

interface LooseObject {
    [key: string]: any
}

interface examResultsType {
    version: number,
    filename: string;
    created_at: string;
    scores: LooseObject,
    total_score: number,
}

const ResultPage = ({questions, answers} : {questions: questionDataType[], answers: examResultsType}) => {
    const date = moment(answers.created_at).format("MM/DD/YYYY, HH:mm:ss Z");
    const answerMetaData = (
        <Card className="mt-2 text-start">
            <Card.Header>Metadata</Card.Header>
            <Card.Body>
                <Card.Text>Version: {answers.version}</Card.Text>
                <Card.Text>Created at: {date}</Card.Text>
                <Card.Text>Total Score: {answers.total_score}</Card.Text>
            </Card.Body>
        </Card>
    );
    return (
        <div>
            {answerMetaData}
        </div>
    );
}

const ExamResultQuestions = () => {

    let params = useParams();
    const {globalState, updateGlobalState} = useGlobalState();
    const navigate = useNavigate();

    const courseName = params.course_name;
    const examId = params.exam_id;

    const [questionList, setQuestionList] = useState<questionDataType[]>([]);
    const [examResults, setExamResults] = useState<examResultsType[]>([]);

    const getQuestionList = useCallback(() => {
        const questionUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/question`);
        const token = globalState.token;
        return axios.get(questionUrl, {headers: {Authorization: "Bearer " + token}});
    }, [globalState.token]);
    const getAnswer = useCallback(() => {
        const answerUrl = getBackendApiUrl(`/courses/${courseName}/assessments/${examId}/submissions`);
        const token = globalState.token;
        return axios.get(answerUrl, {headers: {Authorization: "Bearer " + token}});
    }, []);

    useEffect(() => {
        getQuestionList()
            .then(result => {
                const questionList: questionDataType[] = result.data.data;
                setQuestionList(questionList);
            })
            .then(_ => {
                getAnswer()
                    .then(result => {
                        const answer: examResultsType[] = result.data.data;
                        setExamResults(answer);
                    });
            })
            .catch(badExam => {
                console.error(badExam);
            });
    }, [getQuestionList, getAnswer]);

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
            <Nav.Item>
                <Nav.Link eventKey={"attempt_" + index} href="#">
                    Attempt {examResult.version} ({date})
                </Nav.Link>
            </Nav.Item>
        )
    });

    const tabPanes = examResults.map((examResult, index) => {
        return (
            <Tab.Pane eventKey={"attempt_" + index}>
                <ResultPage questions={questionList} answers={examResult} />
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
