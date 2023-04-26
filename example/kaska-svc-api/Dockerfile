#Build step
FROM docker.io/golang:1.20.2-alpine3.17 as builder

ENV CMD_PATH cmd/service.go
ENV SVC_NAME kaska-svc-api

ENV WORK_DIR /build_dir

WORKDIR $WORK_DIR
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o $SVC_NAME $CMD_PATH

#Run step
FROM docker.io/alpine:3.17

ENV SVC_NAME kaska-svc-api
ENV WORK_DIR /build_dir

RUN apk update && \
    apk add mailcap tzdata && \
    rm /var/cache/apk/*
    
COPY --from=builder $WORK_DIR/$SVC_NAME .
CMD ./$SVC_NAME

EXPOSE 8081
EXPOSE 8082