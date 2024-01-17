
import { useAtomValue } from "jotai";
import { ticketResponseAtom, ticketAvailabilityAtom } from "@/store/index";


import TicketResponse from "@/components/payment/TicketResponse";

interface LoadingSeatsProps{
  onEndLoading: (accept:boolean) => void,
}

const LoadingSeats = ({onEndLoading}:LoadingSeatsProps) => {


  //real tickets and their state. Toggled in cart
  const tickets = useAtomValue(ticketResponseAtom);
  const ticketAvailability = useAtomValue(ticketAvailabilityAtom);

    
  return (
    <div className="flex-col space-y-12">
    {ticketAvailability === "" && <TicketResponse accept = {onEndLoading} available = {""} tickets = {tickets}/>}
    {ticketAvailability === "AVAILABLE" && <TicketResponse accept = {onEndLoading} available = {"AVAILABLE"} tickets = {tickets}/>}
    {(tickets.length <= 0 && ticketAvailability === "UNAVAILABLE") && <TicketResponse accept = {onEndLoading} available = {"UNAVAILABLE"} tickets = {tickets}/>}
    </div>
  )
}

export default LoadingSeats;