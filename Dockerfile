FROM node:16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY package.json .
RUN npm install

COPY *go ./
COPY . .

RUN go build -o /server

CMD ["npm", "start"]

EXPOSE 2024
