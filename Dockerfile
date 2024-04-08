FROM golang:alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux CGO_ENABLED=0 go build -o /streamd ./cmd/streamd/

FROM alpine:latest AS build-release-stage

WORKDIR /
COPY --from=build-stage /streamd /streamd

EXPOSE 8065

CMD ["/streamd"]
