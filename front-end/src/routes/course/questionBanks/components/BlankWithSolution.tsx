import { Form } from "react-bootstrap";
import { subQuestionDataType } from "../../../../components/questionTemplate/subQuestionDataType";

const BlankWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    return (
        <Form.Group className="mb-3">
            <Form.Label>{index + ". " + subQuestion.description}</Form.Label>
            {/* {
                subQuestion.solutions[0].map((solution, index) => (
                    <Form.Control disabled readOnly key={index} value={solution}/>
                ))
            } */}
            <Form.Control disabled readOnly key={index} value={subQuestion.solutions[0]}/>
        </Form.Group>
    );
}

export default BlankWithSolution;
