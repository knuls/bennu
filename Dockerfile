FROM golang:1.18.1 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o bennu

FROM alpine:3
COPY --from=builder /app/bennu /
CMD ["/bennu"]