FROM debian

WORKDIR /app

ADD ./mgo-proto-test /app/mgo-proto-test

CMD ["./mgo-proto-test"]
ENTRYPOINT ["./mgo-proto-test"]
