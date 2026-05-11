FROM golang:1.24.3 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod ./
RUN go mod download

COPY . .
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/ainews ./cmd/ainews

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/ainews /ainews
EXPOSE 8080
ENTRYPOINT ["/ainews"]
