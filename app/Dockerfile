FROM golang:1.24-alpine as build
WORKDIR /src
COPY go.mod .
# COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app -ldflags="-s -w" ./cmd/app

FROM scratch
COPY --from=build /src/app /bin/app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "app" ]
