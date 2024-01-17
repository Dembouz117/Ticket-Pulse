import Cat1Component from "./Cat1Component";
import Cat2Component from "./Cat2Component";
import Cat3Component from "./Cat3Component";

//reminder for parsing: M to move points without drawing, L to draw to an (x,y) and z to roundabout ltr
export const MajorCategory = () => {

  return (
    <svg viewBox="0 0 1250 950" fill="none" xmlns="http://www.w3.org/2000/svg">
        <Cat1Component/>
        <Cat2Component/>
        <Cat3Component/>    
    </svg>
    )
        
}

export default MajorCategory;
