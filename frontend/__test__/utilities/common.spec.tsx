import {
    getAuthApiUrl,
    getTicketingApiUrl,
    getPaymentApiUrl,
    getQueueApiUrl,
    getQueueWebsocketUrl,
    getFrontEndUrl,
  } from "@/utilities/common";
  
  describe('Utility Functions', () => {
    // Test getAuthApiUrl function
    it('should return the frontend API URL from environment variable', () => {
      process.env.NEXT_PUBLIC_FRONTEND_URL = 'https://frontend.example.com';
      const apiUrl = getFrontEndUrl();
      expect(apiUrl).toBe('https://frontend.example.com');
    });
  
    it('should return the default frontend API URL when environment variable is not set', () => {
      delete process.env.NEXT_PUBLIC_FRONTEND_URL;
      const apiUrl = getFrontEndUrl();
      expect(apiUrl).toBe("http://localhost:3000");
    });

    it('should return the authentication API URL from environment variable', () => {
        process.env.NEXT_PUBLIC_AUTH_API_URL = 'https://auth.example.com';
        const apiUrl = getAuthApiUrl();
        expect(apiUrl).toBe('https://auth.example.com');
      });
    
      it('should return the default authentication API URL when environment variable is not set', () => {
        delete process.env.NEXT_PUBLIC_AUTH_API_URL;
        const apiUrl = getAuthApiUrl();
        expect(apiUrl).toBe("http://localhost:8080/auth");
      });

      it('should return the payment API URL from environment variable', () => {
        process.env.NEXT_PUBLIC_PAYMENT_API_URL = 'https://payment.example.com';
        const apiUrl = getPaymentApiUrl();
        expect(apiUrl).toBe('https://payment.example.com');
      });
    
      it('should return the default payment API URL when environment variable is not set', () => {
        delete process.env.NEXT_PUBLIC_PAYMENT_API_URL;
        const apiUrl = getPaymentApiUrl();
        expect(apiUrl).toBe("http://localhost:8082/payment");
      });

      it('should return the queue API URL from environment variable', () => {
        process.env.NEXT_PUBLIC_QUEUE_API_URL = 'https://queue.example.com';
        const apiUrl = getQueueApiUrl();
        expect(apiUrl).toBe('https://queue.example.com');
      });
    
      it('should return the default queue API URL when environment variable is not set', () => {
        delete process.env.NEXT_PUBLIC_QUEUE_API_URL;
        const apiUrl = getQueueApiUrl();
        expect(apiUrl).toBe("http://localhost:8500/queue");
      });

      it('should return the queue websocket API URL from environment variable', () => {
        process.env.NEXT_PUBLIC_QUEUE_WEBSOCKET_URL = 'https://queuesocket.example.com';
        const apiUrl = getQueueWebsocketUrl();
        expect(apiUrl).toBe('https://queuesocket.example.com');
      });
    
      it('should return the default queue web socket API URL when environment variable is not set', () => {
        delete process.env.NEXT_PUBLIC_QUEUE_WEBSOCKET_URL;
        const apiUrl = getQueueWebsocketUrl();
        expect(apiUrl).toBe("ws://localhost:8500/queue/queue");
      });

      it('should return the ticketing websocket API URL from environment variable', () => {
        process.env.NEXT_PUBLIC_TICKETING_API_URL = 'https://ticketing.example.com';
        const apiUrl = getTicketingApiUrl();
        expect(apiUrl).toBe('https://ticketing.example.com');
      });
    
      it('should return the default ticketing web socket API URL when environment variable is not set', () => {
        delete process.env.NEXT_PUBLIC_TICKETING_API_URL;
        const apiUrl = getTicketingApiUrl();
        expect(apiUrl).toBe("http://localhost:8081/ticketing");
      });

      
  
    // Repeat similar tests for other functions like getTicketingApiUrl, getPaymentApiUrl, etc.
  });