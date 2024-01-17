import { userAtom } from "@/state/globals";
import { Button } from "@chakra-ui/react";
import { useAtom } from "jotai";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useRef, useState } from "react";
import Logo from "@/components/TicketPulseLogo";
import axios from "axios";
import { getAuthApiUrl } from "@/utilities/common";

import { useRouter } from "next/router";

import CheckOutButton from "@/components/cart/CheckOutButton";

interface TopNavProps {
  setShowAuthModal: (open: boolean) => void;
}

const TopNav = ({ setShowAuthModal }: TopNavProps) => {
  const [user, setUser] = useAtom(userAtom);
  const router = useRouter();
  const isSeatsPage = router.pathname.includes('/seats');



  useEffect(() => {

    const fetchUser = async () => {
      try {
       
        const response = await axios.get(`${getAuthApiUrl()}/user`, { withCredentials: true });
        console.log("auth response:");
        console.log(response.data.user);
        setUser({email: response.data.user.email, userId: response.data.user.id});
        

        //setUser({userId: response.data.user.id});
      } catch (error) {
        console.error("Axios error:", error);
        console.log("smth went wrong in the topnav");
      }
    };

    fetchUser();
  }, []);

  return (
    <div className="w-full">
      <div className="bg-neutral-900 w-full relative">
        <nav className="w-[90%] px-1 md:px-0 mx-auto py-4 items-center flex justify-between">
          <div className="text-white w-10 h-10">
            <Link href="/">
              <Logo />
            </Link>
          </div>
          <div className="flex flex-row justify-between w-auto items-center"> 
            {isSeatsPage && <CheckOutButton/>}
            {user !== null && (
              <div
                className="font-bold text-white flex flex-row justify-center align-middle"
              >
                <Link href="/profile">
                  <Image
                    className="rounded-full ml-3 border-1 border-white hover:border-2"
                    src={"/images/user_icon.png"}
                    alt={"userIcon"}
                    width={40}
                    height={30}
                    draggable={false}
                  />
                </Link>
              </div>
            )}

            {user === null && (
              <Button
                className="bg-blue-500 font-medium text-white rounded-xl md:py-2 md:px-4 p-2 hover:bg-blue-600"
                onClick={() => setShowAuthModal(true)}
                data-cy="auth-open"
              >
                Sign in
              </Button>
            )}
          </div>

        </nav>
      </div>
    </div>
  );
};

export default TopNav;
