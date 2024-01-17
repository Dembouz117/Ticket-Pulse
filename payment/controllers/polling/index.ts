import { Request, Response } from "express";
import { PrismaClient }from "@prisma/client";
import { paymentSerializer } from "../../utilities/helpers";

export const pollingHandler = (prisma: PrismaClient) => async (req: Request, res: Response) =>{
    const { paymentId } = req.params;
	console.log('paymentId = ' + paymentId);

    try{
        // Query the payment status from db
        const payment: any = await prisma.payments.findFirst({
            where: { payment_id: paymentId as string },
        });
        console.log(payment);

        if (!payment) {
            return res.status(404).json({ error: 'Payment record not found!' });
        }

        // Convert BigInt to String in the payment object. For some reason my utility function doesn't work here
        const paymentSerializable: any = { ...payment };
        let status: string = '';
        for (const prop in paymentSerializable) {
            if (typeof paymentSerializable[prop] === 'bigint') {
                paymentSerializable[prop] = paymentSerializable[prop].toString();
            }
        }
        status = paymentSerializable.payment_status;
        return res.json({ paymentObj: paymentSerializable, status: status });
    }catch(error: any){
        return res.status(500).json({error: "Something went wrong with checking your payment records!"});
    }

} 

module.exports = {
    pollingHandler
}