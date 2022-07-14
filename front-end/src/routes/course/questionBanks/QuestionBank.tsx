import React from 'react';
import {Button, Col, Nav, Row, Tab} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";
import questionDataType from "../../../components/questionTemplate/questionDataType";
import Question from "../../../components/Question";

const QuestionsByTag = ({tag}: {tag: string}) => {
    let questionList: questionDataType[];
    questionList = require('../exams/questions_new.json').data;

    let tagQuestionMap = new Map<string, questionDataType[]>();
    for (let i = 0; i < questionList.length; i++) {
        const tag = questionList[i].questionTag;
        if (!tagQuestionMap.has(tag)) {
            tagQuestionMap.set(tag, [questionList[i]]);
        } else {
            (tagQuestionMap.get(tag) as questionDataType[]).push(questionList[i]);
        }
    }

    if (!tagQuestionMap.has(tag)) return (<p>No questions in this tag!</p>);

    const questionsByTag = tagQuestionMap.get(tag);
    return (
        <Row>
            <Col sm={10}>
                {
                    (questionsByTag as questionDataType[]).map((question) => (
                        <Question key={`Q${question.headerId}`} questionData={question} />
                    ))
                }
            </Col>
        </Row>
    );
}

function QuestionBank () {
    let params = useParams();
    
    const tags = ["Integer", "Float", "Cache", "Memory"];

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <Tab.Container id="questionBankSidebar" defaultActiveKey={tags[0]}>
                <Row>
                    <Col sm={2} className="bg-light vh-100 sticky-top">
                        <Nav className="flex-column my-3">
                            {tags.map((tag) => (
                                <Nav.Item>
                                    <Nav.Link eventKey={tag} href="#">
                                        {tag}
                                    </Nav.Link>
                                </Nav.Item>
                            ))}
                        </Nav>
                        <Link to="#"><Button variant="primary">Add Tag</Button></Link>
                    </Col>
                    <Col sm={10}>
                        <Row className="text-end">
                            <Link to="#"><Button variant="primary" className='me-4 my-4'>Add Question</Button></Link>                                
                        </Row>
                        <Tab.Content className="text-start mx-4">
                            {tags.map((tag) => (
                                <Tab.Pane eventKey={tag}>
                                    <QuestionsByTag tag={tag}/>
                                </Tab.Pane>                           
                            ))}
                        </Tab.Content>
                    </Col>
                </Row>
            </Tab.Container>
        </AppLayout>
    );
}

export default QuestionBank;
