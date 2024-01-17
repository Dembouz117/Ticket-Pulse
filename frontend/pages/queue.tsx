import TopNav from "@/components/TopNav";
import Footer from "@/components/Footer";
import Spinner from "@/components/utilities/Spinner";
import { authModalAtom } from "@/state/globals";
import {
  getAuthApiUrl,
  getQueueApiUrl,
  getQueueWebsocketUrl,
} from "@/utilities/common";
import axios from "axios";
import { useAtom } from "jotai";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const Queue = () => {
  const [showAuthModal, setShowAuthModal] = useAtom(authModalAtom);
  const [position, setPosition] = useState<number | null>(null);
  const router = useRouter();

  useEffect(() => {
    const checkLoginStatus = async () => {
      try {
        const response = await axios.get(`${getAuthApiUrl()}/checkLogin`, {
          withCredentials: true,
        });
        if (response.status !== 200) {
          router.push("/");
        } else {
          //TODO: check if below check fks with positioning
          await axios.post(
            `${getQueueApiUrl()}/join`,
            {},
            { withCredentials: true }
          );
          // establish websocket connection
          const ws = new WebSocket(`${getQueueWebsocketUrl()}`);
          ws.onmessage = (event) => {
            console.log(event);
            const message = event.data;

            // match for position in queue and update state
            const match = message.match(/Position in queue is: (\d+)/);
            if (match && match[1]) {
              setPosition(Number(match[1]));
            }

            // redirect to buy tickets page
            if (message === "Redirect to buy tickets") {
              router.push("/seats/Taylor%20Swift");
            }
          };

          //TODO: string grab the position from the message then useState or smth to display

          ws.onclose = (event) => {
            if (event.wasClean) {
              console.log(
                `[close] Connection closed cleanly, code=${event.code}, reason=${event.reason}`
              );
            } else {
              console.error("[close] Connection died");
            }
          };

          ws.onerror = (error: any) => {
            console.error(`[error] ${error.message}`);
          };
        }
      } catch (error) {
        console.error("Error checking login status:", error);
        router.push("/");
      }
    };

    checkLoginStatus();
  }, []);

  return (
    <div className="flex flex-col h-screen bg-background text-white">
      <TopNav setShowAuthModal={setShowAuthModal} />
      <div className="flex flex-col mt-32 items-center justify-center">
        <div className="text-3xl">You are now in the queue</div>
        <div className="text-7xl my-1">{position !== null ? position : <Spinner />}</div>
        <div className="text-3xl">
          People in front of you
        </div>
        <img src="/queue.gif" className="-mt-28 transform scale-x-[-1]" />
      </div>
      <Footer />
    </div>
  );
};

export default Queue;
