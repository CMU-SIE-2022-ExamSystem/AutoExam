import {Alert} from "react-bootstrap";
import React from "react";

export const RightBottomAlert = ({variant, content, show, onClose} : {variant: string, content: string, show: boolean, onClose: () => void}) => {
    return (
        <div className="position-absolute bottom-0 end-0 p-3 text-start">
            <Alert variant={variant} show={show} onClose={onClose} dismissible>
                {content}
            </Alert>
        </div>
    );
}
