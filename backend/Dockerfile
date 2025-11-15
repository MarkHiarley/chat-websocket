FROM golang:1.25.1 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o main ./cmd/

FROM scratch 
COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]