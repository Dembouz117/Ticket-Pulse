import { PrismaClient } from "@prisma/client";
import { Request, Response } from "express";

import { TicketServiceClient } from "../../proto/entpb/TicketService";
import { SectionServiceClient } from "../../proto/entpb/SectionService";
import { updateTicketHandler } from "../grpc";

import { paymentSerializer } from "../../utilities/helpers/index";

import Stripe from "stripe";
import { Resend } from "resend";
const fs = require("fs");
const path = require("path");

const RESEND_API_KEY = process.env.RESEND_API_KEY || "";
const resend = new Resend(RESEND_API_KEY);

type EmailContentData = {
    orderId: string;
    paidAmount: number;
};

export const webhookHandler =
    (
        prisma: PrismaClient,
        stripe: any,
        ticketClient: TicketServiceClient,
        sectionClient: SectionServiceClient,
        endpointSecret: string
    ) =>
    async (request: Request, response: Response) => {
        console.log("inside webhook handler!");
        const sig = request.headers["stripe-signature"];
        let event;
        const cleanedEndpointSecret = endpointSecret.trim(); // legit dk why, but need to do this even though my secret doesnt have whitespace

        try {
            event = stripe.webhooks.constructEvent(
                request.body,
                sig,
                cleanedEndpointSecret
            );
        } catch (err: any) {
            console.log(err);
            console.log("Webhook Error: " + err.message);
            response
                .status(400)
                .send(`Webhook Error: Unable to process webhook`);
            return;
        }

        console.log("paymentIntent in webhook:");
        const paymentIntent: Stripe.PaymentIntent = event.data.object;
        const paymentIntentId = paymentIntent.id;

        const tickets = JSON.parse(paymentIntent.metadata.tickets);
        const sections = JSON.parse(paymentIntent.metadata.sections);
        const userId = paymentIntent.metadata.userId;
        const email = paymentIntent.metadata.email;
        const paidAmount = paymentIntent.amount / 100;

        // const ticketUuid = tickets[0].ticketId;
        const sectionUuid = sections[0].sectionId;
        const ticketUuidList: Buffer[] = tickets.map((ticket: any) =>
            Buffer.from(ticket.ticketId.replace(/-/g, ""), "hex")
        );

        // Handle the event
        switch (event.type) {
            case "payment_intent.succeeded":
                console.log("successful event!");

                // const ticketUuidBuffer = Buffer.from(
                //     ticketUuid.replace(/-/g, ""),
                //     "hex"
                // );
                const sectionUuidBuffer = Buffer.from(
                    sectionUuid.replace(/-/g, ""),
                    "hex"
                );

                //make CRUD request
                try {
                    const updatedPayment = await prisma.payments.update({
                        where: { payment_id: paymentIntentId },
                        data: { payment_status: "CONFIRMED" },
                    });
                    // console.log("finished making CRUD request!");
                    // console.log(sectionUuid);
                    // console.log(sectionUuid.length);
                    // console.log(ticketUuid);
                    // console.log(ticketUuid.length);

                    // gRPC call to update ticket status to 'BOUGHT' for the user
                    updateTicketHandler(
                        ticketClient,
                        sectionClient,
                        sectionUuidBuffer,
                        userId,
                        ticketUuidList
                    );
                    const paymentSerializable =
                        paymentSerializer(updatedPayment);

                    try {
                        //metadata: tickets, amount, userId, email

                        const emailContent = await generateEmailContent({
                            orderId: paymentIntent.id,
                            paidAmount: paidAmount,
                        });
                        console.log("email:", email);

                        const data = await resend.emails.send({
                            from: "No Reply <no-reply@biddlr.com>",
                            to: [email],
                            subject: "Your Taylor Swift tickets are confirmed!",
                            html: emailContent,
                        });

                        console.log(data);
                    } catch (error) {
                        console.log("email error");
                        console.error(error);
                    }
                    return response.json(paymentSerializable);
                } catch (error: any) {
                    console.error(
                        "Error updating payment status:",
                        error.message
                    );
                    console.log("The entry does not exist yet!");
                }

                break;
            case "payment_intent.requires_action":
                console.log("Payment requires action:", paymentIntent);

                try {
                    console.log(
                        "Attempting to update when it requires action!"
                    );
                    const updatedPayment = await prisma.payments.update({
                        where: { payment_id: paymentIntentId },
                        data: { payment_status: "PENDING_ACTION" },
                    });

                    const paymentSerializable =
                        paymentSerializer(updatedPayment);
                    return response.json(paymentSerializable);
                } catch (error: any) {
                    console.error(
                        "Error updating payment status:",
                        error.message
                    );
                    console.log("The entry does not exist yet!");
                }

                break;
            case "payment_intent.payment_failed":
                console.log("Payment failed:", paymentIntent);
                try {
                    console.log("Attempting to update when it fails!");
                    const updatedPayment = await prisma.payments.update({
                        where: { payment_id: paymentIntentId },
                        data: { payment_status: "FAILED" },
                    });

                    return response.json(updatedPayment);
                } catch (error: any) {
                    console.error(
                        "Error updating payment status:",
                        error.message
                    );
                    console.log("The entry does not exist yet!");
                }
                break;

            default:
                console.log(`Unhandled event type ${event.type}`);
        }

        response.send();
    };

const generateEmailContent = async (data: EmailContentData) => {
    // Read the content of the HTML template
    let emailContent = fs.readFileSync("email.html", "utf8");
    emailContent = emailContent.replace(/\[Order ID\]/g, data.orderId);
    emailContent = emailContent.replace(/\[Paid Amount\]/g, data.paidAmount);
    emailContent = emailContent.replace(
        /\[Support Email Address\]/g,
        "angruiyan@gmail.com"
    );

    // Return the final email content
    return emailContent;
};
