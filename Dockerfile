FROM golang:1.14-alpine AS build
WORKDIR /go/src
COPY . /go/src/my-golang
RUN cd /go/src/my-golang && CGO_ENABLED=0 go build .

FROM scratch
WORKDIR /app
COPY --from=build-env /go/src/my-golang/my-golang /app

EXPOSE 8080
ENTRYPOINT [ "./my-golang" ]
