export const getAuthApiUrl = () => {
    let url = process.env.NEXT_PUBLIC_AUTH_API_URL
        ? process.env.NEXT_PUBLIC_AUTH_API_URL
        : "";
    if (!url) {
        url = "http://localhost:8080/auth";
    }
    return url;
};

export const getTicketingApiUrl = () => {
    let url = process.env.NEXT_PUBLIC_TICKETING_API_URL
        ? process.env.NEXT_PUBLIC_TICKETING_API_URL
        : "";
    if (!url) {
        url = "http://localhost:8081/ticketing";
    }
    return url;
};

export const getPaymentApiUrl = () => {
    let url = process.env.NEXT_PUBLIC_PAYMENT_API_URL
        ? process.env.NEXT_PUBLIC_PAYMENT_API_URL
        : "";
    if (!url) {
        url = "http://localhost:8082/payment";
    }
    return url;
};

export const getQueueApiUrl = () => {
    let url = process.env.NEXT_PUBLIC_QUEUE_API_URL
        ? process.env.NEXT_PUBLIC_QUEUE_API_URL
        : "";
    if (!url) {
        url = "http://localhost:8500/queue";
    }
    return url;
};

export const getQueueWebsocketUrl = () => {
    let url = process.env.NEXT_PUBLIC_QUEUE_WEBSOCKET_URL
        ? process.env.NEXT_PUBLIC_QUEUE_WEBSOCKET_URL
        : "";
    if (!url) {
        url = "ws://localhost:8500/queue/queue";
    }
    return url;
};

export const getFrontEndUrl = () => {
    let url = process.env.NEXT_PUBLIC_FRONTEND_URL
        ? process.env.NEXT_PUBLIC_FRONTEND_URL
        : "";
    if (!url) {
        url = "http://localhost:3000";
    }
    return url;
};
