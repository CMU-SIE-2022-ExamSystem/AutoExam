import {Alert} from "react-bootstrap";
import React from "react";

/**
 * An alert component that appears at the right bottom of the screen.
 * @param variant The background of the alert. See React-Bootstrap Alert documentation
 * @param content The contents display on the component
 * @param show  Whether this component is visible
 * @param onClose  Callback function when the alert is closed / dismissed. Usually you set the boolean variable that contains "show" to false.
 */
export const RightBottomAlert = ({variant, content, show, onClose} : {variant: string, content: string, show: boolean, onClose: () => void}) => {
    return (
        <div className="position-absolute bottom-0 end-0 p-3 text-start">
            <Alert variant={variant} show={show} onClose={onClose} dismissible>
                {content}
            </Alert>
        </div>
    );
}
