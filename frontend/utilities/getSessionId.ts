import { getTicketingApiUrl } from "./common";
import axios from "axios";

export const getSessionId = async (artist:string): Promise<string> => {
    try {
        //gets concerts, then sessions
        const concertSessionId = await axios.get(`${getTicketingApiUrl()}/api/v1/concerts/artist/${artist}`)
        .then(res=>{
            console.log("HERE HERE!");
            console.log(res.data.concerts[0]);
            return res.data.concerts[0];
        })
        .then(
            async concert => {
                console.log("HERE IN ASYNC!");
                console.log(concert);
                console.log(concert.id);
                const sessionResponse = await axios.get(`${getTicketingApiUrl()}/api/v1/concerts/${concert.id}/sessions`)
                .then(res=>{
                    console.log("HERE IS:");
                    console.log(res.data.sessions);
                    return res.data.sessions[0];
                });
                
    

                return sessionResponse.id;
            }
        );

        return concertSessionId;
    } catch (error) {
        console.error("An error occurred sometime during concert fetch:", error);
        throw error; // This will reject the promise
    }
}

