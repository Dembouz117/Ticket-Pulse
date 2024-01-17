import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import { describe } from "node:test";
import Home from "@/pages/index";

// Mock the useRouter hook
jest.mock('next/router', () => ({
    useRouter: jest.fn().mockReturnValue({
      pathname: '/seats', // Adjust this to match the path your component expects
    }),
  }));
  
describe('App', () => {
	it('should render the home component', () => {
		render(<Home/>);
        const homeDiv = screen.getByTestId("homeDiv");
		expect(homeDiv).toBeInTheDocument();
});
});