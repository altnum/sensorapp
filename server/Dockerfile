FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go build -o /tmp/sensorapp ./main.go

FROM alpine:edge

COPY --from=build /tmp/sensorapp /sbin/sensorapp

CMD /sbin/sensorapp

EXPOSE 8080