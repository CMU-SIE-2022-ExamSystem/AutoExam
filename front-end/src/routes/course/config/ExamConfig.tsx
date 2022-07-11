import React, {useEffect} from 'react';
import {Button, Col, Row} from 'react-bootstrap';
import {Link, Outlet, useParams} from "react-router-dom";
import TopNavbar from "../../../components/TopNavbar";
import AppLayout from "../../../components/AppLayout";

function ExamConfig({isNew}: {isNew: boolean}) {
    let params = useParams();

    const getSavedConfig = async () =>  {

    }

    useEffect(() => {
        if (!isNew) {
            getSavedConfig().catch();
        }
    }, [getSavedConfig]);

    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <main>
                <Row>
                    <Col xs={{span: "3"}}>
                        <div className="d-flex flex-column flex-shrink-0 p-3">
                            <div>Exam Config</div>
                            <hr />
                            <ul className="nav nav-pills flex-column mb-auto">
                                <li className="nav-item">
                                    Global Settings
                                </li>
                                <li className="nav-item">
                                    Exam Questions
                                </li>
                            </ul>
                        </div>
                    </Col>
                    <Col xs={{span: "9"}}>
                        <div>
                            <Outlet />
                        </div>
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default ExamConfig;
