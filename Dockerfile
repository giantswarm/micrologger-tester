FROM alpine:3.8
COPY micrologger-tester /micrologger-tester
ENTRYPOINT ["/micrologger-tester"]
