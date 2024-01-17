import Image from "next/image";
import config from "../../config.json";
import Link from "next/link";
import { useAtomValue } from "jotai";
import { userAtom } from "@/state/globals";

interface HeroProps {
  setShowAuthModal: (open: boolean) => void;
}

const Hero = (props: HeroProps) => {
  const user = useAtomValue(userAtom);

  const handleRedirect = () => {
    if (user) {
      window.location.href = "/waitingRoom";
    } else {
      // If no user, open the auth modal
      props.setShowAuthModal(true);
    }
  };

  return (
    <div
      className="w-full h-full flex items-center justify-center"
      data-cy="hero"
    >
      <Image
        src="/images/TaylorSwiftBanner.png"
        alt="Hero"
        layout="fill"
        objectFit="cover"
        className="-z-10"
        draggable={false}
      />

      {/* force content to left side */}
      <div className="h-full w-full grid items-center grid-cols-2 p-36">
        <div className="col-start-1 col-end-2 p-12">
          <div className="text-5xl mb-4 font-header">{config.heroTitle}</div>
          <div className="w-2/3 mb-8">{config.heroSubtitle}</div>
          <div className="flex flex-row justify-between w-[60%]">
            <button
              onClick={handleRedirect}
              className="bg-white hover:bg-gray-300 text-black rounded-2xl px-4 py-2"
            >
              Buy Tickets
            </button>
            <button className=" text-white border border-white rounded-2xl px-4 py-2">
              Learn More
            </button>
          </div>
        </div>
      </div>

      {/* gradient */}
      <div className="absolute pointer-events-none bottom-0 left-0 w-full h-[20%] bg-gradient-to-t from-background to-transparent"></div>
    </div>
  );
};

export default Hero;
