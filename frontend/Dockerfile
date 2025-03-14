FROM node:23-alpine AS build
# Set the working directory inside the container
WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD [ "npm", "run", "dev" ]