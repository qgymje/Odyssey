FROM ubuntu:14.04

RUN apt-get update
RUN apt-get install -y ca-certificates

ADD . /app
ADD ./configs /app

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/http"]
