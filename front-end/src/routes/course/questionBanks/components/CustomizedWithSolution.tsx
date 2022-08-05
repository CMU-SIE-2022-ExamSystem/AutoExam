import { Form } from "react-bootstrap";
import { subQuestionDataType } from "../../../../components/questionTemplate/subQuestionDataType";

const CustomizedWithSolution = ({index, subQuestion}: {index: number, subQuestion: subQuestionDataType}) => {
    return (
        <Form.Group className="mb-3">
            <Form.Label>
                {index + ". " + subQuestion.description}
                {/* {index + ". "}<div dangerouslySetInnerHTML={{__html: subQuestion.description}}/> */}
            </Form.Label>
            <br/>
            
            {
                subQuestion.blanks.map((blank, index) => {
                    if (blank.multiple) {
                        return (
                        <div key={index}>
                            <Form.Label>{"Sub " + (index + 1) + ": multiple choice"}</Form.Label>
                            {
                                subQuestion.choices[index]?.map((choice, choiceIdx) => 
                                    <Form.Label key={choiceIdx}>{choice.content}</Form.Label>
                                )
                            }
                        </div>)
                    } else {
                        return (<Form.Label key={index}>{"Blank " + (index + 1) + ": " + blank.type}</Form.Label>)
                    }
                })
            }
            <Form.Control disabled readOnly key={index} value={subQuestion.solutions[0]}/>
        </Form.Group>
    );
}

export default CustomizedWithSolution;
