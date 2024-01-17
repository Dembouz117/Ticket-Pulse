import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import CartEntry from '@/components/cart/CartEntry';


// // Mock your atoms with sample data
// jest.mock('jotai', () => ({
//   useAtom: (atom:any) => {
//     if (atom === cartEntryAtom) {
//       return [[{ sectionName: 'Sample Section', price: 10 }], () => {}];
//     }
//     if (atom === currSectionAtom) {
//       return ['Sample Section', () => {}];
//     }
//   },
// }));

describe('CartEntry Component', () => {
  it('renders the CartEntry component', () => {
    render(
      <CartEntry section="Sample Section" className="custom-class" closeSlider={() => {}} />
    );

    const sectionName = screen.getByText('Sample Section');
    const price = screen.getByText('$');
    expect(sectionName).toBeInTheDocument();
    expect(price).toBeInTheDocument();
  });


});
