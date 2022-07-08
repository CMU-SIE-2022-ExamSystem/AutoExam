import React, {useEffect} from 'react';
import {Button, Col, Row} from 'react-bootstrap';
import {Link, useParams} from "react-router-dom";
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
                <TopNavbar brand={params.course_name}/>
            </Row>
            <main>
                <Row>
                    <Col xs={{span: "3"}}>
                        <div>
                            Left
                        </div>
                    </Col>
                    <Col xs={{span: "9"}}>
                        <div>
                            Right
                        </div>
                    </Col>
                </Row>
            </main>
        </AppLayout>
    );
}

export default ExamConfig;
