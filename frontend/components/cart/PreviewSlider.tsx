//Mini slider component that shows up on each click of the section
import styles from "@/styles/seats/MiniSlider.module.css";
import { MdClear } from "react-icons/md";

import { cartEntryType } from "@/types/index";

import { showCartAtom, cartEntryAtom, currSectionAtom } from "@/store/index";

import { useAtom } from "jotai";
import { useEffect, useState } from "react";

import CartEntry from "@/components/cart/CartEntry";

//additional components and styles
import {
  Modal,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  ModalOverlay,
} from "@chakra-ui/react";

interface miniSliderProps {
  concertId: string;
}

const PreviewSlider = ({ concertId }: miniSliderProps) => {
  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [sectionQuantityList, setSectionQuantityList] = useAtom(cartEntryAtom);

  const [currSection, setCurrSection] = useAtom(currSectionAtom);

  // const [sectionAmount, setSectionAmount] = useState<number>(0);
  const [windowHeight, setWindowHeight] = useState<number>(0);

  useEffect(() => {
    if (!showCart) {
      document.body.classList.add("overflow-x-hidden");
      document.body.style.overflow = "unset";
    } else {
      document.body.classList.remove("overflow-x-hidden");
      document.body.style.overflow = "hidden";
    }
  }, [showCart]);

  useEffect(() => {
    //adds section to the list if it has not been clicked yet.
    const exists: cartEntryType | undefined = sectionQuantityList.find(
      (el: cartEntryType) => el.sectionName === currSection.name
    );
    if (!exists) {
      console.log(
        "I'm creating a new entry because the current clicked one does not exist"
      );
      const newEntry: cartEntryType = {
        sectionName: currSection.name,
        quantity: 0,
        price: currSection.price,
        sectionId: currSection.sectionUuid,
      };
      setSectionQuantityList([...sectionQuantityList, newEntry]);
      // setSectionAmount(0);
    } else {
      console.log("wtf it exists");
      // setSectionAmount(exists.quantity);
    }
  }, [showCart, currSection]);

  useEffect(() => {
    //unlike self-made modal, chakra ui already keeps track of scroll, so just need to add window offset
    const offsetFactor = 4;
    const windowOffset = window.innerHeight / offsetFactor;
    const handleScroll = () => {
      setWindowHeight(windowOffset);
    };
    //do this first to resolve a bug where window is not loaded yet before scroll
    setWindowHeight(windowOffset);

    window.addEventListener("scroll", handleScroll);

    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <Modal
      isOpen={showCart}
      onClose={() => setShowCart(false)}
      closeOnOverlayClick={true}
      size={"lg"}
      colorScheme="white"
      motionPreset="slideInRight"
    >
      <ModalOverlay />

      <ModalContent
        className={`py-12 px-4 h-auto bg-white gap-4`}
        style={{ top: windowHeight }}
        position="absolute"
        right={5}
      >
        <ModalHeader className="font-3xl">
          Choose how many seats you would like from this section.
        </ModalHeader>
        <ModalCloseButton className="text-button hover:text-black" />
        <ModalBody>
          <div className="h-full w-full flex items-center justify-center rounded-md">
            <CartEntry
              section={currSection.name}
              className="w-1/2"
              closeSlider={() => setShowCart(false)}
            />
          </div>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default PreviewSlider;
