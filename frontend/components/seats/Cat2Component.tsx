import { useState, useEffect } from "react";
import { useAtom } from "jotai";

import { showCartAtom, currSectionAtom, sectionAtom, errorModalAtom } from "@/store/index";
import { sectionSVGInterface, sectionMappingInterface } from "@/types/index";
import styles from "@/styles/seats/cat2.module.css";
import { initialSvgSectionCat2 } from "@/store/constants/sections";

interface Cat2ComponentInterface{
  className?: string
}


export const Cat2Component = ({className} : Cat2ComponentInterface) => {

  const [ currSection, setCurrSection ] = useAtom(currSectionAtom);
  const [showCart, setShowCart] = useAtom(showCartAtom);
  const [ sections, setSections ] = useAtom(sectionAtom);
  //use placeholder unless sections is not empty
  const [ svgMapList, setSVGMapList] = useState<sectionMappingInterface[] | null>(null);
  const [errorModal, setErrorModal] = useAtom(errorModalAtom);

  const initialList = initialSvgSectionCat2;
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
    <svg viewBox="-20 70 990 200" fill="none" xmlns="http://www.w3.org/2000/svg" onClick = {clickHandler}>
  {
    svgMapList ? svgMapList.map(el=>
          <g>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {`${styles.genericShape} stroke-CAT2`} d = {el.pathShape} stroke-width="4.627"></path>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {styles.genericCircle} d = {el.pathCircle}></path>
            <path data-name={el.dataName} data-section = {el.sectionUuid} data-price = {el.price} className = {styles.genericLetter} d = {el.pathLetter}></path>
          </g>)
        :
    initialList.map(el =>
      <g>
            <path data-name={el.dataName} data-section = {""} className = {`${styles.genericShape} stroke-CAT2`} data-price = {initialPrice} d = {el.pathShape} stroke-width="4.627"></path>
            <path data-name={el.dataName} data-section = {""}  className = {styles.genericCircle} data-price = {initialPrice} d = {el.pathCircle}></path>
            <path data-name={el.dataName} data-section = {""}  className = {styles.genericLetter} data-price = {initialPrice} d = {el.pathLetter}></path>
        </g>
      )
      }       
        </svg>

  )
}

export default Cat2Component;