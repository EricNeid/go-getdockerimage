# test dockerfile with 2-stage build process

FROM golang:1.13.0-alpine3.10 as builder

RUN apk update && apk add git

RUN adduser -D -g '' appuser
RUN go get github.com/CombatMage/go-httpserver
WORKDIR $GOPATH/src/github.com/CombatMage/go-httpserver
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/server


# STEP 2 build a small image
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/server /server
COPY ./www /www

USER appuser

CMD  ["/server"]