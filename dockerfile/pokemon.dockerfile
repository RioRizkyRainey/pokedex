#############################
# BUILDing the docker image
#############################
FROM golang:1.14-alpine as builder
LABEL maintainer="https://github.com/RioRizkyRainey"

ENV GO111MODULE on
ENV GOOS linux
ENV GOARCH amd64

COPY . pokedex/
WORKDIR pokedex/

# Create appuser.
ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN apk upgrade --update-cache --available && \
    apk add build-base && \
    apk add openssl-dev


RUN rm -vrf /var/cache/apk/*

RUN go mod tidy

#RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /build/parker-api.idn.media parker-api.idn.media/cmd/master/
RUN go build -o /build/pokedex/pokemon cmd/pokemon/main/main.go

#############################
# CREATE the runtime 
#############################
FROM alpine:latest
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /build/pokedex/pokemon /app/
WORKDIR /app

# Use an unprivileged user.
USER appuser:appuser

CMD ["/app/pokemon"]