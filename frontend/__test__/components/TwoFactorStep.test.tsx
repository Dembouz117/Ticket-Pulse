import React from "react";
import { render, fireEvent, screen, act } from "@testing-library/react";
import axios from "axios";
import TwoFactorStep from "@/components/landing/auth/TwoFactorStep";

jest.mock("axios");

const mockClose = jest.fn();
const mockSetStep = jest.fn();

describe("<TwoFactorStep />", () => {
  // Helper function to type into the OTP input fields
  const typeOtp = (otp) => {
    const inputs = screen.getAllByRole("textbox");
    otp.split("").forEach((digit, index) => {
      fireEvent.change(inputs[index], { target: { value: digit } });
    });
  };

  beforeEach(() => {
    render(
      <TwoFactorStep
        onClose={mockClose}
        setStep={mockSetStep}
        email="test@example.com"
        password="password123"
      />
    );
  });

  it("submits the OTP code and handles success response", async () => {
    axios.post.mockResolvedValue({ status: 200 });
    const otp = "123456";
    typeOtp(otp);

    fireEvent.click(screen.getByText("Log In"));

    await act(() => Promise.resolve()); // Wait for the axios promise to resolve

    expect(axios.post).toHaveBeenCalledWith(
      expect.anything(),
      {
        email: "test@example.com",
        otpCode: otp,
      },
      expect.anything()
    );

    expect(mockClose).toHaveBeenCalled();
  });
});
