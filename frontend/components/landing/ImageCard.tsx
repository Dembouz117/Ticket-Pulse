import { useRouter } from "next/router";
import {useEffect} from 'react';
import { encodeURIParam } from "@/utilities/paramsParsing"


interface ImageCardProps {
  imgUrl: string;
  dataTestId: string,
  artist: string
  
}


const ImageCard: React.FC<ImageCardProps> = ({ imgUrl, dataTestId, artist }) => {
  
  const router = useRouter();

  useEffect(()=>{

    console.log("in ImageCard", imgUrl);
  },[]);

  const goToConcertHandler = () => {
    const encodedArtistToURIParam = encodeURIParam(artist);
    console.log("INSIDE GO TO CONCERTHANDLER");
    router.push(`/seats/${encodedArtistToURIParam}`);
    if (artist==="Taylor Swift"){
      router.push("/queue");
    }
  }
  return (
    <div 
      className="relative w-1/2 sm:w-2/3 p-[50%] overflow-hidden transform transition-transform duration-200 hover:translate-y-[-10px] cursor-pointer group"
      data-testid={dataTestId}
      onClick={goToConcertHandler}
    >
      <img src={imgUrl} className="absolute inset-0 w-full h-full object-cover transition-brightness duration-200 group-hover:brightness-130" />

      {/* Brightness Effect on Hover */}
      <div className="absolute inset-0 bg-gradient-to-t from-white via-transparent opacity-0 group-hover:opacity-10"></div>
    </div>
  );
};

export default ImageCard;
