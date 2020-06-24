FROM golang:alpine AS build-env
WORKDIR /go/src
COPY . /go/src/my-golang
RUN cd /go/src/my-golang && go build .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/my-golang/my-golang /app

EXPOSE 8080
ENTRYPOINT [ "./my-golang" ]
