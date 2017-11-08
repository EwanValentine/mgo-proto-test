FROM alpine

RUN mkdir -p /app

WORKDIR /app

ADD ./svc /app/svc

CMD ["./svc", "--server_address=0.0.0.0:50051"]
