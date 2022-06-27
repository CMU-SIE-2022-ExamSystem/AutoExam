import React from 'react';
import { Button, Card } from 'react-bootstrap';
import { useParams } from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import Question from "../../components/Question";
import CountdownTimer from "../../components/CountdownTimer";

const QuestionList = () => {

}

function ExamQuestions() {
    let params = useParams();
    const questionList = QuestionList();

    const targetTime = new Date(Date.now() + 1000 * 100).toString();
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1 className="my-3">{params.exam_id}</h1>
                    <h2 className="text-start my-4"><strong>Instructions</strong></h2>
                    <p className="text-start">Some detailed instructions.</p>
                    <br/>
                    <Question></Question>
                    <br/>
                    {/* {questionList} */}
                    <CountdownTimer targetTime={targetTime} callback={() => {}} />
                    <div><Button variant="primary" className="my-4">Submit</Button></div>
                </>
            </AppLayout>
        </div>
    );
}

export default ExamQuestions;
