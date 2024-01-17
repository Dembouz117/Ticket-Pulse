

const createMockStripe = () => {
    return {
      // Mock methods
      elements: jest.fn().mockReturnValue({
        create: jest.fn(),
      }),
      createPaymentMethod: jest.fn(),
      redirectToCheckout: jest.fn(),
      confirmCardPayment: jest.fn(),
      // Mock properties
      elements: {
        create: jest.fn(),
      },
    };
  };
  
module.exports = {
    createMockStripe: createMockStripe
}
  