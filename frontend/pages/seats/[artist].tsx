import { useState, useEffect } from "react";
import TopNav from "@/components/TopNav";
import Footer from "@/components/Footer";
import Hero from "@/components/seats/Hero";
import Stage from "@/components/seats/Stage.svg";
import MajorCategory from "@/components/seats/majorCategory";
import AuthModal from "@/components/landing/auth/AuthModal";

import { useRouter } from 'next/router';

//types
import { seatCategoryProps, sectionInterface } from "@/types/index";


//state management
import { showCartAtom, showStripeCheckoutAtom, errorModalAtom, cartCoordinatesAtom, sectionAtom } from "@/store/index";
import { userAtom } from "@/state/globals";
import { useAtom } from "jotai";

//Preview slider
import PreviewSlider from "@/components/cart/PreviewSlider";
import PaymentModal from '@/components/payment/PaymentModal';

//Error modal
import ErrorModal from "@/components/errors/ErrorModal";

import axios from "axios";
import { getTicketingApiUrl, getAuthApiUrl } from "@/utilities/common";
import { getSessionId } from "@/utilities/getSessionId";
import { getConcertSession } from "@/utilities/getConcertSession";
import { encodeURIParam, decodeURIParam } from "@/utilities/paramsParsing";
import { useToast } from "@chakra-ui/react";



const initialState: seatCategoryProps = {
  CAT1: false,
  CAT2: false,
  CAT3: false,
};



export default function SeatSelection() {
  const [showAuthModal, setShowAuthModal] = useState<boolean>(false);

  const toast = useToast();

  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [showStripeCheckout, setShowStripeCheckout] = useAtom(showStripeCheckoutAtom);
  const [coordinates, setCoordinates] = useAtom(cartCoordinatesAtom);
  const [errorModal, setErrorModal] = useAtom(errorModalAtom);
  const [sections, setSections] = useAtom(sectionAtom);
  const [catPrices, setCatPrices] = useState({CAT1:1000, CAT2:1000, CAT3:1000});
  const [concertSession, setConcertSession] = useState<any>(null);
  const [heroImage, setHeroImage] = useState<string>("");
  const [user, setUser] = useAtom(userAtom);

  const router = useRouter();
  const { artist } = router.query;


  useEffect(()=>{
    document.body.style.overflowX = 'hidden';
  },[]);

  useEffect(() => {
    // Define the checkLogin function
    const url: string = getAuthApiUrl();
    console.log("seats auth url:")
    console.log(url);
    const checkLogin = async () => {
      console.log("in handler");
      try {
        const response = await axios.get(`${url}/checkLogin`, { withCredentials: true });
        console.log("response status:");
        console.log(response.status);
        if (response.status === 200) {
          console.log("User is logged in");
        } else {
          // console.error("Error checking JWT token broski");
          // setErrorModal({ visible: true, message: "You are not logged in!" });
          router.push("/");
        }
      } catch (error) {
        console.error("Error checking logging status in seats:", error);
        toast({
          title: "Unauthorised User",
          description:
              "Please login",
          status: "error",
          duration: 5000,
          isClosable: true,
          position: "top",
          onCloseComplete: ()=>router.push("/")
      });
        setErrorModal({ visible: true, message: "You are not logged in!" });
        router.push("/");
      }
    };
  
    checkLogin();
  
    const intervalId = setInterval(() => {
      checkLogin();
    }, 3000);
  
    return () => {
      clearInterval(intervalId);
    };
  }, [user]);

  useEffect(()=>{
    const url:string = `${getTicketingApiUrl()}`;

    const artistName: string = decodeURIParam(artist as string);

    // const sessionId = getSessionId(artistName); //consult plan to map concert to session. Right now default to first session of taylor swift
    const getTickets = async () => {
      if (!artist){
        return;
      }
      try{
        const tickets = await getConcertSession(artistName).then((session:any)=> {
          setConcertSession(session);
          setHeroImage(session);
          return session.id;
        })
        .then(sessionId => 
          axios.get(`${url}/api/v1/sessions/${sessionId}/sections`,{withCredentials:true})
          .then(res => {
            console.log("res.data:");
            console.log(res.data);
            // //adjust types later
            const cat1 = res.data.filter((section:sectionInterface) => section.category === "CAT1");
            const cat2 = res.data.filter((section:sectionInterface) => section.category === "CAT2");
            const cat3 = res.data.filter((section:sectionInterface) => section.category === "CAT3");
            setCatPrices({CAT1:cat1[0].price, CAT2:cat2[0].price, CAT3:cat3[0].price})
            setSections(res.data);
          })
          .catch(err => {
            console.log(err.message);
            setErrorModal({visible:true, message:"Something went wrong with fetching the sections!"});
        }))
        .catch(err => setErrorModal({visible:true, message:"Something went wrong with fetching the Concert Session!"}));
        
      }catch (error){
        setErrorModal({visible:true, message:"Something went wrong with fectching the sections!"});
      }
    }
    getTickets();

  },[artist]);



  return (
    <div className="bg-background text-white" style = {{overflowX:'hidden'}}>
      <TopNav setShowAuthModal={setShowAuthModal} />
      {showAuthModal && <AuthModal open={showAuthModal} onClose={() => setShowAuthModal(false)} />}

      <div className="h-1/3 relative flex-shrink-0 z-0">
        <Hero bannerUrl={heroImage} session={concertSession}/>
      </div>

      <div className="my-4 md:my-8 lg:my-12 mx-12 sm:mx-24 md:mx-32 lg:mx-48">
        <div className="flex flex-row items-center justify-between my-4 md:my-8 lg:my-12 w-full sm:w-[65%] md:w-[60%] lg:w-[55%] xl:w-[50%]">
          <div className="flex flex-col text-xs sm:text-sm md:text-base">
            <div onClick={e=>console.log(e)}>Dates and Door Times</div>
            <div className="flex flex-row text-sm sm:text-base md:text-lg lg:text-xl">
              <div>Monday 12 Feb 2024</div>
              <div className="opacity-70 px-2">-</div>
              <div className="opacity-70">6:30 PM</div>
            </div>
          </div>
          <div className="bg-white opacity-50 h-12 w-[1px] mx-4"></div>{" "}
          {/* Hide the divider on small screens */}
          <div className="flex flex-col text-xs sm:text-sm md:text-base">
            <div>Price From</div>
            <div className="flex flex-row text-sm sm:text-base md:text-lg lg:text-xl">
              <div>${catPrices.CAT3}</div>
            </div>
          </div>
        </div>

        <div className="text-base sm:text-lg md:text-xl lg:text-2xl xl:text-3xl mb-4">
          Choose your section
        </div>
        <div className="flex flex-col items-center space-y-12">
          <MajorCategory/>
          <div className="flex justify-center">
            <Stage className="w-[40%] sm:w-[50%] md:w-[60%] lg:w-[70%] xl:w-[100%] -mt-[40%] sm:-mt-[30%] md:-mt-[10%] xl:-mt-[15%]" />
          </div>
        </div>

        <div className="flex flex-col items-center justify-between w-[60%] sm:w-1/2 md:w-1/2 lg:w-2/5 xl:w-1/3 my-4 md:my-8 lg:my-12  text-sm sm:text-base md:text-lg lg:text-xl">
          <div className="flex flex-row justify-between w-full">
            <div className="flex flex-row items-center">
              <div className="w-8 h-4 bg-CAT1"></div>
              <div className="px-2">CAT1</div>
            </div>
            <div>${catPrices.CAT1}</div>
          </div>
          <div className="flex flex-row justify-between w-full">
            <div className="flex flex-row items-center">
              <div className="w-8 h-4 bg-CAT2"></div>
              <div className="px-2">CAT2</div>
            </div>
            <div>${catPrices.CAT2}</div>
          </div>
          <div className="flex flex-row justify-between w-full">
            <div className="flex flex-row items-center">
              <div className="w-8 h-4 bg-CAT3"></div>
              <div className="px-2">CAT3</div>
            </div>
            <div>${catPrices.CAT3}</div>
          </div>
        </div>

        <ul className="text-sm sm:text-base list-disc ml-4 mt-4 opacity-70">
          <li>Seat plan is not drawn to scale</li>
          <li>Colour indicates price category</li>
          <li>Ticket prices exclude booking fees</li>
          <li>Seating layout subject to change</li>
        </ul>
      </div>

      <div className = "inset-0 h-full w-full">
        {/* This is the grey background */}
          {showCart && <div className = "inset-0 w-full h-full opacity-50 bg-gray-900 absolute backdrop-blur-lg" onClick = {() => {
            setShowCart(false);
            }}/>}
          <PreviewSlider concertId={"Taylor Swift"}/>
      </div>

      {showStripeCheckout.visible && <PaymentModal open = {showStripeCheckout.visible} onClose = {() => setShowStripeCheckout({visible: false, amount: 0})}/>}
      {errorModal?.visible && <ErrorModal settings={errorModal} modalHandler={setErrorModal}/>}
      <Footer />
    </div>
  );
}
