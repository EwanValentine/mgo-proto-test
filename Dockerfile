FROM alpine

RUN mkdir -p /app

WORKDIR /app

ADD ./mgo-proto-test /app/mgo-proto-test

CMD ["./mgo-proto-test", "--server_address=0.0.0.0:50051"]
