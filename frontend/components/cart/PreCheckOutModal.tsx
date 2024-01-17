import { useState, useEffect } from "react";
import { useAtom } from "jotai";

import { ticketResponseAtom, ticketAvailabilityAtom, cartEntryAtom, showStripeCheckoutAtom, sectionSelectionAtom } from "@/store";

//additional components and styles
import { Modal, ModalCloseButton, ModalContent, ModalOverlay } from "@chakra-ui/react";
import InteractiveButton from "@/components/utilities/InteractiveButton";
import CartEntry from "./CartEntry";
import { getTicketingApiUrl } from "@/utilities/common";
import { useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";

import axios from "axios";


interface preCheckOutModalProps{
    open: boolean,
    onClose: () => void
}


const PreCheckOutModal = ({open, onClose} : preCheckOutModalProps) => {

    const toast = useToast();
    const router = useRouter();

    const [showStripeCheckout, setShowStripeCheckout] = useAtom(showStripeCheckoutAtom);
    const [cartEntry, setCartEntry] = useAtom(cartEntryAtom);

    //for actual ticket response
    const [tickets, setTickets] = useAtom(ticketResponseAtom);
    const [ticketsAvailability, setTicketsAvailability] = useAtom(ticketAvailabilityAtom);
    
    const [validButton, setValidButton] = useState<boolean>(false);
    const [totalAmount, setTotalAmount] = useState(cartEntry.reduce((total:number, ticket:any) => total + ticket.quantity*ticket.price, 0));

    const handleClose = ():void => {
        onClose();
    }

    const validQuantity = (quantity:number): boolean => {
      return quantity > 0 && quantity <= 4
    }

    useEffect(()=>{
      const newTotalQuantity = cartEntry.reduce((total:number, ticket:any) => total + ticket.quantity, 0);
      const newTotalAmount = cartEntry.reduce((total:number, ticket:any) => total + ticket.quantity*ticket.price, 0);
      setTotalAmount(newTotalAmount);
      const validQuant = validQuantity(newTotalQuantity);
      if (!validQuant && newTotalQuantity > 0){
        toast({
          title: "Invalid Seat Selection Quantity",
          description:
              "You can select no more than 4 seats",
          status: "error",
          duration: 5000,
          isClosable: true,
          position: "top",
          onCloseComplete: onClose
      });
      }
      setValidButton(validQuant);
    },[cartEntry]);


    const confirmHandler = async () => {
      setShowStripeCheckout({visible:true, amount: totalAmount});
      console.log(showStripeCheckout);
      
      const sectionsList = cartEntry.filter(el=>el.quantity>0&&el.sectionId.length>0).map(el=>{return{sectionId:el.sectionId, quantity:el.quantity}});
      const checkoutParams : any = {
        seats: sectionsList
      }
  
      try {
        //this endpoint checks for available tickets and immediately reserves them
        const url = `${getTicketingApiUrl()}/api/v1/sessions/available`
        axios.post(url, checkoutParams,{
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
          },
        })
        .then((response : any) => {
          return response.data;
        })
        .then((data : any)=> {
          setTickets(data.tickets);
          console.log("TICKETSSSS AVAILABILITY:", data.tickets.length);
          console.log(data.tickets);
          setTicketsAvailability("AVAILABLE"); 
        })
        .catch((error : any) => {
          setTicketsAvailability("UNAVAILABLE");
          console.log("Seat availability check error: ",error.message);
        });
  
  
    }catch (error:any){
      console.log(error.message);
    }
  
  }

    return(
        <Modal
        isOpen={open}
        onClose={handleClose}
        closeOnOverlayClick={true}
        size={"lg"}
        colorScheme="white"
    >
        <ModalOverlay />
        <ModalContent textColor={"white"} className = "py-12 px-4 h-auto bg-white"   position="absolute" right={5}>
            <ModalCloseButton className="text-button ease-in-out transition-all duration-200 hover:border-2 hover:border-black hover:-translate-y-1 hover:scale-125 hover:shadow" />
            <div className = "flex-col space-y-6 h-full justify-center items-center">
                <h2 className = "text-3xl text-black font-bold">My Cart</h2>
                {cartEntry.map(entry=><div>{entry.quantity>0&&<CartEntry section={entry.sectionName} closeSlider={() => true}/>}</div>)}
                <div className="w-full flex justify-center items-center">
                  <h3 className="text-black font-bold underline">Total Amount:&nbsp;&nbsp;${cartEntry.reduce((total, ticket) => total + ticket.quantity*ticket.price, 0)}</h3>
                </div>
                <div className="w-full flex justify-center">
                  <InteractiveButton onClick={confirmHandler} valid={validButton}>Confirm Seats</InteractiveButton>
                </div>

            </div>
        </ModalContent>
    </Modal>
    )
}

export default PreCheckOutModal;