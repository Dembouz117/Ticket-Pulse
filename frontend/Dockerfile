# Set the base image to Node
FROM node:16-alpine

# Set the working directory in Docker
WORKDIR /usr/app

# Copy package.json and package-lock.json to the root directory of our app
COPY pnpm-lock.yaml package.json ./

# Install dependencies in the container
RUN npm install -g pnpm

# Install dependencies in the container using pnpm
RUN pnpm install 

# Copy the content of your repository into the container
COPY . .

ARG NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY
ENV NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY=$NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY

ARG NEXT_PUBLIC_AUTH_API_URL
ENV NEXT_PUBLIC_AUTH_API_URL=$NEXT_PUBLIC_AUTH_API_URL

ARG NEXT_PUBLIC_TICKETING_API_URL
ENV NEXT_PUBLIC_TICKETING_API_URL=$NEXT_PUBLIC_TICKETING_API_URL

ARG NEXT_PUBLIC_PAYMENT_API_URL
ENV NEXT_PUBLIC_PAYMENT_API_URL=$NEXT_PUBLIC_PAYMENT_API_URL

ARG NEXT_PUBLIC_QUEUE_API_URL
ENV NEXT_PUBLIC_QUEUE_API_URL=$NEXT_PUBLIC_QUEUE_API_URL

ARG NEXT_PUBLIC_QUEUE_WEBSOCKET_URL
ENV NEXT_PUBLIC_QUEUE_WEBSOCKET_URL=$NEXT_PUBLIC_QUEUE_WEBSOCKET_URL

ARG NEXT_PUBLIC_FRONTEND_URL
ENV NEXT_PUBLIC_FRONTEND_URL=$NEXT_PUBLIC_FRONTEND_URL

# Build the Next.js app
RUN pnpm run build

# Expose port 3000 to be accessed outside the container
EXPOSE 3000


# Command to run the application
CMD ["pnpm", "start"]
