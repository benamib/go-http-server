FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server -ldflags="-w -s" main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/server /usr/bin/

ENTRYPOINT ["server"]