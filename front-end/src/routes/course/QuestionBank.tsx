import React from 'react';
import {Col, Row} from 'react-bootstrap';
import {useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";

const QuestionBank = () => {
    let params = useParams();
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name}/>
            </Row>
            <main>
                <Row>
                    <Col xs={{span: "8", offset: "2"}}>
                        <h1>Question Bank</h1>
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default QuestionBank;
