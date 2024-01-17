import { useState, useEffect } from "react";

import TabUnderline from "../utilities/TabUnderline";
import SearchBar from "../utilities/SearchBar";
import { concertInterface } from "@/types/index"

// type Concert = {
//   id: string;
//   title: string;
//   artist: string;
//   imageUrl: string;
//   edges: any;
// };

type SearchTabProps = {
  data: concertInterface[];
  onResults: (results: concertInterface[]) => void;
  activeTab: string;
  setActiveTab: (tab: string) => void;
};

const SearchTab = ({
  data,
  onResults,
  activeTab,
  setActiveTab,
}: SearchTabProps) => {
  const [searchTerm, setSearchTerm] = useState("");

  const searchHandler = (query: string) => {
    setSearchTerm(query);

    const queryLowered = query.toLowerCase();
    const filteredData = data.filter(
      (concert) =>
        concert.title.toLowerCase().includes(queryLowered) ||
        concert.artist.toLowerCase().includes(queryLowered)
    );

    onResults(filteredData);
  };

  const changeActiveTab = (tabName: string) => {
    setActiveTab(tabName);
  };

  return (
    <div
      className="w-full flex justify-around items-start"
      data-testid="searchTabTest"
    >
      <div className="flex flex-row">
        <div className="flex justify-between mx-6">
          <div className="mx-2" onClick={() => changeActiveTab("All Events")}>
            <TabUnderline
              text={"All Events"}
              active={activeTab === "All Events"}
            />
          </div>
          <div className="mx-2" onClick={() => changeActiveTab("Featured")}>
            <TabUnderline text={"Featured"} active={activeTab === "Featured"} />
          </div>
          <div className="mx-2" onClick={() => changeActiveTab("Upcoming")}>
            <TabUnderline text={"Upcoming"} active={activeTab === "Upcoming"} />
          </div>
        </div>
        <div className="flex items-end">
          <SearchBar setInput={searchHandler} />
        </div>
      </div>
    </div>
  );
};

export default SearchTab;
