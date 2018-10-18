FROM alpine:3.7

RUN apk update && apk add --no-cache ca-certificates

COPY bin/ /
WORKDIR /

CMD ["/main"]
