import {Form} from "react-bootstrap";
import {choiceDataType, subQuestionDataType} from "../../../../components/questionTemplate/subQuestionDataType";

const ChoiceWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    const choices = (subQuestion.choices[0] as choiceDataType[]).map((choice) => {
        return (
            <div key={choice.choice_id}>
                <Form.Label>{choice.choice_id + ". " + choice.content}</Form.Label>
                <br/>
            </div>
        );
    })

    return (
        <Form.Group className="mb-3">
            <Form.Label>{index + ". " + subQuestion.description}</Form.Label><br/>
            {choices}
            <Form.Control disabled readOnly value={subQuestion.solutions[0]}/>
        </Form.Group>
    );
}

export default ChoiceWithSolution;
