FROM alpine:3.7
COPY micrologger-tester /micrologger-tester
ENTRYPOINT ["/micrologger-tester"]