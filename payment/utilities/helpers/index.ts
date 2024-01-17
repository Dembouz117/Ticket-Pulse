import { payments } from "@prisma/client"

export const paymentSerializer = (payment: payments) => {
    const paymentSerializable: any = { ...payment };
		
    for (const prop in paymentSerializable) {
        if (typeof paymentSerializable[prop] === 'bigint') {
            paymentSerializable[prop] = paymentSerializable[prop].toString();
        }
    }
}