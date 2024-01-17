import SeatSelection from "@/pages/seats/[artist]";
import { sectionAlgo, TicketDataInterface } from "@/types/index";

//need to consider how many seats in a row? Then modulo
//points for improvement:

/*
a) people typically prefer front seats -> if cannot find, direct to another section
b) mb use probabilistic model to determine which seats are more likely to be bought
c) P-Queue for front seats? (queue for each section)
d) instead of returning null, return the next best seat
e) if i buy 2 section a, 2 section b (checkout successful if both conditions are met)
*/


// For initialising seat data
export const initializeSeats = (sections:sectionAlgo[]) => {
    const seats:TicketDataInterface = {};
    sections.forEach((section : sectionAlgo) => {
      seats[`${section.section}-${section.category}`] = Array.from({ length: section.seats }, (_, index) => ({
        id: `${section.section}-${section.category}-${index + 1}`,
        status: "AVAILABLE",
      }));
    });
    return seats;
};

//main algorithm
export const findContiguousSeats = async (numSeats:number, section: string, category: string, seatData: TicketDataInterface) => {

    const lock = new Promise<void>((resolve) => {
        // Simulate locking for concurrency control
        setTimeout(() => resolve(), 0);
      });
    
      await lock;

    const keyId = `${section}-${category}`;

    const seatsInSection = seatData[keyId];
  
    console.log(seatsInSection);
    let consecutiveSeats = 0;
    let startSeatIndex = 0;

    for (let seatIndex = 0; seatIndex < seatsInSection.length; seatIndex++) {
        const seat = seatsInSection[seatIndex];
        if (seat.status === "AVAILABLE") {
        consecutiveSeats++;

        if (consecutiveSeats === numSeats) {
            // Allocate seats
            for (
            let i = startSeatIndex;
            i < startSeatIndex + numSeats;
            i++
            ) {
            seatsInSection[i].status = "RESERVED";
            }
            return seatsInSection.slice(
            startSeatIndex,
            startSeatIndex + numSeats
            );
            //return false
        }
        } else {
        consecutiveSeats = 0;
        startSeatIndex = seatIndex + 1;
        }
    }
    // No consecutive seats found
    return null;
    //return true;
};


// to handle results from findContiguousSeats
export const allocateSeatsConcurrently = async (seatAmount:number, section:string, category:string, seatData : TicketDataInterface) => {
    try {
      const allocatedSeats = await findContiguousSeats(seatAmount,"A", "CAT3", seatData);
      if (allocatedSeats) {
        console.log("Allocated Seats:", allocatedSeats);
      } else {
        console.log("No consecutive seats available.");
      }
    } catch (error: any) {
      console.error("Error:", error.message);
    }
  };