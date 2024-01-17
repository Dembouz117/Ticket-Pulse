import { useEffect } from "react";
import { useAtom } from "jotai";
//components
// import Section from "@/components/algorithm/Section";
//algorithm function
import { allocateSeatsConcurrently } from "@/utilities/algorithm";
//mock data
import { sections } from "@/data/sections"
//import { ticketsAtom } from "@/store/index";

export const seatAlgorithmPage = () => {

  // const [seatData, setSeatData] = useAtom(ticketsAtom)
  
  // //for testing purposes
  // useEffect (()=>{
  //   allocateSeatsConcurrently(1,"A","CAT3", seatData);
  //   allocateSeatsConcurrently(2,"A","CAT3", seatData);
  //   allocateSeatsConcurrently(3,"A","CAT3",seatData);
  //   },[]);
  // return (
  //   <div>
  //       {/* Map sections, then map seats */}
  //     {sections.map((section) => (
  //       <Section
  //         key={section.id}
  //         section={section}
  //         onSelectSeat={() => console.log("Seat selected")}
  //         seat={seatData[`${section.section}-${section.category}`]}
  //       />
  //     ))}
  //   </div>
  // )
  return (<div></div>)
}

export default seatAlgorithmPage;