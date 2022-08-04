import { Form } from "react-bootstrap";
import { subQuestionDataType } from "../../../../components/questionTemplate/subQuestionDataType";

const CustomizedWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    return (
        <Form.Group className="mb-3">
            <Form.Label>
                {index + ". "}
                <div dangerouslySetInnerHTML={{__html: subQuestion.description}}/>
            </Form.Label>
            <Form.Control disabled readOnly key={index} value={subQuestion.solutions[0]}/>
        </Form.Group>
    );
}

export default CustomizedWithSolution;
