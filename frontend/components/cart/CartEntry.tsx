import { BiSolidDownArrow } from "react-icons/bi";

import { useState, useEffect } from "react";
import QuantityAdjuster from "./QuantityAdjuster";
import { useAtom } from "jotai";

import { cartEntryType } from "@/types/index";
import { cartEntryAtom, currSectionAtom } from "@/store/index";

interface CartEntryProps {
  section: string | String;
  className?: string;
  closeSlider: () => void;
}

//This component handles the individual box entries in the mini slider
const CartEntry = ({ section, className, closeSlider }: CartEntryProps) => {
  //toggles the +/- controls
  const [visibleBox, setVisibleBox] = useState<boolean>(true);
  const [cartEntryList, setCartEntryList] =
    useAtom<cartEntryType[]>(cartEntryAtom);
  const [currEntry, setCurrEntry] = useState<cartEntryType | undefined>(
    undefined
  );

  useEffect(() => {
    const currEntry = cartEntryList.find(
      (entry: cartEntryType) => entry.sectionName === section
    );
    setCurrEntry(currEntry);
  }, [cartEntryList, section]);

  const clickHandler = () => {
    setVisibleBox(!visibleBox);
  };

  return (
    <div
      className={`w-full justify-center items-center flex-col relative px-4`}
    >
      <div className={`flex justify-center items-center`}>
        <div className="bg-white text-black flex justify-between items-center w-full h-auto border-2 border-gray rounded px-2 relative">
          <div className="flex-col w-full m-2">
            <h3 className="mb-4 font-bold">{section}</h3>
            <h3>${currEntry?.price}</h3>
            <div className="flex justify-center w-full">
              <div className="relative w-full transition-transform transform translate-y-0 text-black transition transition-all ease-in duration-1000 rounded border-gray">
                {visibleBox && (
                  <QuantityAdjuster
                    section={section}
                    onClick={clickHandler}
                    onConfirm={closeSlider}
                  />
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CartEntry;
