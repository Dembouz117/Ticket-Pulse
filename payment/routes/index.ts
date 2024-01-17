import express from "express";
import {
    createPaymentIntentHandler,
    createPaymentStatusHandler,
} from "../controllers/stripe/index";
import { webhookHandler } from "../controllers/webhook/index";
import { pollingHandler } from "../controllers/polling/index";

import {
    prisma,
    stripe,
    ticketClient,
    sectionClient,
    endpointSecret,
} from "../config/index";

const router = express.Router();

//when buyer clickes "Yes" to acknowledge suggested tickets
router.post(
    "/payment/create-payment-intent",
    express.json(),
    createPaymentIntentHandler(prisma, stripe, ticketClient)
);
//when buyer clickes "Pay" after filling his credit card details
router.post(
    "/payment/create-payment-status/:paymentId",
    express.json(),
    createPaymentStatusHandler(prisma, stripe, ticketClient)
);

//this endpoint is for polling after a payment attempt has been made
router.get(
    "/payment/payment-status/:paymentId",
    express.json(),
    pollingHandler(prisma)
);

router.post(
    "/payment/webhook",
    express.raw({ type: "application/json" }),
    webhookHandler(prisma, stripe, ticketClient, sectionClient, endpointSecret)
);

router.get("/payment/health", express.json(), (req, res) => {
    return res.status(200).json({ status: "healthy" });
});

console.log(endpointSecret);

export default router;
