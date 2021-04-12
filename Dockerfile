FROM golang:alpine AS builder

# Move all files into /app
RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -v -o danksongs .

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/danksongs .
ENTRYPOINT ./danksongs
# CMD ["danksongs"]
