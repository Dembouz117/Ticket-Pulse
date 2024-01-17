import { render, waitFor, screen, getByTestId, fireEvent } from '@testing-library/react';
// import { useRouter } from '@/__mocks__/router';

import QuantityAdjuster from '@/components/cart/QuantityAdjuster';
import PreCheckoutModal from "@/components/cart/PreCheckOutModal";


jest.mock('next/router', () => ({
  useRouter() {
    return {
      route: '/seats',
      pathname: '',
      query: '',
      asPath: '',
      push: jest.fn(),
      events: {
        on: jest.fn(),
        off: jest.fn()
      },
      beforePopState: jest.fn(() => null),
      prefetch: jest.fn(() => null)
    };
  },
}));

describe("Quantity Adjuster", () => {
    //user acceptance criteria 0
    it('I want the default quantity to be 0 before I add to cart/as I browse sections', () => {
        const onClickMock = jest.fn();
        const onConfirmMock = jest.fn();
        render(<QuantityAdjuster section="sectionTest" onClick={onClickMock} onConfirm={onConfirmMock}/>);
        expect(screen.queryByText("0")).toBeInTheDocument();
    });

    //user acceptance criteria 1
    it('I want to be able to add quantities of a section to my cart', async () => {
            const onClickMock = jest.fn();
            const onConfirmMock = jest.fn();
            render(<QuantityAdjuster section="sectionTest" onClick={onClickMock} onConfirm={onConfirmMock}/>);
            fireEvent.click(screen.getByTestId("addTest"));
            
            await waitFor(() => expect(screen.getByTestId("temporaryQuantityTest").textContent).toBe("1"));
            expect(screen.queryByText("1")).toBeInTheDocument();
    });
    //user acceptance criteria 2
    it('I want to be able to add and subtract quantities in sequence', async () => {
   
      const onClickMock = jest.fn();
      const onConfirmMock = jest.fn();
      render(<QuantityAdjuster section="sectionTest" onClick={onClickMock} onConfirm={onConfirmMock}/>);

        for (let i = 0; i < 5; i++) {
            fireEvent.click(screen.getByTestId("addTest"));
            if (i%2==0){
                fireEvent.click(screen.getByTestId("minusTest"));
            }    
        }
        await waitFor(() => expect(screen.getByTestId("temporaryQuantityTest").textContent).toBe("2"));
        expect(screen.queryByText("2")).toBeInTheDocument();
});
    //user acceptance criteria 3
    it('Additional validation so that I do not submit negative amounts', () => {
        const onClickMock = jest.fn();
        const onConfirmMock = jest.fn();
        render(<QuantityAdjuster section="sectionTest" onClick={onClickMock} onConfirm={onConfirmMock}/>);
        fireEvent.click(screen.getByTestId("minusTest"));
        expect(screen.queryByText("0")).toBeInTheDocument();
    });

    //user acceptance criteria 4
    it('should trigger a toast when quantity exceeds 4', async () => {
        const onClickMock = jest.fn();
        const onConfirmMock = jest.fn();
      
        render(
          <>
            <QuantityAdjuster section="sectionTest" onClick={onClickMock} onConfirm={onConfirmMock}/>
            <PreCheckoutModal open={true} onClose={()=>{}}/> {/* The component responsible for showing Chakra UI toasts */}
          </>
        );
      
        // Trigger quantity adjustments
        const addTestButton = screen.getByTestId('addTest');
        for (let i = 0; i < 5; i++) {
          fireEvent.click(addTestButton);
        }
      
        // Invalid amount validation
        await waitFor(() => {
          const disabledButton = screen.getByText('Confirm Seats', { selector: 'button.bg-gray-500.text-white.py-2.px-4' });
          expect(disabledButton).toBeInTheDocument();
        });
      });
});


