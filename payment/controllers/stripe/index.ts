import { PrismaClient } from "@prisma/client";
import { Request, Response } from "express";
import { TicketServiceClient } from "../../proto/entpb/TicketService"; 
import { checkTicketsReservedHandler } from "../grpc/index";
import { paymentSerializer } from "../../utilities/helpers";
import { MAXIMUM_TICKETS_PER_USER } from "../../config";
import Stripe from "stripe";


//when you click "yes" to confirm the tickets suggested
export const createPaymentIntentHandler = (prisma : PrismaClient, stripe: Stripe, client: TicketServiceClient) => async (req : Request, res : Response) => {

    console.log("This is the create-payment-intent endpoint");
    const data = req.body;

    console.log("cookies:");
    console.log(req.cookies);

    if (data.tickets.length > MAXIMUM_TICKETS_PER_USER){
        return res.status(400).send({error: 'Too many tickets were requested. Select no more than 4.'});
    }

    // check if tickets are reserved via gRPC
    try{
        for (const ticket of data.tickets) {
            console.log(ticket);
            try {
                // Prepare request payload (assuming ticketId is in UUID format)
                console.log("inside the try block of gRPC call!");
                const ticketIdBuffer = Buffer.from(ticket.ticketId.replace(/-/g, ""), "hex");
                
                // Make gRPC call
                const ticketResponse = await checkTicketsReservedHandler(client, ticketIdBuffer)
        
                //status 3 is reserved, 2 is bought, 1 is available, 0 is unspecified. Refer to protobuff
                if(ticketResponse && ticketResponse.status === 3 && ticketResponse.userId) {
                    console.log("These tickets have been reserved by you!");
                }else{
                    return res.status(400).send({error:"Something went wrong. Your tickets may have been reserved by someone else or bought, or there are no more available tickets in the sections you have chosen. The ticketId = " + ticketResponse.id.toString("hex")});
                }
            } catch (error) {
                console.error('Error in gRPC call:', error);
                return res.status(500).send({error: 'Error checking the tickets'});
            }
        }
    }catch(error:any){
        console.log("Error in passing in data");
        return res.status(500).send({error: 'Internal Server Error'});
    }

    // Create a PaymentIntent with the order amount and currency
    const paymentIntent = await stripe.paymentIntents.create({
        amount: data.amount,
        currency: "sgd",
        payment_method_types: ["card"],
        // automatic_payment_methods: {
        //     enabled: true,
        // },
        metadata:{
            tickets:JSON.stringify(data.tickets),
            amount:data.amount,
            userId:data.userId,
            sections:JSON.stringify(data.sections),
            email:data.email
        }
    });

    res.send({
        clientSecret: paymentIntent.client_secret,
        paymentId: paymentIntent.id
    });
}

//when you click "Pay Now" after filling your credit card details
export const createPaymentStatusHandler = (prisma: PrismaClient, stripe: Stripe, client: TicketServiceClient) => async (req : Request, res: Response) => {
    console.log('Creating payment status');
	try {
		console.log('In server payment_id:', req.body.payment_id);
		const data = req.body;
		console.log(data);

		// Create a new payment entry
        
		const payment: any = await prisma.payments.create({
			data: {
				payment_id: data.payment_id,
				payment_status: data.payment_status,
				datetime: new Date().getTime(),
				payment_user: data.payment_user,
				payment_ticket: data.payment_ticket,
				payment_amount: data.payment_amount,
				payment_provider: data.payment_provider,
			},
		});
		console.log('Ran prisma.payments.create');

		// Convert BigInt to String in the payment object
		const paymentSerializable = paymentSerializer(payment);

		return res.status(201).json({ payment: paymentSerializable });
	} catch (error) {
		console.error('Error creating payment: ', error);
		return res.status(500).json({ error: 'An error occurred while creating the payment' });
	}
}

module.exports = {
    createPaymentIntentHandler,
    createPaymentStatusHandler,
}