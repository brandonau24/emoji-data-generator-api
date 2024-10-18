FROM docker.io/library/golang:1.23 AS build

WORKDIR /emoji-data-generator-api

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/emoji-data-generator-api ./cmd/api_server/main.go

FROM docker.io/library/alpine

RUN adduser -D nonroot
WORKDIR /home/nonroot

COPY --from=build /bin/emoji-data-generator-api .

RUN chown -R nonroot:nonroot /home/nonroot

EXPOSE 8080

USER nonroot

CMD ["/home/nonroot/emoji-data-generator-api"]
