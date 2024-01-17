import TopNav from "@/components/TopNav";
import Hero from "@/components/landing/Hero";
import Footer from "@/components/Footer";
import AuthModal from "@/components/landing/auth/AuthModal";
import Body from "@/components/landing/Body";
import { useAtom } from "jotai";
import { authModalAtom } from "@/state/globals";


export default function Home() {
  const [showAuthModal, setShowAuthModal] = useAtom(authModalAtom);
  return (
    <div className="h-screen  bg-background text-white" data-testid="homeDiv">
      <AuthModal open={showAuthModal} onClose={() => setShowAuthModal(false)} />
      <TopNav setShowAuthModal={setShowAuthModal} />
      <div className="h-2/3 relative flex-shrink-0 z-0">
        <Hero 
            setShowAuthModal={setShowAuthModal}
        />
      </div>
      <div className="h-full 2xl:h-1/2 w-full relative">
        <Body />
      </div>
      <div className="h-[750px] bg-background"></div>
      <Footer />
    </div>
  );
}
