FROM golang:alpine AS builder

RUN mkdir /application
ADD . /application

WORKDIR /application

RUN go mod download

RUN go build -v -o danksongs .

FROM alpine:latest
WORKDIR /root
COPY --from=builder /application/danksongs .
EXPOSE 8080

CMD ["./danksongs"]
