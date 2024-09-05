FROM docker.io/library/golang:1.23

WORKDIR /emoji-data-generator-api

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./emoji-data-generator-api

EXPOSE 8080

CMD ["./emoji-data-generator-api"]