

const createMockStripe = require('../handlers/createMockStripe');

// Mock the Stripe object using the createMockStripe function
const mockStripe = createMockStripe();

module.exports = {
  loadStripe: jest.fn().mockReturnValue(mockStripe), // Simulate the loadStripe method
};
