import React, { useEffect, useRef, useState } from "react";
import { Box, Input } from "@chakra-ui/react";
import axios from "axios";
import { getAuthApiUrl } from "@/utilities/common";
import { useSetAtom } from "jotai";
import { userAtom } from "@/state/globals";
import OtpInput from "react-otp-input";

interface TwoFactorStepProps {
    onClose: () => void;
    setStep: (step: string) => void;
    email: string;
    password: string;
}

const TwoFactorStep: React.FC<TwoFactorStepProps> = ({
    onClose,
    setStep,
    email,
    password,
}) => {
    const [otp, setOtp] = useState("");
    const [hasError, setHasError] = useState<boolean>(false);
    const [remainingSeconds, setRemainingSeconds] = useState<number>(30);
    const setUser = useSetAtom(userAtom);
    const inputs = useRef<(HTMLInputElement | null)[]>([]);

    //for timer
    useEffect(() => {
        const intervalId = setInterval(() => {
            setRemainingSeconds((prevSeconds) => prevSeconds - 1);
        }, 1000);

        if (remainingSeconds <= 0) {
            clearInterval(intervalId);
        }

        return () => {
            clearInterval(intervalId);
        };
    }, [remainingSeconds]);

    //handle resend code
    const handleResend = async () => {
        console.log("Making API call to resend code...");
        try {
            const response = await axios.post(
                `${getAuthApiUrl()}/login`,
                {
                    email: email,
                    password: password,
                },
                { withCredentials: true }
            );

            if (response.status === 200) {
                setHasError(false);
            }
        } catch (error) {
            setHasError(true);
        }
        setRemainingSeconds(30); // Reset timer
    };

    const onSubmit = async () => {
        console.log("Calling API with code:", otp);
        try {
            const response = await axios.post(
                `${getAuthApiUrl()}/otp`,
                {
                    email: email,
                    otpCode: otp,
                },
                { withCredentials: true }
            );

            if (response.status === 200) {
                setHasError(false);
                setUser({ email: email });
                onClose();
            }
        } catch (error) {
            setHasError(true);
        }
    };

    return (
        <Box w="full" p={4}>
            <div className="flex flex-col align-middle justify-start p-3">
                <div className="text-3xl mb-2">Verification Required</div>
                <div className="text-[#AFAFAF] mb-16">
                    Enter the 6-digit code sent to your phone
                </div>

                <div className="flex gap-2 mb-16">
                    <OtpInput
                        value={otp}
                        onChange={setOtp}
                        numInputs={6}
                        renderInput={(props) => <input {...props} />}
                        containerStyle={{
                            display: "flex",
                            justifyContent: "center",
                            height: "65px",
                            color: "black",
                        }}
                        inputStyle={{
                            width: "50px",
                            height: "65px",
                            borderRadius: "10px",
                            border: "2px solid #AFAFAF",
                            fontSize: "36px",
                            color: "black",
                            margin: "0 10px",
                            outline: "none",
                            textAlign: "center",
                        }}
                    />
                </div>

                <button
                    onClick={onSubmit}
                    className="bg-button hover:bg-button-hover text-white rounded-lg w-full py-2 mb-1"
                >
                    Log In
                </button>
                <div className="text-sm text-[#AFAFAF]">
                    {remainingSeconds > 0 ? (
                        `Resend Code in ${remainingSeconds}s`
                    ) : (
                        <a
                            href="#"
                            onClick={handleResend}
                            className="hover:underline"
                        >
                            Resend Code
                        </a>
                    )}
                </div>
            </div>
        </Box>
    );
};

export default TwoFactorStep;
