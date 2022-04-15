FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /build-api

EXPOSE ${APP_PORT}

CMD [ "/build-api" ]