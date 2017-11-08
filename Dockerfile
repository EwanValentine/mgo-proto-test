FROM alpine

MKDIR /app

WORKDIR /app

ADD ./svc /app/svc

CMD ["./svc"]
