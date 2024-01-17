import React from 'react';
import { render, screen } from '@testing-library/react';
import Hero from '@/components/landing/Hero';
import { Provider } from 'jotai';

describe('Hero', () => {
  it('checks if any image within the Hero component has loaded', async () => {
    render(
      <Provider>
        <Hero setShowAuthModal={() => {}} />
      </Provider>
    );

    const images = screen.getAllByRole('img');
    expect(images).toHaveLength(1); 
    images.forEach((image) => {
      expect(image).toHaveAttribute('src');
      expect(image).toBeVisible();
    });
  });
});
