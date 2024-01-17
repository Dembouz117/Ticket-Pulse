//This is the main component that slides out and stuff

import styles from "@/styles/Cart.module.css";
import { MdClear } from "react-icons/md";

import { cartEntryType } from "@/types/index";

import {
  showCartAtom,
  cartEntryAtom,
  showStripeCheckoutAtom,
  currSectionAtom,
  ticketResponseAtom,
  ticketAvailabilityAtom,
} from "@/store/index";

import { useAtom } from "jotai";
import { useEffect, useState } from "react";

import CartAmountButton from "@/components/cart/CartAmountButton";
import CartEntry from "@/components/cart/CartEntry";

import axios from "axios";
import { getTicketingApiUrl } from "@/utilities/common";

interface cartProp {
  concertId: string;
}

const Cart = ({ concertId }: cartProp) => {
  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [sectionQuantityList, setSectionQuantityList] = useAtom(cartEntryAtom);
  const [showStripeCheckout, setShowStripeCheckout] = useAtom(
    showStripeCheckoutAtom
  );
  const [currSection, setCurrSection] = useAtom(currSectionAtom);

  const [sectionAmount, setSectionAmount] = useState<number>(0);
  const [windowHeight, setWindowHeight] = useState<number>(0);

  //for actual ticket response
  const [tickets, setTickets] = useAtom(ticketResponseAtom);
  const [ticketsAvailability, setTicketsAvailability] = useAtom(
    ticketAvailabilityAtom
  );

  useEffect(() => {
    if (!showCart) {
      document.body.classList.add("overflow-x-hidden");
    } else {
      document.body.classList.remove("overflow-x-hidden");
    }
  }, [showCart]);

  useEffect(() => {
    setSectionAmount(
      sectionQuantityList.reduce(
        (total: number, currVal: cartEntryType) =>
          currVal.quantity * currVal.price + total,
        0
      )
    );
  }, [sectionQuantityList]);

  useEffect(() => {
    const handleScroll = () => {
      const middleOfScreen = 0;
      setWindowHeight(window.scrollY + middleOfScreen);
    };

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  const confirmButtonHandler = async () => {
    console.log("Inside confirm button handler");
    setShowStripeCheckout({ visible: true, amount: sectionAmount });

    const checkoutParams: any = {
      seats: [
        // can just uncomment these seats with invalid IDs to test error handling
        // {
        //   sectionId:"bba46785-6e27-4f53-9664-01e4abe76ee3",
        //   quantity:2
        // },
        {
          sectionId: "dc493beb-c9e4-4851-8feb-68a9232721b3",
          quantity: 3,
        },
      ],
    };

    try {
      //this endpoint checks for available tickets and immediately reserves them
      const url = `${getTicketingApiUrl}/api/v1/sessions/available`;
      axios
        .post(url, checkoutParams, {
          withCredentials: true,
          headers: {
            "Content-Type": "application/json",
          },
        })
        .then((response: any) => {
          return response.data;
        })
        .then((data: any) => {
          setTickets(data.tickets);
          console.log("TICKETSSSS AVAILABILITY:", data.tickets.length);
          setTicketsAvailability("AVAILABLE");
        })
        .catch((error: any) => {
          setTicketsAvailability("UNAVAILABLE");
          console.log("Seat availability check error: ", error.message);
        });
    } catch (error: any) {
      console.log(error.message);
    }
  };

  return (
    <div
      className={`h-screen absolute top-1/4 right-0 w-1/4 z-20 ${
        showCart ? `${styles["slide-in"]}` : `${styles["slide-out"]}`
      }`}
      style={{ top: `${windowHeight}px` }}
    >
      <button
        className="inline-block px-2 py-2 bg-slate-200 rounded-md absolute top-4 right-4 ease-in-out transition-all duration-200 hover:border-2 hover:border-black hover:-translate-y-2 hover:scale-125 hover:shadow"
        onClick={() => {
          setShowCart(false);
        }}
      >
        <MdClear className="text-red-200 hover:inherit" />
      </button>

      <div className="h-full w-full bg-white flex items-center justify-center">
        <div className="flex-col w-full space-y-3">
          <CartEntry section={"Cat 3"} closeSlider={() => true}/>
          <CartEntry section={"Cat 2"} closeSlider={() => true}/>
          <CartEntry section={"Cat 1"} closeSlider={() => true}/>

          <div className="flex justify-center w-full">
            <div className="flex-col space-y-3">
              <div className="text-black mt-3 mb-1 underline font-semibold">
                Total Amount:&nbsp;&nbsp;&nbsp;${sectionAmount}
              </div>
              <button
                className="bg-button rounded-md p-2 m-2 text-white"
                onClick={confirmButtonHandler}
              >
                Confirm
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Cart;
