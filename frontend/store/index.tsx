import { atom } from 'jotai';
import {
	cartEntryType,
	showStripeCheckout,
	cartCoordinateInterface,
	mockTicketsInterface,
	ticketResponseInterface,
	sectionInterface
} from '@/types/index';

import { PaymentIntent } from '@stripe/stripe-js';

/*
------ Cart UI Data ------
*/
export const showCartAtom = atom<boolean>(false);

//modal coordinate data
export const cartCoordinatesAtom = atom<cartCoordinateInterface>({
	x: 0,
	y: 0,
});


/*
------ PreCheckout Modal ------
*/
//will use cart data and its own toggle
export const showPreCheckoutAtom = atom<boolean>(false);



/*
------ Section Amount Data ------
*/

export const cartEntryAtom = atom<cartEntryType[]>([]);

/*
------ Stripe Modal ------
*/
export const showStripeCheckoutAtom = atom<showStripeCheckout>({
    visible: false,
    amount: 0
});

/*
------ Error Modal ------
*/
interface errorModalInterface{
	visible: boolean,
	message: string
}
export const errorModalAtom = atom<errorModalInterface>({
	visible: false,
	message:""
})

/*
------ Stripe Data ------
*/
export const paymentIntentAtom = atom<PaymentIntent | null>(null);


/*
------ Seat/Ticket Response ------
*/
export const ticketResponseAtom = atom<ticketResponseInterface[]>([]);
export const ticketAvailabilityAtom = atom<string>("");


//for selection of the seats
interface sectionSelectionInterface{
	sectionId: string,
	quantity: number
}
const initialSectionSelectionData: sectionSelectionInterface[] = [{
	sectionId:"",
	quantity:0
  }]
export const sectionSelectionAtom = atom<sectionSelectionInterface[]>(initialSectionSelectionData);

//for fetching the sections from backend
export const sectionAtom = atom<sectionInterface[]>([]);

/*
------ Category Selection Data ------
*/
//For the selection of the section and then displaying them onto the preview sidebar
interface currSectionInterface{
	sectionUuid:string,
	price:number,
	name:string
}
const intialCurrSectionData: currSectionInterface = {
	sectionUuid:"",
	price:1000,
	name:""
}
export const currSectionAtom = atom<currSectionInterface>(intialCurrSectionData);