import React, {useEffect, useState} from 'react';
import {Card, Button, Collapse} from 'react-bootstrap';

const computeDiffTime = (targetTime: string) => {
    const currentTime = Date.now();
    const finishTime = new Date(targetTime).getTime();
    const difference = finishTime - currentTime; // This difference is in milliseconds.
    return difference > 0 ? difference : 0;
}

const formTime = (remainingTime: number) => {
    // remainingTime is the remaining milliseconds to the targetTime.

    // Convert to seconds, use ceiling to ignore milliseconds part.
    const remainingSeconds = Math.ceil(remainingTime / 1000);

    const seconds = remainingSeconds % 60;
    const minutes = Math.floor(remainingSeconds / 60) % 60;
    const hours = Math.floor(remainingSeconds / 3600);

    const zeroPaddingSeconds = seconds < 10 ? `0${seconds}` : seconds;
    const zeroPaddingMinutes = minutes < 10 ? `0${minutes}` : minutes;
    const zeroPaddingHours = hours < 10 ? `0${hours}` : hours;

    return `${zeroPaddingHours}:${zeroPaddingMinutes}:${zeroPaddingSeconds}`;
}

const CountdownTimer = ({targetTime, callback} : {targetTime: string, callback: Function}) => {
    const [countdownTime, setTime] = useState(computeDiffTime(targetTime));
    const [displayState, setDisplayState] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => {
            const diffTime = computeDiffTime(targetTime);
            if (diffTime !== countdownTime) {
                setTime(diffTime);
                if (diffTime <= 0) callback();
            }
        }, 1000);

        return () => clearTimeout(timer);
    })

    const toggleDisplayState = () => setDisplayState(!displayState);

    return (
        <>
            <Card className="text-start w-100">
                <Card.Header>Remaining Time</Card.Header>
                <Card.Body className="d-flex flex-column text-center">
                    <Collapse in={displayState}>
                        <Card.Title>
                            {formTime(countdownTime)}
                        </Card.Title>
                    </Collapse>
                    <Button variant="primary" size="sm" onClick={toggleDisplayState}>{displayState ? "Hide" : "Show"}</Button>
                </Card.Body>
            </Card>
        </>
    );
}

export default CountdownTimer;
