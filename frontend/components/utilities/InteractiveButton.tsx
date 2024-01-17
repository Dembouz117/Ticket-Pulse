
import { FC, useEffect } from 'react';

interface interactiveButtonProps{
    onClick:()=>void,
    children?: React.ReactNode,
    valid:boolean
}

const InteractiveButton: FC<interactiveButtonProps> = ({onClick, children, valid}) => {
  useEffect(()=>{
    console.log("valid: ", valid);
  }
  ,[valid]);
  return (
    <button disabled = {!valid} onClick={onClick} className={
        valid ? 
        "bg-button text-white py-2 px-4 rounded hover:bg-pink-500"
        :
        "bg-gray-500 text-white py-2 px-4"
    }>
        {children}
    </button>
  )
}

export default InteractiveButton;