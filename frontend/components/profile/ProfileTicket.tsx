import { concertInterface, ticketSchemaInterface } from "@/types/index"; 
import styles from "@/styles/tickets/ProfileTicket.module.css";
import Logo from "@/components/TicketPulseLogo.svg";
import { parseDate } from "@/utilities/parseDate";
import Image from "next/image";

interface ProfileTicketInterface{
    concert: concertInterface,
    category: string,
    ticket: ticketSchemaInterface
}

const ProfileTicket = ({concert, category, ticket}: ProfileTicketInterface) => {
  const formattedDate = parseDate(ticket.edges.withinSection.edges.atConcertSession.sessionDateTime);
  return (
    <div key={`${concert.title}-${ticket.seatNumber}`} className="group perspective w-full h-36">
        <div className="relative preserve-3d group-hover:rotate-x-180 w-full h-full duration-500 flex-col space-y-4">
          
            <div className = "p-4 rounded-lg absolute shadow-md bg-slate-200 absolute backface-hidden w-full h-full">
                <div className = {`absolute bg-white right-[-1px] top-12 ${styles["ticket-circle-right"]}`}></div>
                <div className = {`absolute bg-white left-[-2px] top-12 ${styles["ticket-circle-left"]}`}></div>
                <div className="ml-8 flex space-x-4">
                    <img src = {concert.imageUrl} className = "max-w-[10.7rem] max-h-[6rem]"/>
                    <div>
                        <h2 className="text-xl font-semibold">
                            {concert.artist}'s {concert.title}
                        </h2>
                        <p>Category: {category}</p>
                        <p>Seat Number: {ticket.seatNumber}</p>
                    </div>
                </div>
            </div>
  
            {/* front face content */}
            <div className="absolute backface-hidden w-full h-full rotate-x-180 border p-4 rounded-lg relative shadow-sm bg-slate-200" >
                <div className = {`absolute bg-white right-[-1px] top-12 ${styles["ticket-circle-right"]}`}></div>
                <div className = {`absolute bg-white left-[-2px] top-12 ${styles["ticket-circle-left"]}`}></div>
                {/* <div className = "absolute h-full w-full flex justify-center items-center inset-0">
                    <img src = {concert.imageUrl} className = "h-full w-11/12 object-conver object-top opacity-5"/>
                </div> */}
                <div className="relative h-full w-full flex px-8">
                    <div className="relative flex-col h-full">
                        <h2 className="text-2xl font-medium">{ticket.edges.withinSection.category}</h2>
                        <h2 className="text-2xl">${ticket.edges.withinSection.price}</h2>
                        <h4 className="text-sm"><em>Inclusive of booking fee</em></h4>
                        <Logo className = "w-8 h-6 mt-2 opacity-70"/>
                    </div>
                    <div className="w-3/6 justify-center flex">
                        <div className="relative">
                            <h2 className="text-3xl">{concert.artist}</h2>
                            <h2 className="text-xl font-light">{concert.title}</h2>
                            <h2>Singapore National Stadium<br/></h2>
                            <h2>{formattedDate}</h2>
                            {/* <Logo className="object-cover"/> */}
                        </div>
                    </div>
                    {/* <Image src = {"images/ticketpulseQR.png"} width={96} height={96} alt = "QR CODE"/> */}
            
                    <div className="h-full w-auto lg:ml-12">
                    <img src="https://utfs.io/f/b4829693-e96a-4193-a85e-363ea6e09c02-ooe5gz.png" className="w-24 h-24"/>
                    </div>
                    
                </div>
            </div>

        </div>
    </div>
  )
}

export default ProfileTicket