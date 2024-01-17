import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import SignUpStep from "@/components/landing/auth/SignUpStep";

describe("<SignUpStep />", () => {
  const onClose = jest.fn();
  const setStep = jest.fn();
  const setEmail = jest.fn();
  const setPassword = jest.fn();

  beforeEach(() => {
    render(
      <SignUpStep
        onClose={onClose}
        setStep={setStep}
        setEmail={setEmail}
        setPassword={setPassword}
      />
    );
  });

  it("allows entering details for signup", () => {
    // Find input fields by placeholder or label text
    const emailInput = screen.getByPlaceholderText("johndoe@example.com");
    const nameInput = screen.getByPlaceholderText("Name");
    const passwordInput = screen.getByPlaceholderText("Password");
    const confirmPasswordInput =
      screen.getByPlaceholderText("Confirm Password");
    const phoneNumberInput = screen.getByPlaceholderText("Enter phone number");

    const createAccountButton = screen.getByRole("button", {
      name: /create account/i,
    });

    // Simulate user typing into the input fields
    fireEvent.change(emailInput, { target: { value: "test@example.com" } });
    fireEvent.change(nameInput, { target: { value: "Test User" } });
    fireEvent.change(passwordInput, { target: { value: "Password123!" } });
    fireEvent.change(confirmPasswordInput, {
      target: { value: "Password123!" },
    });
    fireEvent.change(phoneNumberInput, { target: { value: "+1234567890" } }); // Assuming '+1234567890' is a valid phone number format

    // Check if the input values are updated
    expect(emailInput.value).toBe("test@example.com");
    expect(nameInput.value).toBe("Test User");
    expect(passwordInput.value).toBe("Password123!");
    expect(confirmPasswordInput.value).toBe("Password123!");
    expect(phoneNumberInput.value).toBe('+1 234 567 890');

    // Simulate clicking the create account button
    fireEvent.click(createAccountButton);

    // Here, you might want to check if the onSubmit function was called, but that requires
    // the function to be injected or spying on the component's internal function call.
    // Instead, we can mock network requests and wait for them to be called, but that's beyond
    // the scope of this test. This test just ensures that we can type into the fields and
    // click the button.
  });

  // More tests...
});
