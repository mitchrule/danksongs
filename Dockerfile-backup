FROM node AS nodebuild
RUN mkdir /application
ADD . /application
WORKDIR /application/ui
# RUN cd ui/
RUN npm install
RUN npm run-script build

FROM golang:alpine AS gobuild
WORKDIR /application
COPY --from=nodebuild /application .
WORKDIR /application/backend
RUN go mod download
RUN go build -v -o danksongs .

FROM alpine:latest
WORKDIR /root
COPY --from=gobuild /application .
EXPOSE 8080

# CMD [ "pwd" ]
CMD ["./backend/danksongs"]
