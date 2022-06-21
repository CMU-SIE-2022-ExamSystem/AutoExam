import React from 'react';
import { Button } from 'react-bootstrap';
import { Link, useParams } from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

const Question = () => {

}

const QuestionList = () => {

}

function ExamQuestions() {
    let params = useParams();
    const questionList = QuestionList();
    return (
        <div>
            <TopNavbar brand={null}/>
            <AppLayout>
                <>
                    <h1 className="my-3">{params.exam_id}</h1>
                    <h2 className="text-start my-4"><strong>Instructions</strong></h2>
                    <p className="text-start">Some detailed instructions.</p>
                    {questionList}
                    <div><Button variant="primary" className="my-4">Submit</Button></div>
                </>
            </AppLayout>
        </div>
    );
}

export default ExamQuestions;
