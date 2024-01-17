import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe } from "node:test";
import ErrorModal from "@/components/errors/ErrorModal";
 
const mockErrorSettings= {
  visible: true,
  message: "This is a test error message"
}

 describe('ErrorModal', () => {
   it('Renders the error text', () => {
     render(<ErrorModal settings={mockErrorSettings} modalHandler={jest.fn()}/>)
     const expectedMessage = "Unfortunately, something went wrong.";
     expect(screen.getByText(expectedMessage)).toBeInTheDocument();
   });

   it('renders with the given error message', () => {

    render(<ErrorModal settings={mockErrorSettings} modalHandler={jest.fn()} />);

    const modalTitle = screen.getByText('Unfortunately, something went wrong.');
    const modalMessage = screen.getByText(mockErrorSettings.message);

    expect(modalTitle).toBeInTheDocument();
    expect(modalMessage).toBeInTheDocument();
  });

  it('calls the onClose function when the modal is closed', () => {

    const onClose = jest.fn();
    render(<ErrorModal settings={mockErrorSettings} modalHandler={onClose} />);

    const closeButton = screen.getByLabelText('Close');

    fireEvent.click(closeButton);

    expect(onClose).toHaveBeenCalled();
  });

  it('does not render if visible is false', () => {
    const mockErrorSettingsInvisible = {
      visible: false,
      message: "Test error message"
    }
    const onClose = jest.fn();
    render(<ErrorModal settings={mockErrorSettingsInvisible} modalHandler={onClose}/>);

    const modalTitle = screen.queryByText('Unfortunately, something went wrong.');
    const modalMessage = screen.queryByText('Error message');

    expect(modalTitle).toBeNull();
    expect(modalMessage).toBeNull();
  });
 })