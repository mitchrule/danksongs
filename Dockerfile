FROM golang:alpine

# Move all files into /app
RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go mod download

RUN go install

CMD ["danksongs"]
