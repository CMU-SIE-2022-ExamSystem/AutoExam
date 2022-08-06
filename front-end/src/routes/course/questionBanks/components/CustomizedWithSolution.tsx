import { Form } from "react-bootstrap";
import { subQuestionDataType } from "../../../../components/questionTemplate/subQuestionDataType";

const CustomizedWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    return (
        <Form.Group className="mb-3">
            <Form.Label>
                {index + ". " + subQuestion.description}
                {/* {index + ". "}<div dangerouslySetInnerHTML={{__html: subQuestion.description}}/> */}
            </Form.Label>

            {subQuestion.blanks.map((blank, index) => {
                let choices = subQuestion.choices[index];
                if (choices !== null) {
                    return (
                        <div key={index}>
                            <Form.Label>{(index + 1) + (blank.multiple === true ? ". Multiple" : ". Single") + " Choice"}</Form.Label>
                            {choices.map((choice) => (
                                <div key={choice.choice_id}>
                                    <Form.Label>{choice.choice_id + ". " + choice.content}</Form.Label>
                                    <br/>
                                </div>
                            ))}
                            <Form.Control disabled readOnly value={subQuestion.solutions[index]}/>
                        </div>
                    );
                } else {
                    return (
                        <div key={index}>
                            <Form.Label>{(index + 1) + (blank.type === "string"? ". Blank" : ". Code")}</Form.Label>
                            <Form.Control disabled readOnly value={subQuestion.solutions[index]}/>
                        </div>
                    )
                }
            })}
        </Form.Group>
    );
}

export default CustomizedWithSolution;
