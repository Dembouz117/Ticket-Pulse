import { TicketServiceClient } from "../../proto/entpb/TicketService"; 
import { SectionServiceClient } from "../../proto/entpb/SectionService";
// import { SurfaceCall } from "@grpc/grpc-js/build/src/call";


export const checkTicketsReservedHandler = async (client: TicketServiceClient, ticketId: Buffer) => {
    // Make gRPC call
    const ticketResponse : any = await new Promise((resolve, reject) => {
        client.Get({
            id: ticketId,
            view: "VIEW_UNSPECIFIED"
        }, (err, response) => {
            if (err) {
                console.log(err.message);
                return reject(err)};
            resolve(response);
        });
        });
        console.log("after grpc call!");
        console.log(ticketResponse);

    return ticketResponse;
}

export const 
updateTicketHandler = async (ticketClient: TicketServiceClient, sectionClient: SectionServiceClient, sectionBuffer: Buffer, userId: string, ticketBuffers:Buffer[]) => {
    ticketBuffers.forEach((ticketBuffer) => {
        try{
            const ticketResponse = ticketClient.Get(
                {
                    id: ticketBuffer,
                    view: 0
                },
                (err, result) => {
                    console.log("IN TICKET GET BOI");
                    if (err) {
                        console.log("There's an error with the gRPC Ticket Get call!");
                        console.error(err.message);
                        return;
                    }
    
                  
                    //make sure to use spread. Be careful if passing your own arguments via another variable as JS objects have default properties
                    //e.g it may convert buffer to string, and long to number etc. Specify the keys properly
                    ticketClient.Update(
                        {
                            ticket: {...result, status: 2}
                        },
                        //err may not suggest network layer error, but also includes app level errors so just kim to check first while debugging
                        (err , result) => {
                            console.log("IN TICKET CLIENT UPDATE BOI");
                            if (err) {
                                console.log("There's an error with the gRPC Ticket call!");
                                console.error(err.message);
                                return;
                            }
                            if (result && result.status == 3) {
                                console.log('The ticket has been reserved by someone else!');
                                return;
                            }
                            if (result && result.status == 2) {
                                console.log('The ticket has already been purchased from someone else!');
                                return;
                            } else {
                                console.log('The ticket is available for purchase!');
                                return;
                            }
                        }
                    )
                }
           
        
            );
        }catch(err){
            console.log(err);
        }   
    });       
}

