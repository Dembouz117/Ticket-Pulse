import React from "react";
import { render, fireEvent, screen } from "@testing-library/react";
import PreCheckOutModal from "@/components/cart/PreCheckOutModal";
import { useToast } from "@chakra-ui/react";
import { useRouter } from "next/router";

// Mock the useToast hook before your tests
jest.mock("@chakra-ui/react", () => ({
  ...jest.requireActual("@chakra-ui/react"),
  useToast: jest.fn(), // Mock useToast
}));

jest.mock("next/router", () => ({
  useRouter: jest.fn(),
}));

// Define the mock function for onClose
const mockClose = jest.fn();

describe("PreCheckOutModal", () => {
  beforeEach(() => {
    // Clear all mocks before each test
    jest.clearAllMocks();

    // Define a mock implementation for useToast
    useToast.mockImplementation(() => ({
      toast: jest.fn(), // this mock function represents the toast method
    }));
  });

  it("closes the modal when the onClose function is triggered", () => {
    // Render your component with the open prop set to true
    render(<PreCheckOutModal open={true} onClose={mockClose} />);
    // Simulate the close action, for example by clicking a close button
    fireEvent.click(screen.getByRole("button", { name: /close/i }));
    // Assertions ...
    expect(mockClose).toHaveBeenCalled();
  });

  // Other tests...
});
