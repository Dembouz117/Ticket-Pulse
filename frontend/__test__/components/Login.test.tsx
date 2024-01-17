import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import LoginStep from "@/components/landing/auth/LoginStep";

describe("<LoginStep />", () => {
  const setStep = jest.fn();
  const setEmail = jest.fn();
  const setPassword = jest.fn();

  beforeEach(() => {
    // Render the component before each test
    render(
      <LoginStep
        setStep={setStep}
        setEmail={setEmail}
        setPassword={setPassword}
      />
    );
  });

  it("allows typing an email and password and clicking the login button", () => {
    // Query the input elements
    const emailInput = screen.getByPlaceholderText(/johndoe@example.com/i);
    const passwordInput = screen.getByPlaceholderText(/Password/i);
    const loginButton = screen.getByRole("button", { name: /log in/i });

    // Simulate typing into the input fields
    fireEvent.change(emailInput, { target: { value: "test@example.com" } });
    fireEvent.change(passwordInput, { target: { value: "password123" } });

    // Check if the input values were updated
    expect(emailInput.value).toBe("test@example.com");
    expect(passwordInput.value).toBe("password123");

    // Simulate clicking the login button
    fireEvent.click(loginButton);
  });
});
