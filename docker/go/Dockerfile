FROM golang:1.22-alpine

RUN apk add --no-cache \
    chromium \
    chromium-chromedriver

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api_server ./app/api

RUN go build -o cli_tool ./app/cmd

ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/ \
    CHROMEDRIVER_PATH=/usr/bin/chromedriver

CMD ["./api_server"]
