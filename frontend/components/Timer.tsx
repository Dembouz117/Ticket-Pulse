import React, { useEffect, useState, useRef } from "react";
import { DateTime, Duration } from "luxon";

interface TimerProps {
    startDateTime: number;
    endDateTime: number;
}

const Timer = ({ startDateTime, endDateTime }: TimerProps) => {
    const [daysLeft, setDaysLeft] = useState(0);
    const [hoursLeft, setHoursLeft] = useState(0);
    const [minutesLeft, setMinutesLeft] = useState(0);
    const [secondsLeft, setSecondsLeft] = useState(0);
    const [bidStarted, setBidStarted] = useState(false);
    const [bidEnded, setBidEnded] = useState(false);

    const countdownIntervalRef = useRef<number | undefined>();

    useEffect(() => {
        const updateCountdown = () => {
            let diff: Duration;
            if (!bidStarted) {
                diff = DateTime.fromSeconds(startDateTime).diffNow();
                setBidStarted(
                    DateTime.fromSeconds(startDateTime).diffNow().toMillis() <=
                        0
                );
            } else {
                diff = DateTime.fromSeconds(endDateTime).diffNow();
                setBidEnded(
                    DateTime.fromSeconds(endDateTime).diffNow().toMillis() <= 0
                );
            }

            const diffObj = diff
                .shiftTo("days", "hours", "minutes", "seconds")
                .toObject();

            setDaysLeft(Math.floor(diffObj.days || 0));
            setHoursLeft(Math.floor(diffObj.hours || 0));
            setMinutesLeft(Math.floor(diffObj.minutes || 0));
            setSecondsLeft(Math.floor(diffObj.seconds || 0));
        };

        updateCountdown(); // Run initially to set the correct countdown immediately

        countdownIntervalRef.current = window.setInterval(
            updateCountdown,
            1000
        ) as unknown as number;

        return () => {
            if (countdownIntervalRef.current !== undefined) {
                clearInterval(countdownIntervalRef.current);
                countdownIntervalRef.current = undefined;
            }
        };
    }, [startDateTime, endDateTime, bidStarted]);

    return (
        <div className="w-full">
            {bidEnded ? (
                <div className="text-center py-4 w-full rounded-2xl bg-light-button-active bg-opacity-60 border-2 border-dark-primary">
                    <div className="text-2xl font-extrabold mb-1 md:text-2xl text-white">
                        Ending Phase Active
                    </div>
                </div>
            ) : (
                <div className="text-white flex flex-row align-middle justify-between">
                    <div className="flex flex-col align-middle text-center py-8 w-1/5 rounded-2xl bg-white text-button">
                        <div className="text-2xl font-semibold mb-1 md:text-5xl">
                            {daysLeft}
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center py-8 w-1/5 rounded-2xl bg-white text-button">
                        <div className="text-2xl font-semibold mb-1 md:text-5xl">
                            {hoursLeft}
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center py-8 w-1/5 rounded-2xl bg-white text-button">
                        <div className="text-2xl font-semibold mb-1 md:text-5xl">
                            {minutesLeft}
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center py-8 w-1/5 rounded-2xl bg-white text-button">
                        <div className="text-2xl font-semibold mb-1 md:text-5xl">
                            {secondsLeft}
                        </div>
                    </div>
                </div>
            )}
            {!bidEnded && (
                <div className="text-white flex flex-row align-middle justify-between mt-2">
                    <div className="flex flex-col align-middle text-center w-1/5">
                        <div className="text-white md:text-2xl">
                            Days
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center w-1/5">
                        <div className="text-white md:text-2xl">
                            Hours
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center w-1/5">
                        <div className="text-white md:text-2xl">
                            Mins
                        </div>
                    </div>
                    <div className="flex flex-col align-middle text-center w-1/5">
                        <div className="text-white md:text-2xl">
                            Secs
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Timer;
