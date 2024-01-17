import { useRouter } from "next/router";
import { useState, useEffect } from "react";
import axios from "axios";

import { Spinner } from "@chakra-ui/react";
import { useToast } from "@chakra-ui/react";
import { getPaymentApiUrl } from "@/utilities/common";

interface PaymentStatusProps {
    paymentId: any;
}

function PaymentStatus() {
    const router = useRouter();

    const toast = useToast({
        position: "top",
    });

    const paymentId = router.query.paymentId;

    const [status, setStatus] = useState<string>("");
    const [error, setError] = useState("");

    useEffect(() => {
        // Function to fetch payment status =>

        const fetchStatus = async () => {
            try {
                const response = await axios.get(
                    `${getPaymentApiUrl()}/payment-status/${paymentId}`
                );
                setStatus(response.data.status);

                console.log("in fetch, before confirmation");
                console.log(response.data);
                console.log(response.data.confirmed);
                const resStatus = response.data.status;
                if (resStatus === "CONFIRMED") {
                    console.log("In if statement!");
                    clearInterval(intervalId); // Clear interval if the status is in final state
                    setStatus(resStatus);
                    toast({
                        title: "Payment Status",
                        description:
                            "Payment was successful! Congratulations on your tickets! 🎉",
                        status: "success",
                        duration: 5000,
                        isClosable: true,
                        position: "top",
                        // className:"bg-button",
                        onCloseComplete: () => router.push("/"),
                    });
                } else if (resStatus === "FAILED") {
                    console.log("In else statement!");
                    clearInterval(intervalId);
                    setStatus(resStatus);
                    toast({
                        title: "Payment Status",
                        description:
                            "Payment was unsuccessful. Redirecting you to the home page...",
                        status: "error",
                        duration: 5000,
                        isClosable: true,
                        position: "top",
                        // className:"bg-button",
                        onCloseComplete: () => router.push("/"),
                    });
                }
            } catch (err) {
                console.error("Error fetching payment status", err);
                setError("Error fetching payment status");
            }
        };

        // Initial fetch
        fetchStatus();

        // Set up an interval to poll the status every 30 seconds
        const intervalId = setInterval(fetchStatus, 3000);

        // Clean up the interval on component unmount
        return () => clearInterval(intervalId);
    }, [paymentId]);

    //for the content in the modal to reflect outcome
    let displayOutcome: any;
    if (status === "CONFIRMED") {
        displayOutcome = (
            <p>
                Your tickets have been secured! Redirecting you to the home
                page...
            </p>
        );
    } else if (status === "FAILED") {
        displayOutcome = (
            <p>Payment was unsuccessful. Redirecting you to the home page...</p>
        );
    } else {
        displayOutcome = <Spinner />;
    }
    return (
        <div className="min-h-screen flex items-center justify-center bg-background">
            <div className="bg-white p-8 rounded-lg shadow-md text-center">
                <svg
                    className="mb-4 w-16 h-16 mx-auto text-button"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="M6 19V9a9 9 0 0 1 12 0v10m-6-13.5V5m0 15.5v-2.5m0-10a2 2 0 1 0 0 4 2 2 0 1 0 0-4z"
                    ></path>
                </svg>
                {status !== "CONFIRMED" && status !== "FAILED" ? (
                    <h2 className="mb-4 text-xl font-semibold text-gray-700">
                        Processing your payment...
                    </h2>
                ) : (
                    <h2></h2>
                )}
                <div className="flex-col space-y-4">
                    {status !== "CONFIRMED" && status !== "FAILED" && (
                        <p className="text-gray-500">
                            Please wait while we confirm your payment and secure
                            your tickets. Do not refresh or close this page.
                        </p>
                    )}
                    {displayOutcome}
                </div>
            </div>
        </div>
    );
}

export default PaymentStatus;
