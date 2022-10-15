FROM golang:1.16-alpine AS build
COPY . /go/src/module3
WORKDIR /go/src/module3
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/httpserver

FROM alpine:latest
COPY --from=build /go/bin/httpserver /usr/local/bin/
ENV MYPATH /usr/local/bin
WORKDIR $MYPATH
EXPOSE 8080
CMD [ "./httpserver"]
