FROM golang:1.26.2-bookworm AS build
WORKDIR /app

# copying dependency first so if they haven't changed
# docker can reuse this cached layer and skip downloading them again
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# CGO_ENABLED=0 statically compile the binary (required for distroless)
# -ldflags="-s -w" string debugging info to reduce binary size
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o server ./cmd/main.go

FROM gcr.io/distroless/static-debian13 AS run
WORKDIR /
COPY --from=build /app/server ./server
EXPOSE 8080

# distroless image comes with a built-in 'nonroot'
USER nonroot:nonroot
ENTRYPOINT [ "./server" ]
