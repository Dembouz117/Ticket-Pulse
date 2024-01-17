import React from 'react';
import { render, waitFor, screen, getByTestId, act, fireEvent } from '@testing-library/react';
// import userEvent from '@testing-library/user-event';

import Body from '@/components/landing/Body';
import SearchTab from '@/components/landing/SearchTab';
import SearchBar from '@/components/utilities/SearchBar';

import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import { getTicketingApiUrl } from '@/utilities/common';
import { concertInterface } from '@/types';



const mock = new MockAdapter(axios);

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

const mockResponse = {
    "concerts": [
        {
            "id": "2948e5b6-ae2b-45bb-8b86-5d2dda89a3e8",
            "title": "NewJeans: Bunnie Fest",
            "artist": "NewJeans",
            "imageUrl": "https://www.allkpop.com/upload/2022/08/content/022201/web_data/allkpop_1659492146_20220802-newjeans.jpg",
            "edges": {}
        },
        {
            "id": "75a7c937-b96c-4941-b1eb-761bfea62407",
            "title": "NIKI Nicole Tour 2022",
            "artist": "NIKI",
            "imageUrl": "https://ahboy.com/wp-content/uploads/2022/10/NIKI-Concert-Singapore-2022.jpg",
            "edges": {}
        },
    ]
}

describe('Body', () => {

  it('Renders the concert cards after a successful API call', async () => {
    // Mock Axios response

    const ticketUrl = getTicketingApiUrl();
    console.log("test ticketUrl: ", ticketUrl);
    mock.onGet(`${ticketUrl}/api/v1/concerts/featured`).reply(200, mockResponse);

    render(<Body />);

    //expects multiple ticket elements to be rendered
    await waitFor(() => {
        expect(screen.getAllByTestId("concertCardRenderTest").length).toBeGreaterThan(0);
    });

    // Verify that loading state is no longer displayed
    expect(screen.queryByText('Loading...')).toBeNull();
  });

  it ('Renders the mini navigation tab across sizes', () => {
    const ticketUrl = getTicketingApiUrl();
    mock.onGet(`${getTicketingApiUrl()}/api/v1/concerts`).reply(200, mockResponse);
    const mockConcertData: concertInterface[] = [{
        id: "artistId",
        title: "artist concert title",
        artist: "artist name",
        imageUrl: "../../public/images/TaylorSwiftBannerShort.png",
        edges: {}
    }];
    const mockOnResultsHandler = (results: concertInterface[])=>{};
    const mockActiveTabHandler = (tab: string) => {};


    render(<SearchTab data={mockConcertData} onResults={mockOnResultsHandler} setActiveTab={mockActiveTabHandler} activeTab={"All Events"}/>);
    const element = screen.getByTestId("searchTabTest");
    expect(element).toBeInTheDocument();
  });

  it('Renders the search bar', () => {
    //mock
    const onSearch = jest.fn();
    render(<SearchBar setInput={onSearch}/>);
    const element = screen.getByTestId("searchBarTest");
    //assert
    expect(element).toBeInTheDocument();
  });

  it('calls onSearch after debouncing when input changes', () => {
    //arrange
    //mock timers
    jest.useFakeTimers(); 

    const onSearch = jest.fn();
    render(<SearchBar setInput={onSearch} />);

    const inputElement = screen.getByTestId('searchBarTest');
    fireEvent.change(inputElement, { target: { value: 'search term' } });

    // fast-forward timers to trigger the debounce
    act(() => {
      jest.runAllTimers(); 
    });

    expect(onSearch).toHaveBeenCalledWith('search term');
  });

  it('clears the timer on unmount', () => {
    jest.useFakeTimers();
    const onSearch = jest.fn();
    const { unmount } = render(<SearchBar setInput={onSearch} />);

    unmount();

    act(() => {
      jest.runAllTimers();
    });

    expect(onSearch).not.toHaveBeenCalled(); 
  });

  it('handles input change correctly', () => {
    const onSearch = jest.fn();
    render(<SearchBar setInput={onSearch} />);

    const inputElement: HTMLInputElement = screen.getByRole('textbox');
    fireEvent.change(inputElement, { target: { value: 'new value' } });

    // onSearch should not be called immediately cos of debouncing logic
    expect(onSearch).not.toHaveBeenCalled();
    // test state value is updating immediately
    expect(inputElement.value).toBe('new value'); 
  });
});
