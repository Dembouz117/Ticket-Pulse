import { FaShoppingCart } from "react-icons/fa";

import { useState, useEffect } from "react";
import { useAtom } from "jotai";
import { cartEntryAtom, showPreCheckoutAtom } from "@/store";

import { cartEntryType } from "@/types";

import PreCheckoutModal from "@/components/cart/PreCheckOutModal";


export const CheckOutButton = () => {
  const [ selectedSections, setSelectedSections ] = useAtom(cartEntryAtom);
  const [ showPreCheckOut, setShowPreCheckOut ] = useAtom(showPreCheckoutAtom);
  const [ totalQuantityAndPrice, setTotalQuantityAndPrice ] = useState({totalPrice:0, totalQuantity:0});

  useEffect(()=>{
    const totalPrice = selectedSections.reduce((total:number, currVal:cartEntryType) => currVal.quantity*currVal.price + total, 0);
    const totalQuantity = selectedSections.reduce((total:number, currVal:cartEntryType) => currVal.quantity+total, 0);
    setTotalQuantityAndPrice({totalPrice:totalPrice, totalQuantity:totalQuantity});
  },[selectedSections]);

  const closeModalHandler = () => {
    setShowPreCheckOut(false);
  }

  const cartClickHandler = () => {
    console.log("I am clicked in cartClickHandler");
    setShowPreCheckOut(true);
  }

  return (
    <button className="mx-6 border-2 border-solid border-white px-4 py-2 relative rounded" onClick={cartClickHandler}>
        <FaShoppingCart className="w-6 h-6"/>
        {totalQuantityAndPrice.totalQuantity > 0 && 
        <div className="absolute top-[-10px] right-[-10px] bg-button text-white rounded-full w-6 h-6 flex justify-center items-center">
            {totalQuantityAndPrice.totalQuantity}
        </div>
        }
        <div className="absolute top-6 left-25">
          <PreCheckoutModal open={showPreCheckOut} onClose={closeModalHandler}/>
        </div>
    </button>
  )
}

export default CheckOutButton;