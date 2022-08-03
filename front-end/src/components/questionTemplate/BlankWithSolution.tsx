import { Form } from "react-bootstrap";
import { subQuestionDataType } from "./subQuestionDataType";

const BlankWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    return (
        <Form.Group className="mb-3">
            <Form.Label>{index + ". " + subQuestion.description}</Form.Label>
            <Form.Control disabled readOnly value={subQuestion.solutions[0]}/>
        </Form.Group>
    );
}

export default BlankWithSolution;
