import { useState, useEffect } from "react";
import { useAtom, useAtomValue } from "jotai";
import {
    showStripeCheckoutAtom,
    ticketResponseAtom,
    errorModalAtom,
} from "@/store/index";
import { userAtom } from "@/state/globals";

import { paymentIntentAtom } from "@/store/index";
import axios from "axios";

import {
    PaymentElement,
    useStripe,
    useElements,
} from "@stripe/react-stripe-js";
import { stripeConfirmPayment } from "@/utilities/check-payment-confirmation";
import cardElementStyle from "@/styles/stripe/cardElementStyle";

import { Spinner } from "@chakra-ui/react";
import {
    PaymentIntent,
    PaymentIntentResult,
    StripePaymentElementOptions,
} from "@stripe/stripe-js";
import { getPaymentApiUrl } from "@/utilities/common";

interface checkoutProps {
    clientSecret: string;
}

const CheckOut = ({ clientSecret }: checkoutProps) => {
    const stripe = useStripe();
    const elements = useElements();

    const [isLoading, setIsLoading] = useState(false);
    const [axiosCallStatus, setAxiosCallStatus] = useState<boolean>(false);

    const [attemptedPayment, setAttemptedPayment] = useState(false);
    const [paymentIntent, setPaymentIntent] = useAtom(paymentIntentAtom);
    const [errorModal, setErrorModal] = useAtom(errorModalAtom);

    const [stripeCheckout, setStripeCheckout] = useAtom(showStripeCheckoutAtom);
    const [user, setUser] = useAtom(userAtom);
    const tickets = useAtomValue(ticketResponseAtom);
    

    const paymentElementOptions: StripePaymentElementOptions = {
        layout: "tabs",
    };

    useEffect(() => {
        setIsLoading(true);
        if (!stripe || !clientSecret) {
            return;
        }

        if (axiosCallStatus) {
        }
        stripe
            .retrievePaymentIntent(clientSecret)
            .then(({ paymentIntent }: PaymentIntentResult) => {
                if (paymentIntent !== undefined) {
                    setPaymentIntent(paymentIntent);
                    if (axiosCallStatus) {
                        switch (paymentIntent.status) {
                            case "succeeded":
                                setErrorModal({
                                    visible: true,
                                    message: "Payment succeeded!",
                                });
                                break;
                            case "processing":
                                setErrorModal({
                                    visible: true,
                                    message: "Your payment is processing.",
                                });
                                break;
                            case "requires_payment_method":
                                setErrorModal({
                                    visible: true,
                                    message:
                                        "Your payment was not successful, please try again.",
                                });
                                break;
                            default:
                                setErrorModal({
                                    visible: true,
                                    message: "Something went wrong.",
                                });
                                break;
                        }
                    }
                }
            });

        setIsLoading(false);
    }, [stripe, attemptedPayment]);

    const handleSubmit = async (e: any) => {
        e.preventDefault();

        if (!stripe || !elements) {
            return;
        }

        setIsLoading(true);
        console.log("Pay now button clicked");

        const createPaymentHandler = async () => {
            try {
                if (
                    tickets.length > 0 &&
                    paymentIntent &&
                    stripeCheckout.amount > 0
                ) {
                    console.log("tried to pay boss");
                    const res = axios
                        .post(
                            `${getPaymentApiUrl()}/create-payment-status/${
                                paymentIntent!.id
                            }`,
                            {
                                payment_id: paymentIntent!.id,
                                payment_provider: "stripe",
                                payment_amount: stripeCheckout.amount,
                                payment_status: "UNCONFIRMED",
                                payment_user: user!.userId,
                                payment_ticket: tickets[0].ticketId,
                            }
                        )
                        .then(async (data) => {
                            setAxiosCallStatus(true);
                            setIsLoading(true);
                            stripeConfirmPayment(
                                stripe,
                                elements,
                                paymentIntent,
                                setErrorModal
                            );
                        })
                        .catch((error) => {
                            if (error.response) {
                                setErrorModal({
                                    visible: true,
                                    message: error.message,
                                });
                            } else if (error.request) {
                                setErrorModal({
                                    visible: true,
                                    message: error.message,
                                });
                            } else {
                                setErrorModal({
                                    visible: true,
                                    message: error.message,
                                });
                            }
                        });
                    console.log("Payment status added to DB: ", res);
                }
            } catch (error: any) {
                setErrorModal({ visible: true, message: error.message });
            }
        };

        createPaymentHandler();
        setAttemptedPayment(true);
        setIsLoading(false);
    };

    return (
        <div>
            <form id="payment-form" onSubmit={handleSubmit}>
                <PaymentElement
                    id="payment-element"
                    options={paymentElementOptions}
                />
                <div className="mt-3 mb-1 underline font-semibold">
                    Total Amount:&nbsp;&nbsp;&nbsp;${stripeCheckout.amount}
                </div>
                <button
                    disabled={isLoading || !stripe || !elements}
                    id="submit"
                    className="px-4 py-2 bg-button mt-4 rounded-md"
                    type="submit"
                >
                    {isLoading ? (
                        <Spinner />
                    ) : (
                        <span id="button-text">Pay Now</span>
                    )}
                </button>

                {attemptedPayment && errorModal.message && (
                    <div id="payment-message">{errorModal.message}</div>
                )}
            </form>
        </div>
    );
};

export default CheckOut;
