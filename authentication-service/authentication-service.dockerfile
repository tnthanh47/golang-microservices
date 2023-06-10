FROM alpine:latest

RUN mkdir /app

COPY authenApp /app

CMD ["/app/authenApp"]
