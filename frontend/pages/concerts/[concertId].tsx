import { useRouter } from 'next/router';

import Cart from "@/components/cart/Cart";
import PaymentModal from '@/components/payment/PaymentModal';


import { showCartAtom, showStripeCheckoutAtom } from "@/store/index";
import { useAtom } from "jotai";
import { useEffect } from 'react';

const Concert = () => {
  const router = useRouter();
  const { concertId } = router.query;
  console.log("router.query = " , router.query);
  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [showStripeCheckout, setShowStripeCheckout] = useAtom(showStripeCheckoutAtom);
  
  useEffect(() => {
    if (typeof concertId !== 'string') {
      router.push('/404');
    }
  }, [concertId]);

  return (
    <div className = {`w-screen h-screen bg-red-200 z-10${showCart?"opacity-50":""}`}>

      <button className = "w-auto h-auto bg-slate-200" onClick = {() => {
        setShowCart(!showCart); //should be a jotai state
        setShowStripeCheckout({visible:false, amount:0});
      }}>Toggle cart</button>
      <h1>Concert: {concertId}</h1>
      
      {showStripeCheckout && <PaymentModal open = {showStripeCheckout.visible} onClose = {() => setShowStripeCheckout({visible:false, amount:0})}/>}

      <div className = "inset-0 h-full w-full">
          {showCart && <div className = "inset-0 w-full h-full opacity-50 bg-gray-900 absolute" onClick = {() => {
            setShowCart(false);
            }}/>}
          <Cart concertId = {concertId as string}/>
      </div>
   

    </div>
  )
}

export default Concert;