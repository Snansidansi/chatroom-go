FROM golang:1.23.2 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN ls -al
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /chatroom-server


FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /

COPY --from=build-stage /chatroom-server /chatroom-server

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT ["/chatroom-server", "server", "5000"]
