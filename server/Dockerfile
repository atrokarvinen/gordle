# syntax=docker/dockerfile:1

FROM golang:1.21.6 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /gordle

# CMD ["/gordle"]

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /gordle /gordle

EXPOSE 8080

# USER nonroot:nonroot
USER 1001:1001

ENTRYPOINT ["/gordle"]