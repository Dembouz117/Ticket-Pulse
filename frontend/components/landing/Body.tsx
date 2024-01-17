import { useEffect, useState } from "react";
import ImageCard from "./ImageCard";
import SearchTab from "./SearchTab";
import axios from "axios";
import { getTicketingApiUrl } from "@/utilities/common";
import Spinner from "../utilities/Spinner";
import { addAbortListener } from "events";

type ConcertData = {
  id: string;
  title: string;
  artist: string;
  imageUrl: string;
  edges: any;
};

const Body = () => {
  const [displayedConcertData, setDisplayedConcertData] = useState<
    ConcertData[]
  >([]);
  const [concertData, setConcertData] = useState<ConcertData[]>([]);
  const [concertLoading, setConcertLoading] = useState(true);
  const [activeTab, setActiveTab] = useState("Featured");

  const fetchConcerts = async (tab: string) => {
    let url = `${getTicketingApiUrl()}/api/v1/concerts`;
    if (tab === "Featured") {
      url = `${getTicketingApiUrl()}/api/v1/concerts/featured`;
    }

    try {
      const response = await axios.get(url);
      setDisplayedConcertData(response.data.concerts);
      setConcertData(response.data.concerts);
    } catch (error) {
      console.error("Fetch concerts error:", error);
    } finally {
      setConcertLoading(false);
    }
  };

  useEffect(() => {
    fetchConcerts(activeTab);
  }, [activeTab]);

  const handleSearchResults = (results: ConcertData[]) => {
    setDisplayedConcertData(results);
  };

  if (concertLoading) {
    return <Spinner />;
  }

  return (
    <div className="w-full h-full flex justify-center items-start bg-background">
      <div className="w-full px-36 flex flex-col justify-start items-center gap-[60px]">
        <div className="w-full flex justify-between items-start gap-4">
          <SearchTab
            data={concertData}
            onResults={handleSearchResults}
            activeTab={activeTab}
            setActiveTab={setActiveTab}
          />
        </div>
        <div className="w-full flex justify-center bg-background">
          <div className="w-full grid place-items-center gap-8 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-4">
            {activeTab === "Upcoming"
              ? displayedConcertData
                  .filter((concert: any) => !concert.featured)
                  .map((concert, index) => (
                    <ImageCard
                      key={index}
                      imgUrl={concert.imageUrl}
                      dataTestId="concertCardRenderTest"
                      artist={concert.artist}
                    />
                  ))
              : displayedConcertData.map((concert, index) => (
                  <ImageCard
                    key={index}
                    imgUrl={concert.imageUrl}
                    dataTestId="concertCardRenderTest"
                    artist={concert.artist}
                  />
                ))}
          </div>
        </div>
        <div className="h-12 px-[31px] py-px bg-pink-600 rounded-3xl shadow-inner flex justify-center items-center">
          <div className="text-white text-[15px] font-semibold leading-relaxed">
            More Events
          </div>
        </div>
      </div>
    </div>
  );
};

export default Body;
