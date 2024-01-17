import React, { useState, useEffect } from "react";

interface SearchBarInterface {
  // input: string;
  setInput: (query: string) => void;
}

const SearchBar = ({ setInput }: SearchBarInterface) => {
  const [searchText, setSearchText] = useState("");

  useEffect(() => {
    const timer = setTimeout(() => {
      setInput(searchText);
    }, 300); // Set a delay of 300 milliseconds

    return () => clearTimeout(timer); // Clear the timer on component unmount or before the next useEffect execution
  }, [searchText, setInput]); // Only re-run the effect if searchText or setInput changes

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchText(event.target.value);
  };

  return (
    <input
      className="text-gray-400 w-inline py-2 px-2 border-solid border-2 border-button rounded-lg bg-background"
      placeholder="search"
      value={searchText}
      onChange={handleInputChange}
      data-testid="searchBarTest"
    />
  );
};

export default SearchBar;