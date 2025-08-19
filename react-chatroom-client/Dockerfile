FROM node:20-alpine AS build-stage

WORKDIR /app

COPY package*.json ./
RUN npm install

# Quellcode kopieren und bauen
COPY . .
RUN npm run build


FROM node:20-alpine

WORKDIR /app

RUN npm install -g serve

COPY --from=build-stage /app ./build

EXPOSE 3000

CMD ["serve", "-s", "build/dist", "-l", "3000"]
