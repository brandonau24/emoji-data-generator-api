FROM docker.io/library/golang:1.23 AS build

WORKDIR /emoji-data-generator

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/emoji-data-generator-api ./cmd/api_server/main.go

FROM docker.io/library/alpine

COPY --from=build /bin/emoji-data-generator-api /bin/emoji-data-generator-api

EXPOSE 8080

CMD ["/bin/emoji-data-generator-api"]