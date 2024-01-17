import { Stripe, loadStripe } from '@stripe/stripe-js';

// require('dotenv').config({ path: `../.env.local` });

let stripePromise: Promise<Stripe | null>;
const getStripe = () => {
	if (!stripePromise) {
		console.log("HERE IN STRIPE!");
		const publicKey = process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY!.trim();
		console.log(process.env.NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY!);
		stripePromise = loadStripe(publicKey);
	}
	return stripePromise;
};

export default getStripe;
