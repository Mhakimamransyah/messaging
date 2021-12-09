FROM golang:1.16-alpine AS builder

RUN mkdir /projects
ADD . /projects
WORKDIR /projects
RUN go clean --modcache
WORKDIR /projects/app
RUN go build -o apps

FROM alpine:3.14
WORKDIR /root/
RUN mkdir /config
WORKDIR /root/config
COPY --from=builder /projects/config/.env .
WORKDIR /root/
COPY --from=builder /projects/app/apps .
EXPOSE 8005
CMD ["./apps"]

