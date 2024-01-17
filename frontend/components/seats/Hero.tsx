import Image from "next/image";
import { useEffect } from "react";

interface HeroInterface{
  bannerUrl: string
  session?:any
}

const Hero = ({ bannerUrl, session }: HeroInterface) => {

  useEffect(()=>{
    console.log("In Hero");
    
  },[bannerUrl]);
  return (
    <div className="w-full h-full flex items-center justify-center">
      {/* <Image
        src={"/images/TaylorSwiftBannerShort.png"}
        alt="Hero"
        layout="fill"
        objectFit="cover"
        className="-z-10 object-top"
        draggable={false}
      /> */}
      <img
        src={session?session.edges.ofConcert[0].imageUrl:"https://people.com/thmb/6DM8_L0YQwa2B4Sx3qky2dh2pMY=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc():focal(713x199:715x201)/Taylor-swift-Mexico-City-0824232-af5045dd8be549bcb4c1c85a007c85de.jpg"}
        alt="Hero"
        className="-z-10 object-top absolute"
        style={{ width: '100%', height: '100%', objectFit: 'cover' }}
        draggable={false}
      />


      {/* force content to left side */}
      <div className="h-full w-full items-center px-12 sm:px-24 md:px-32 lg:px-48 pt-8 md:pt-16 lg:pt-48">
        <div className="text-2xl sm:text-3xl md:text-4xl lg:text-5xl xl:text-6xl md:mb-4 font-header">
          {session?session.edges.ofConcert[0].artist:"Taylor Swift"}
        </div>
        <div className="w-full md:w-3/4 lg:w-2/3 mb-12 text-base sm:text-lg md:text-xl xl:text-2xl opacity-80">
          {session?session.edges.ofConcert[0].title:"Era's Tour Singapore"}
        </div>
      </div>
      <div className="absolute bottom-0 left-0 w-full h-[20%] bg-gradient-to-t from-background to-transparent"></div>
    </div>
  );
};

export default Hero;
