import React, {useCallback, useEffect, useState} from 'react';
import {Row, Col, Button, Modal, Form} from 'react-bootstrap';
import {Link, useNavigate, useParams} from "react-router-dom";
import TopNavbar from "../../components/TopNavbar";
import AppLayout from "../../components/AppLayout";
import {getBackendApiUrl} from "../../utils/url";
import axios from "axios";
import {useGlobalState} from "../../components/GlobalStateProvider";
import moment from 'moment';
import BaseCourseRelationshipManageModal from "./baseCourse/BaseCourseRelationshipManageModal";

interface assessmentProps {
    name: string;
    display_name: string;
    start_at: string;
    due_at: string;
    end_at: string;
    category_name: string;
    grading_deadline?: string;
    autoexam: boolean;
    autolab: boolean;
    draft: boolean;
    submitted: boolean;
    can_submit: boolean;
}

interface ICourseInfo {
    name: string;
    display_name: string;
    auth_level: string;
}

interface extAssessmentProps extends assessmentProps {
    permission: boolean
}

interface tagProps {
    id: string;
    name: string;
}

const AssessmentRow = ({name, display_name, start_at, due_at, permission, draft, submitted, can_submit}: extAssessmentProps) => {
    let now = moment();
    let start = moment(start_at);
    let end = moment(due_at);
    let startTime = start.format("MMMM Do YYYY, h:mm:ss a");
    let dueTime = end.format("MMMM Do YYYY, h:mm:ss a");

    let beforeExamTime = now.diff(start) < 0;
    let inExamTime = now.diff(start) >= 0 && end.diff(now) >= 0;
    let afterExamTime = end.diff(now) < 0;

    let actionList = [];
    if (permission) {
        actionList.push(<Link to={"examConfig/" + name} key="_EditExam" className="btn btn-success m-1">Edit Exam</Link>)
        //actionList.push(<Link to={"exams/" + name} key="_ProctorExam" className="btn btn-primary m-1">Proctor Exam</Link>)
    } else {
        if (beforeExamTime) {
            actionList.push(<span>Remaining {now.to(start, true)}</span>)
        } else if (inExamTime && can_submit) {
            actionList.push(<Link to={"exams/" + name} key="_TakeExam" className="btn btn-primary m-1">Take Exam</Link>);
        }
        if (afterExamTime && !submitted) {
            actionList.push(<span>No Submissions</span>);
        }
        if (submitted) {
            actionList.push(<Link to={"examResults/" + name} key="_ExamResults" className="btn btn-success m-1">Exam Results</Link>);
        }
    }

    let trClassName = "align-middle";
    if (draft) {
        trClassName += " bg-light";
    }

    return (
        <tr className={trClassName}>
            <th scope="row" className="text-center">{display_name + (draft ? " (Draft)" : "")}</th>
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
                                      onChange={clearMessage}
                        />
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
    const [showNexExamModal, setShowNexExamModal] = useState<boolean>(false);
    const [showBaseCourseModal, setShowBaseCourseModal] = useState(false);
    const [badExamConfig, setBadExamConfig] = useState<string>("");
    const navigate = useNavigate();

    const getCourseInfo = useCallback(async () => {
        const infoUrl = getBackendApiUrl("/courses/" + params.course_name + "/info");
        const assessmentUrl = getBackendApiUrl("/courses/" + params.course_name + "/assessments");
        const token = globalState.token;
        const infoResult = await axios.get(infoUrl, {headers: {Authorization: "Bearer " + token}});
        setCourseInfo(infoResult.data.data);
        if (infoResult.data.data.base_course.length === 0) setShowBaseCourseModal(true);
        const assessmentResult = await axios.get(assessmentUrl, {headers: {Authorization: "Bearer " + token}});
        if (assessmentResult.data.data) setExamList(assessmentResult.data.data);
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
    }, [globalState.token])

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
                const configPage = "/courses/" + params.course_name + "/examConfig/" + name;
                setShowNexExamModal(false);
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
                        <Button variant="secondary" className="me-3 text-white" onClick={() => {setShowBaseCourseModal(true);}}>Base Course Edit</Button>
                        <Button variant="info" className="me-3 text-white" onClick={() => {setShowNexExamModal(true);}}>New Exam</Button>
                        <Link to={"questionBank/null"}><Button variant="primary">Question Bank</Button></Link>
                    </div>
                }
                <Row>
                    <Col xs={{span: "10", offset: "1"}}>
                        <h1>Assessments</h1>
                        {assessmentTable}
                    </Col>
                </Row>
            </main>
            <NewExamConfig show={showNexExamModal}
                           onClose={() => {setShowNexExamModal(false);}}
                           clearMessage={() => {setBadExamConfig("")}}
                           onSubmit={(categoryName, name) => {createNewExam(categoryName, name).catch();}}
                           categoryList={categoryList}
                           errorMessage={badExamConfig}
            />
            {courseInfo?.auth_level === 'instructor' && <BaseCourseRelationshipManageModal show={showBaseCourseModal} toClose={() => setShowBaseCourseModal(false)} />}
        </AppLayout>
    );
}

export default Assessments;
