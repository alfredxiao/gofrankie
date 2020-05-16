FROM golang:1.14.1 AS builder

RUN mkdir /appbuild 
WORKDIR /appbuild
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.11.6

RUN mkdir /application 
WORKDIR /application

COPY --from=builder /appbuild/gofrankie /application/

CMD ["/application/gofrankie"]
EXPOSE 8080

# TODO: not run as root