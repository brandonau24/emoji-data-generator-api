FROM docker.io/library/golang:1.23

WORKDIR /emoji-data-generator-api

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build

CMD ["go ./..."]