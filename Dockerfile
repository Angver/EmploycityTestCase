FROM alpine:latest

COPY ./employcity-test-case /usr/bin/employcity-test-case

CMD ["employcity-test-case"]
