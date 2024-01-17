import React, { MouseEvent, TouchEvent } from "react";

interface BidChangeButtonProps {
    amount: number;
    action: (amount: number) => void;
    endAction: () => void;
    label: string;
}

const CartAmountButton: React.FC<BidChangeButtonProps> = ({
    amount,
    action,
    endAction,
    label,
}) => {
    const handleMouseDown = (
        e: MouseEvent<HTMLButtonElement> | TouchEvent<HTMLButtonElement>
    ) => {
        e.preventDefault();
        action(amount);
    };

    const handleMouseUp = (
        e: MouseEvent<HTMLButtonElement> | TouchEvent<HTMLButtonElement>
    ) => {
        e.preventDefault();
        endAction();
    };

    return (
        <button
            className="text-white bg-neutral-700 hover:bg-neutral-600 h-full p-4 rounded-lg text-center w-full md:text-md lg:text-lg"
            onMouseDown={handleMouseDown}
            onMouseUp={handleMouseUp}
            onMouseLeave={handleMouseUp}
            onTouchStart={handleMouseDown}
            onTouchEnd={handleMouseUp}
            onContextMenu={(e) => e.preventDefault()}
        >
            {label}
        </button>
    );
};

export default CartAmountButton;
