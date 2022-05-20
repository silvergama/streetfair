FROM alpine:3.13
RUN apk add --no-cache ca-certificates
COPY --from=silvergama/streetfair:test /go/bin/streetfair /usr/bin/streetfair
COPY --from=silvergama/streetfair:test /go/src/github.com/silvergama/streetfair/docs/ /docs/
ENTRYPOINT ["/usr/bin/streetfair"]
EXPOSE 9000