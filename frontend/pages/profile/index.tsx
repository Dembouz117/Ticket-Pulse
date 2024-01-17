// pages/profile.tsx
import React, { useState, useEffect } from "react";
import TopNav from "@/components/TopNav";
import Footer from "@/components/Footer";
import { userAtom } from "@/state/globals";
import Link from "next/link";

import { Skeleton, SkeletonCircle, SkeletonText } from '@chakra-ui/react'

import axios from "axios";
import { useAtom } from "jotai";
import { getAuthApiUrl, getTicketingApiUrl } from "@/utilities/common";
import { useRouter } from "next/router";
import ProfileTicket from "@/components/profile/ProfileTicket";
import { ticketSchemaInterface, concertInterface } from "@/types";
import AuthModal from "@/components/landing/auth/AuthModal";


// COMMENTED CODE IS FOR GROUP VIEW

const Profile: React.FC = () => {
  const router = useRouter();
  const [showAuthModal, setShowAuthModal] = useState<boolean>(false);
  const [user, setUser] = useAtom(userAtom);

  const [fetchedTickets, setFetchedTickets] = useState<ticketSchemaInterface[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  const headers = {
    Authorization: "Access-Control-Allow-Origin",
  };

  const tickets = useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        const response = await axios.get(
          `${getTicketingApiUrl()}/api/v1/tickets/user`,
          {
            withCredentials: true,
          }
        );
        setLoading(false);
        setFetchedTickets(response.data);
      } catch (error) {
        setLoading(true);
        console.error("Axios error:", error);
      }
    };

    fetchData();
  }, [showAuthModal, user]);

  // const processData = (tickets: any) => {
  //   const groupedData: any = {};

  //   if (!tickets || tickets.length === 0) {
  //     return [];
  //   }
  //   tickets.forEach((ticket: any) => {
  //     const concert =
  //       ticket.edges.withinSection.edges.atConcertSession.edges.ofConcert[0];
  //     const category = ticket.edges.withinSection.category;

  //     const key = `${concert.id}-${category}`;
  //     if (!groupedData[key]) {
  //       groupedData[key] = {
  //         artist: concert.artist,
  //         title: concert.title,
  //         category: category,
  //         count: 0,
  //       };
  //     }
  //     groupedData[key].count += 1;
  //   });

  //   return Object.values(groupedData);
  // };

  // const processedTickets = processData(fetchedTickets);

  const [userDetails, setUserDetails] = useState<any>();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(`${getAuthApiUrl()}/user`, {
          withCredentials: true,
        });
        setUserDetails(response.data);
      } catch (error) {
        console.error("Fetch user details error:", error);
      }
    };

    fetchData();
  }, [showAuthModal]);

  // const [user, setUser] = useAtom(userAtom);

  const handleLogout = async () => {
    try {
      await axios.post(
        `${getAuthApiUrl()}/logout`,
        {},
        { withCredentials: true }
      );
      setUser(null);
      router.push("/");
    } catch (error) {
      console.error("Logout error:", error);
    }
  };

  return (
    <div className="h-full bg-background">
      <TopNav setShowAuthModal={setShowAuthModal} />
      {showAuthModal && <AuthModal open={showAuthModal} onClose={() => setShowAuthModal(false)} />}
      <div className="min-h-screen">
        <div className="max-w-4xl mx-auto bg-white p-6 m-6 rounded-lg shadow-md">
          <div className="flex flex-row justify-between items-center mb-4">
            <h1 className="text-2xl font-bold">Profile</h1>
            <div
              className="bg-blue-500 font-medium text-white rounded-xl md:py-2 md:px-4 p-2 hover:bg-blue-600 cursor-pointer"
              onClick={handleLogout}
            >
              Sign Out
            </div>
          </div>
          <section className="mb-8">

            <h2 className="text-xl font-semibold mb-2">Personal Information</h2>
            <p>
            <Skeleton isLoaded={!loading}>
              <strong>Name:</strong> {userDetails?.user?.name || ""}
            </Skeleton>
            </p>
            <p className="mt-2">
              <Skeleton isLoaded={!loading}>
                <strong>Email:</strong> {userDetails?.user?.email || ""}
              </Skeleton> 
            </p>
            <p className="mt-2">
              <Skeleton isLoaded={!loading}>
                <strong>Mobile:</strong> {userDetails?.user?.phone || ""}
              </Skeleton>
            </p>
          </section>

          <section>
            <h2 className="text-xl font-semibold mb-4">Purchased Tickets</h2>
            <div className="max-h-[30rem] overflow-y-auto">
              <ul>
                <Skeleton isLoaded={!loading}>
                  {fetchedTickets.map((ticket, index) => {
                    const concert: concertInterface =
                      ticket.edges.withinSection.edges.atConcertSession.edges
                        .ofConcert[0];
                    const category = ticket.edges.withinSection.category;
                    console.log(ticket);
                    const boughtTicket = ticket.status==="BOUGHT";


                    return (
                      // <div key={index} className="border p-4 rounded">
                      //   <h2 className="text-xl font-semibold">
                      //     {concert.artist}'s {concert.title}
                      //   </h2>
                      //   <p>Category: {category}</p>
                      //   <p>Seat Number: {ticket.seatNumber}</p>
                      // </div>
                      <>
                      {boughtTicket?<ProfileTicket concert={concert} category={category} ticket={ticket}/>:<></>}
                      </>
                    );
                  })}
                </Skeleton>
                {/* {processedTickets.map((ticket: any, index) => (
                <div key={index} className="border p-4 rounded">
                  <h2 className="text-xl font-semibold">{ticket.title}</h2>
                  <p className="text-gray-600">Artist: {ticket.artist}</p>
                  <p>Category: {ticket.category}</p>
                  <p>Number of Tickets: {ticket.count}</p>
                </div>
              ))} */}

              </ul>
            </div>
            
          </section>
        </div>
      </div>

      <Footer />
    </div>
  );
};

export default Profile;
