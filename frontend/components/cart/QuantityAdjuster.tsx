import { useState, useEffect } from "react";
import { cartEntryAtom } from "@/store/index";
import { useAtom } from "jotai";
import { cartEntryType } from "@/types/index";

interface QuantityAdjusterProps {
  section: string | String;
  onClick: () => void;
  onConfirm: () => void;
}

const QuantityAdjuster = ({ section, onConfirm }: QuantityAdjusterProps) => {
  const [cartEntry, setCartEntry] = useAtom(cartEntryAtom);
  const [temporaryQuantity, setTemporaryQuantity] = useState(0); // Temporary state
  const [isCartEntryOpen, setIsCartEntryOpen] = useState(true);

  // Handler to toggle CartEntry
  // const toggleCartEntry = () => {
  //   setIsCartEntryOpen(!isCartEntryOpen);
  // };

  // This finds the current entry in the cart for the specific section
  const currEntry = cartEntry.find(
    (entry: cartEntryType) => entry.sectionName === section
  );

  // Initialize the temporary quantity from the cart entry when the component mounts or section changes
  useEffect(() => {
    setTemporaryQuantity(currEntry?.quantity || 0);
  }, [currEntry, section]);

  // Handler to change the temporary quantity before confirmation
  const handleTemporaryQuantity = (type: string) => {
    setTemporaryQuantity((prevQuantity) => {
      if (type === "add") {
        return prevQuantity + 1;
      } else if (type === "minus") {
        return prevQuantity > 0 ? prevQuantity - 1 : 0;
      }
      return prevQuantity;
    });
  };

  // Confirm button handler to update the cart
  const confirmQuantity = () => {
    setCartEntry((currCartEntry: cartEntryType[]) =>
      currCartEntry.map((entry: cartEntryType) => {
        if (entry.sectionName === section) {
          return { ...entry, quantity: temporaryQuantity };
        }
        return entry;
      })
    );

    // onClick();
    onConfirm();
    setIsCartEntryOpen(!isCartEntryOpen);
  };

  return (
    <div className="flex flex-col justify-center h-full">
      <div className="flex flex-row bg-white justify-between h-full mb-4">
        <button
          className="w-1/3 ml-4 flex justify-center items-center"
          onClick={() => handleTemporaryQuantity("minus")}
          data-testid="minusTest"
        >
          -
        </button>
        <div
          className="w-1/3 h-auto p-2 flex justify-center items-center"
          data-testid="temporaryQuantityTest"
        >
          {temporaryQuantity}
        </div>
        <button
          className="w-1/3 mr-4 h-auto flex justify-center items-center"
          onClick={() => handleTemporaryQuantity("add")}
          data-testid="addTest"
        >
          +
        </button>
      </div>
      <button
        className="bg-pink-500 hover:bg-pink-600 text-white py-1 mx-20 rounded-lg"
        onClick={confirmQuantity}
        data-testid="confirmQuantityTest"
      >
        Confirm
      </button>
    </div>
  );
};

export default QuantityAdjuster;
