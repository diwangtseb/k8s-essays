FROM golang:1.19-alpine as builder

WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./code/main.go

FROM busybox
COPY --from=builder /build/main . 
EXPOSE 1001
ENTRYPOINT ["./main"]