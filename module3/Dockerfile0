FROM alpine:latest

COPY ./httpserver /usr/local/bin/
ENV MYPATH /usr/local/bin
WORKDIR $MYPATH
EXPOSE 8080
CMD [ "./httpserver"]
