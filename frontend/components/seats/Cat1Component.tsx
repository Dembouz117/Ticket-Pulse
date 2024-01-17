import { useState, useEffect } from "react";
import { useAtom } from "jotai";

import styles from "@/styles/seats/cat1.module.css";

import { showCartAtom, currSectionAtom, sectionAtom, errorModalAtom } from "@/store/index";
import { sectionSVGInterface, sectionMappingInterface } from "@/types/index";
import { initialSvgSectionCat1 } from "@/store/constants/sections";

interface Cat1ComponentInterface{
    className?: string
}

//reminder for parsing: M to move points without drawing, L to draw to an (x,y) and z to roundabout ltr
export const Cat1Component = ({className} : Cat1ComponentInterface) => {

  const [ currSection, setCurrSection ] = useAtom(currSectionAtom);
  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [ sections, setSections ] = useAtom(sectionAtom);
  //use placeholder unless sections is not empty
  const [ svgMapList, setSVGMapList] = useState<sectionMappingInterface[] | null>(null);
  const [errorModal, setErrorModal] = useAtom(errorModalAtom);

  const initialList = initialSvgSectionCat1;
  const initialPrice = 1000;

  const cartSliderHandler = () => {
    setShowCart(true);
  }
  
  const clickHandler = (e: any) => {
    setCurrSection(
      {
        sectionUuid: e.target.dataset.section,
        price: e.target.dataset.price,
        name: e.target.dataset.name
      }
    )
    cartSliderHandler();
  }

  //map database sections to svg sections
  useEffect(() => {
    let errorDetectingUuid = false;
    //map db to svg sections
    const newMappedList: sectionMappingInterface[] = initialList.map(el => {
      const matchedSection = sections.find(section => `${section.category}-${section.name}` === el.dataName);
      if (!matchedSection && sections.length>0){
        errorDetectingUuid = true;
        setErrorModal({visible:true, message:"Error: section not found in database"});
      }
      return {
        ...el,
        sectionUuid: matchedSection ? matchedSection.id : "",
        price: matchedSection ? matchedSection.price : initialPrice

      };
    });
    if(!errorDetectingUuid){
      setSVGMapList(newMappedList);
    }

  }, [sections]);



  return (
    <svg id = "cat1SVG" viewBox="-130 0 1000 10" fill="none" xmlns="http://www.w3.org/2000/svg" className = "bg-red" onClick = {clickHandler}>
  {
    svgMapList ? svgMapList.map(el=>
          <g>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {`${styles.genericShape} stroke-CAT1`} d = {el.pathShape} stroke-width="4.627"></path>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {styles.genericCircle} d = {el.pathCircle}></path>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {styles.genericLetter} d = {el.pathLetter}></path>
          </g>)
        :
    initialList.map(el =>
      <g>
            <path data-name={el.dataName} data-section = {""} className = {`${styles.genericShape} stroke-CAT1`} data-price = {initialPrice} d = {el.pathShape} stroke-width="4.627"></path>
            <path data-name={el.dataName} data-section = {""}  className = {styles.genericCircle} data-price = {initialPrice} d = {el.pathCircle}></path>
            <path data-name={el.dataName} data-section = {""}  className = {styles.genericLetter} data-price = {initialPrice} d = {el.pathLetter}></path>
        </g>
      )
      }       
        </svg>
  )
}

export default Cat1Component;
