ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION} AS builder
LABEL version="0.1.0"
LABEL org.opencontainers.image.description "A tiny tool to keep track of mobile devices battery"


RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /src

COPY . .
RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -tags 'osusergo netgo static_build' \
  -o /bin/devices \
  ./main.go


RUN upx -q -9 /bin/devices

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/devices /bin/devices

ENTRYPOINT ["/bin/devices"]
