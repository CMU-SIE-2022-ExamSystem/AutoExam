import React, {useCallback, useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import {useGlobalState} from "../../../components/GlobalStateProvider";
import {ExamConfigSettingsType, useConfigStates} from "./ExamConfigStates";
import {getBackendApiUrl} from "../../../utils/url";
import axios from "axios";
import {Button, Card, Table} from "react-bootstrap";

interface tagProps {
    id: string;
    name: string;
    course: string;
}

const settingToQuestion = (setting: ExamConfigSettingsType, qIndex: number, tags: tagProps[]) => {
    const myTag = tags.find(tag => tag.id === setting.tag);
    const idLength = setting.id.length;
    const appointed = idLength > 0 ? ("Chosen from " + idLength + "questions") : "Random pick according to tag";
    const subquestionTitles = Array.from({length: setting.sub_question_number}, (value, index) => "sub" + (index + 1).toString());
    const subquestionScores = (
        <Table bordered>
            <thead>
                <tr><th>Questions</th>{subquestionTitles.map(title => <th scope="col" key={title}>{title}</th>)}</tr>
            </thead>
            <tbody>
                <tr><td>Scores</td>{setting.scores.map(score => <td>{score}</td>)}</tr>
            </tbody>
        </Table>
    )
    return (
        <Card className="my-3 text-start" key={qIndex}>
            <Card.Header>
                {qIndex + 1}. {myTag ? myTag.name : setting.title}
            </Card.Header>
            <Card.Body>
                <Card.Text>Question title: {setting.title}</Card.Text>
                <Card.Text>Score: {setting.max_score}</Card.Text>
                <Card.Text># of sub questions: {setting.sub_question_number}</Card.Text>
                {subquestionScores}
                <Card.Text>Generation source: {appointed}</Card.Text>
            </Card.Body>
            <Card.Footer>
                <Button variant="warning" size="sm" className="me-1"><i className="bi bi-pencil-square me-1"/>Edit</Button>
                <Button variant="danger" size="sm"><i className="bi bi-trash me-1"/>Delete</Button>
            </Card.Footer>
        </Card>
    );
};

const ExamConfigQuestions = () => {
    let params = useParams();
    const courseName = params.course_name;
    const examId = params.exam_id;
    const {globalState} = useGlobalState();
    let {examConfigState, setExamConfigState}  = useConfigStates();

    const [tags, setTags] = useState<tagProps[]>([]);
    const getTags = useCallback(() => {
        const url = getBackendApiUrl("/courses/" + courseName + "/tags");
        const token = globalState.token;
        return axios.get(url, {headers: {Authorization: "Bearer " + token}});
    }, []);


    useEffect(() => {
        getTags()
            .then(response => {
                const data : tagProps[] = response.data.data;
                setTags(data);
            })
    }, [])

    const settingsToQuestion = examConfigState?.settings?.map((setting, index) => settingToQuestion(setting, index, tags));
    return (
        <div>
            <div className="text-end mb-2">
                <Button variant="success" className="me-1"><i className="bi bi-plus-square me-1"/>Add Question</Button>
                <Button variant="warning"><i className="bi bi-list me-1"/>Change Order</Button>
            </div>
            {settingsToQuestion}
        </div>
    )
}

export default ExamConfigQuestions;