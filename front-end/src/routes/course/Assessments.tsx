import React, {useCallback, useEffect, useState} from 'react';
import {Row, Col, Button, Modal, Form} from 'react-bootstrap';
import {Link, useNavigate, useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import {getBackendApiUrl} from "../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../components/GlobalStateProvider";
import moment from 'moment';

interface assessmentProps {
    name: string;
    display_name: string;
    start_at: string;
    due_at: string;
    end_at: string;
    category_name: string;
    grading_deadline?: string;
}

interface ICourseInfo {
    name: string;
    display_name: string;
    auth_level: string;
}

interface extAssessmentProps extends assessmentProps {
    permission: boolean
}

const AssessmentRow = ({name, display_name, start_at, due_at, permission}: extAssessmentProps) => {
    let startTime = moment(start_at).format("MMMM Do YYYY, h:mm:ss a");
    let dueTime = moment(due_at).format("MMMM Do YYYY, h:mm:ss a");

    let actionList = [];
    if (permission) {
        actionList.push(<Link to={"examConfig/" + name + "/base"} key="_EditExam" className="btn btn-success m-1">Edit Exam</Link>)
        actionList.push(<Link to={"exams/" + name} key="_ProctorExam" className="btn btn-primary m-1">Proctor Exam</Link>)
    } else {
        actionList.push(<Link to={"exams/" + name} key="_TakeExam" className="btn btn-primary m-1">Take Exam</Link>);
    }

    return (
        <tr className="align-middle">
            <th scope="row" className="text-center">{display_name}</th>
            <td>{startTime}</td>
            <td>{dueTime}</td>
            <td>{actionList}</td>
        </tr>
    )
}

const Table = (listOfAssessments: assessmentProps[], courseInfo?: ICourseInfo) => {

    const permission = courseInfo?.auth_level === 'instructor';

    const tableBody = listOfAssessments.map(assessment => <AssessmentRow key={assessment.name} permission={permission} {...assessment}/>)

    return (
        <table className="table text-start">
            <thead>
            <tr>
                <th scope="col">Assessment</th>
                <th scope="col">Start At</th>
                <th scope="col">Due At</th>
                <th scope="col">Actions</th>
            </tr>
            </thead>
            <tbody>
            {tableBody}
            </tbody>
        </table>
    )
}


const NewExamConfig = ({show, onSubmit, clearMessage, onClose, categoryList, errorMessage} :{ show: boolean, onSubmit: (categoryName: string, name: string) => void, clearMessage: () => void, onClose: () => void, categoryList: string[], errorMessage: string}) => {
    const categoryRadio = categoryList.map((item) => {
        return (
            <Form.Check type='radio'
                        name="newExamCategory"
                        key={item}
                        label={item}
                        inline
                        id={"new-exam-category-" + item}
                        value={item}
                        required
            />
        )
    });
    const validate = () => {
        const examName = (document.getElementById('new-exam-name') as HTMLInputElement).value;

        const categoryNodeList = (document.getElementsByName('newExamCategory'));
        let checked: string = "";
        categoryNodeList.forEach((item) => {
            const inputItem = (item as HTMLInputElement);
            if (inputItem.checked) checked = inputItem.value;
        })

        onSubmit(checked, examName);
    }
    return (
        <Modal show={show}>
            <Modal.Header>
                <Modal.Title>New Exam</Modal.Title>
            </Modal.Header>

            <Modal.Body>
                <p>Please type the new exam name:</p>
                <Form onSubmit={(event) => {event.preventDefault(); validate();}}>
                    <Form.Group className="pb-4">
                        <Form.Control type="text"
                                      className="mb-2"
                                      required
                                      id="new-exam-name"
                                      pattern="[A-za-z][A-za-z0-9_]{0,15}"
                                      onChange={clearMessage}
                        />
                        <small>Use only alphabets, numbers and underscore(_).</small>
                        <div>
                            <small className="text-danger">{errorMessage}</small>
                        </div>
                    </Form.Group>
                    <Form.Group className="border-top py-4">
                        <Form.Label className="me-3">Category</Form.Label>
                        {categoryRadio}
                    </Form.Group>
                    <div className="text-end">
                        <Button variant="primary" type="submit" className="me-2">Submit</Button>
                        <Button variant="danger" onClick={onClose}>Back</Button>
                    </div>
                </Form>
            </Modal.Body>
        </Modal>
    );
}


function Assessments() {
    const params = useParams();
    const {globalState} = useGlobalState();
    const [examList, setExamList] = useState<assessmentProps[]>([]);
    const [courseInfo, setCourseInfo] = useState<ICourseInfo>();
    const [showModal, setShowModal] = useState<boolean>(false);
    const [badExamConfig, setBadExamConfig] = useState<string>("");
    const navigate = useNavigate();

    const getCourseInfo = useCallback(async () => {
        const infoUrl = getBackendApiUrl("/courses/" + params.course_name + "/info");
        const assessmentUrl = getBackendApiUrl("/courses/" + params.course_name + "/assessments");
        const token = globalState.token;
        const infoResult = await axios.get(infoUrl, {headers: {Authorization: "Bearer " + token}});
        setCourseInfo(infoResult.data.data);
        const assessmentResult = await axios.get(assessmentUrl, {headers: {Authorization: "Bearer " + token}});
        setExamList(assessmentResult.data.data);
    }, [globalState.token, params.course_name]);

    useEffect(() => {
        getCourseInfo().catch();
    }, [getCourseInfo]);

    const [categoryList, setCategoryList] = useState<string[]>([]);

    const getCategoryList = useCallback(async () => {
        const categoryListUrl = getBackendApiUrl("/courses/assessments/config/categories");
        const token = globalState.token;
        const categoryListResult = await axios.get(categoryListUrl, {headers: {Authorization: "Bearer " + token}});
        setCategoryList(categoryListResult.data.data.categories);
    }, [])

    useEffect(() => {
        getCategoryList().catch();
    }, [getCategoryList]);

    const createNewExam = async (categoryName: string, name: string) => {
        const postUrl = getBackendApiUrl("/courses/" + params.course_name + "/assessments");
        const token = globalState.token;
        const data = {
            category_name: categoryName,
            name: name,
        };
        axios.post(postUrl, data, {headers: {Authorization: "Bearer " + token}})
            .then(_ => {
                const configPage = "/courses" + params.course_name + "/examConfig/" + name + "/base";
                setShowModal(false);
                navigate(configPage);
            })
            .catch((error: any) => {
                let response = error.response.data;
                setBadExamConfig(response.error.message);
            });
    }

    const assessmentTable = Table(examList, courseInfo);
    return (
        <AppLayout>
            <Row>
                <TopNavbar brand={courseInfo?.display_name || params.course_name} brandLink={"/courses/"+params.course_name}/>
            </Row>
            <main>
                {courseInfo?.auth_level === "instructor" &&
                    <div className="text-end pe-5">
                        <Button variant="info" className="me-3 text-white" onClick={() => {setShowModal(true);}}>New Exam</Button>
                        <Link to={"questionBank"}><Button variant="primary">Question Bank</Button></Link>
                    </div>
                }
                <Row>
                    <Col xs={{span: "10", offset: "1"}}>
                        <h1>Assessments</h1>
                        {assessmentTable}
                    </Col>
                </Row>
            </main>
            <NewExamConfig show={showModal}
                           onClose={() => {setShowModal(false);}}
                           clearMessage={() => {setBadExamConfig("")}}
                           onSubmit={(categoryName, name) => {createNewExam(categoryName, name).catch();}}
                           categoryList={categoryList}
                           errorMessage={badExamConfig}
            />
        </AppLayout>
    );
}

export default Assessments;
