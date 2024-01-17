import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe } from "node:test";
import Footer from "@/components/Footer";


describe('Footer', () => {
    it("Should render the title", () => {
        render(<Footer/>);
        const expectedMessage = "© 2023 Ticket Pulse. All rights reserved.";
        expect(screen.getByText(expectedMessage)).toBeInTheDocument();
    })
})