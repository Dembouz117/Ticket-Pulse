import { useState, useEffect } from "react";
import { useAtom, useAtomValue } from "jotai";

//stripe-related
import getStripe from "@/utilities/get-stripejs";
import { Elements } from "@stripe/react-stripe-js";

//stripe dependency to toggle the loading and also store total amount
import {
    showStripeCheckoutAtom,
    ticketResponseAtom,
    errorModalAtom,
    cartEntryAtom,
} from "@/store/index";
import { userAtom } from "@/state/globals";

//additional components and styles
import {
    Modal,
    ModalCloseButton,
    ModalContent,
    ModalOverlay,
} from "@chakra-ui/react";
import LoadingSeats from "./LoadingSeats";
import CheckOut from "./CheckOut";
import { StripeElementsOptions } from "@stripe/stripe-js";

import { appearance } from "@/styles/stripe/stripeAppearance";
import { getPaymentApiUrl } from "@/utilities/common";

import PaymentModalContent from "@/components/payment/PaymentModalContent";

interface paymentModalProps {
    open: boolean;
    onClose: () => void;
}
const stripePromise = getStripe();

const PaymentModal = ({ open, onClose }: paymentModalProps) => {
    const [clientSecret, setClientSecret] = useState<string>("");
    const [loading, setIsLoading] = useState<boolean>(true);
    const [accepted, setAccepted] = useState<boolean>(false);
    const [errorModal, setErrorModal] = useAtom(errorModalAtom);
    //to get the total amount of the tickets
    const [showStripeCheckout, setShowStripeCheckout] = useAtom(
        showStripeCheckoutAtom
    );
    const [user, setUser] = useAtom(userAtom);
    const [cartEntry, setCartEntry] = useAtom(cartEntryAtom);

    //tickets are set in PreCheckOutModal.tsx
    const tickets = useAtomValue(ticketResponseAtom);

    useEffect(() => {
        // Try to create PaymentIntent as soon as the page loads
        const filteredSectionsInCart = cartEntry.filter(
            (el) => el.quantity > 0 && el.sectionId.length > 0
        );
        if (tickets.length > 0) {
            console.log("in payment modal user check:");
            console.log(user!.email);
            fetch(`${getPaymentApiUrl()}/create-payment-intent`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    tickets: tickets,
                    amount: showStripeCheckout.amount*100,
                    userId: user!.userId,
                    sections: filteredSectionsInCart,
                    email: user!.email,
                }),
                credentials: "include",
            })
                .then(async (res) => {
                    if (!res.ok) {
                        const errorData = await res.json();
                        const errorMessage = errorData.error;

                        throw new Error(errorMessage);
                    }
                    return res.json();
                })
                .then((data) => {
                    setClientSecret(data.clientSecret);
                })
                .catch((error) => {
                    console.log("inside this error block");
                    setErrorModal({ visible: true, message: error.message });
                });
        }
    }, [tickets]);

    const options: StripeElementsOptions = {
        clientSecret,
        appearance,
    };

    const handleClose = () => {
        onClose();
    };

    const acceptSeatHandler = (accept: boolean) => {
        //set loading to false, mock the transfer to the payment element
        if (accept) {
            setAccepted(true);
        } else {
            onClose();
        }
        setIsLoading(false);
    };

    return (
        <PaymentModalContent
            handleClose={onClose}
            clientSecret={clientSecret}
            loading={loading}
            accepted={accepted}
            onEndLoading={acceptSeatHandler}
            stripePromise={stripePromise}
            options={options}
        />
    );
};

export default PaymentModal;
