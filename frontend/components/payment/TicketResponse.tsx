import { Spinner } from '@chakra-ui/react';
import Logo from "@/components/TicketPulseLogo.svg";
import { ticketResponseInterface } from "@/types/index"

interface TicketResponseProps{
  available: string,
  tickets: ticketResponseInterface[],
  accept: (accept: boolean) => void
}

export const TicketResponse = ({ available, tickets, accept } : TicketResponseProps) => {

  const successfulResponse = 
    <div className="flex-col space-y-2">
      {tickets && tickets.map((el: ticketResponseInterface, idx: number) => {
        console.log("inside ticketDisplay");
        console.log(el);
        return (
          <div className="border p-4 rounded" key={idx}>
            <div>Id: {el.ticketId}</div>
            <div>Seat Number: {el.seatNumber}</div>
            <div>Section: {el.sectionCategory}</div>
          </div>
        );
      })}
      <div className="flex justify-between">
        <div>Do you wish to accept these seats?</div>
        <button className="bg-button px-3 py-2 rounded" onClick={() => accept(true)}>Yes</button>
        <button className="bg-button px-3 py-2 rounded" onClick={() => accept(false)}>No</button>
      </div>
    </div>;

  const failedResponse =     
    <div className = "w-full flex-col space-y-6">
      <h3>Sorry, there are no available tickets that fit your requirements at the moment! Please come back in a moment.</h3>
      <div className = "w-full flex justify-center items-center">
          <Logo className = "w-14 h-14"/>
      </div>
    </div>;

  const loadingResponse = 
    <div>
      <div className="text-3xl tex-bold">Please wait while we find available seats for you...</div>
      <div className = "w-full flex justify-center">
        <Spinner/>
      </div>
    </div>

  return (
    <div>
      {available === "AVAILABLE" && successfulResponse}
      {available === "UNAVAILABLE" && failedResponse}
      {available === "" && loadingResponse}
    </div>

  )
}

export default TicketResponse;
