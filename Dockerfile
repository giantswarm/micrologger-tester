FROM alpine:3.7
COPY logger-test /logger-test
ENTRYPOINT ["/logger-test"]