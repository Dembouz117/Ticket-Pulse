import { Stripe, StripeElements, PaymentIntent } from "@stripe/stripe-js";
import { getFrontEndUrl } from "@/utilities/common"

export const stripeConfirmPayment = async (stripe: Stripe, elements: StripeElements, paymentIntent: PaymentIntent, setError: (error: any) => void) => {


    const { error } = await stripe.confirmPayment({
        elements,
        confirmParams: {
            return_url: `${getFrontEndUrl()}/payment-status/${
                paymentIntent!.id
            }`,
        },
    });
    if (error.type === "card_error" || error.type === "validation_error") {
        if (error.message) {
            setError({visible: true, message:"There is something wrong with the card you input. Please check your input."});
        }
    } else {
        setError({visible: true, message:"An unexpected error occurred."});
    }
}
