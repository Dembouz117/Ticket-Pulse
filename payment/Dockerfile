# Use the Node 20 Alpine image
FROM node:20-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and pnpm-lock.yaml for pnpm installation
COPY ./package.json ./pnpm-lock.yaml ./

# Install global pnpm
RUN npm install -g pnpm

# Install dependencies
RUN pnpm install 

# Copy all the other source files
COPY . .

# Generate Prisma client
RUN pnpm prisma generate

# Build the project
RUN pnpm build

# Start the application
CMD ["pnpm", "start"]
