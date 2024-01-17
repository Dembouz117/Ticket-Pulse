import Confetti from "@/components/Confetti";
import Timer from "@/components/Timer";
import TopNav from "@/components/TopNav";
import Footer from "@/components/Footer";
import { authModalAtom } from "@/state/globals";
import { getAuthApiUrl, getQueueApiUrl } from "@/utilities/common";
import axios from "axios";
import { get } from "http";
import { useAtom } from "jotai";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const WaitingRoom = () => {
    const [showAuthModal, setShowAuthModal] = useAtom(authModalAtom);
    const [showConfetti, setShowConfetti] = useState(false);
    const router = useRouter();

    useEffect(() => {
        const checkLoginStatus = async () => {
            try {
                const response = await axios.get(
                    `${getAuthApiUrl()}/checkLogin`,
                    { withCredentials: true }
                );
                console.log("auth checkLogin response:");
                console.log(response);
                if (response.status !== 200) {
                    router.push("/");
                } else {
                    const addToWaitingRoomResponse = await axios.post(
                        `${getQueueApiUrl()}/join`,
                        {},
                        { withCredentials: true }
                    );
                    console.log(addToWaitingRoomResponse);
                }
            } catch (error) {
                console.error("Error checking login status:", error);
                router.push("/");
            }
        };

        checkLoginStatus();
    }, []);

    const handleEndWaitingRoom = async () => {
        try {
            const res = await axios.get(`${getQueueApiUrl()}/randomise`);
            setShowConfetti(true);
            setTimeout(() => {
                setShowConfetti(false);
                router.push("/queue");
            }, 2000);
        } catch (error) {
            console.error("Error ending waiting room:", error);
        }
    };

    return (
      <div className="flex flex-col h-screen bg-background text-white">
          <TopNav setShowAuthModal={setShowAuthModal} />
          <div className="flex flex-col p-56 items-center justify-center flex-grow">
              <div className="w-2/3" data-cy="timer">
                  <Timer
                      startDateTime={1696608329}
                      endDateTime={1728230726}
                  />
              </div>
              <div className="w-2/3 mt-6 mb-6 text-center text-2xl font-semibold" data-cy="waiting-room-message">
                  You are currently in the waiting room. When the sale begins,
                  you will be automatically moved to the queue.
              </div>
              <button
                  onClick={handleEndWaitingRoom}
                  className="bg-button hover:bg-button-hover px-6 py-3 rounded-lg"
                  data-cy="end-waiting-room-button"
              >
                  Simulate End Waiting Room
              </button>
          </div>
          {showConfetti && <Confetti data-cy="confetti" />}
          <Footer />
      </div>
  );
};

export default WaitingRoom;