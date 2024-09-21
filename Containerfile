FROM docker.io/library/golang:1.23 AS lint

WORKDIR /emoji-data-generator

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3

COPY . ./

RUN golangci-lint run

FROM lint AS unit_tests

RUN go test ./... -v

FROM unit_tests AS build

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/emoji-data-generator-api

FROM docker.io/library/alpine

COPY --from=build /bin/emoji-data-generator-api /bin/emoji-data-generator-api

EXPOSE 8080

CMD ["/bin/emoji-data-generator-api"]