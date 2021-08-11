FROM golang:1.16.3 AS BUILDER
LABEL stage="builder"
WORKDIR /go/src/github.com/yannice92/be_tsel_fernando/

ARG CGO_ENABLED=0

COPY . .

RUN go build -o app main.go

FROM alpine
LABEL stage="deploy"

RUN apk update
RUN apk upgrade
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add curl
RUN apk add busybox-extras

WORKDIR /apps/
COPY --from=BUILDER /go/src/github.com/yannice92/be_tsel_fernando/app .
COPY --from=BUILDER /go/src/github.com/yannice92/be_tsel_fernando/.env.example /apps/.env

RUN chgrp -R 0 /apps/ && \
    chmod -R g=u /apps/ /etc/passwd

USER 1001

EXPOSE 8085

RUN pwd & ls -lah

CMD ["./app"]