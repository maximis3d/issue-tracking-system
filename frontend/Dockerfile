# Step 1: Build the React app with Vite
FROM node:23 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (make sure they exist)
COPY ./package.json ./package-lock.json ./

# Install dependencies (Vite and other dependencies)
RUN npm install

# Copy the rest of the frontend app
COPY . .

# Build the app with Vite (this creates the /dist folder with production build)
RUN npm run build

# Step 2: Serve the app with Nginx
FROM nginx:alpine

# Copy the build files from the previous stage to Nginx's default folder
COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port the frontend will be served on (default port 80)
EXPOSE 80

# Run Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
