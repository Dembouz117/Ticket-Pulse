import {
    Modal,
    ModalCloseButton,
    ModalContent,
    ModalOverlay,
} from "@chakra-ui/react";

import LoadingSeats from "./LoadingSeats";
import CheckOut from "./CheckOut";
import { Elements } from "@stripe/react-stripe-js";

import { Stripe, StripeElementsOptions } from '@stripe/stripe-js';

interface PaymentModalContentProps{
    handleClose: () => void,
    clientSecret: string,
    loading: boolean,
    accepted: boolean,
    onEndLoading: (accept: boolean) => void,
    stripePromise: Promise<Stripe | null>,
    options: StripeElementsOptions
}

const PaymentModalContent = ({handleClose, clientSecret, loading, accepted, onEndLoading, stripePromise, options}: PaymentModalContentProps) => {
  return (
    <Modal
    isOpen={true}
    onClose={handleClose}
    closeOnOverlayClick={true}
>
    <ModalOverlay />
    <ModalContent
        bg={"#1F1F1F"}
        textColor={"white"}
        className="py-12 px-4 h-auto"
        role="checkoutModal"
    >
        <ModalCloseButton />
        {clientSecret && !loading && accepted && (
            <Elements options={options} stripe={stripePromise}>
                <CheckOut clientSecret={clientSecret} />
            </Elements>
        )}
        {loading && <LoadingSeats onEndLoading={onEndLoading} />}
    </ModalContent>
</Modal>
  )
}

export default PaymentModalContent;