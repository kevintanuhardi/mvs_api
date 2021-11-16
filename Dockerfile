# generate proto file
FROM namely/protoc-all:1.30_0 as proto_builder
COPY . /go/src/gitlab.warungpintar.co/sales-platform/brook
WORKDIR /go/src/gitlab.warungpintar.co/sales-platform/brook

# Generate Proto
RUN apk add --update make
RUN protoc --version
RUN make proto

# Stage build
FROM golang:1.16.5-alpine AS builder
ARG SSH_PRIVATE_KEY
ENV GO111MODULE=on
ENV GOPRIVATE=gitlab.warungpintar.co
ENV BUILDDIR /go/src/gitlab.warungpintar.co/sales-platform/brook
COPY --from=proto_builder $BUILDDIR /go/src/gitlab.warungpintar.co/sales-platform/brook
RUN apk add --update gcc openssh git bash libc-dev ca-certificates make g++

COPY . /go/src/gitlab.warungpintar.co/sales-platform/brook
WORKDIR /go/src/gitlab.warungpintar.co/sales-platform/brook

RUN mkdir -p /root/.ssh/ \
    && touch /root/.ssh/config

RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa \
    && chmod 600 /root/.ssh/id_rsa \
    && echo "IdentityFile /root/.ssh/id_rsa" >> /root/.ssh/config \
    && echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > /root/.ssh/config \
    && git config --global url."git@gitlab.warungpintar.co:".insteadOf "https://gitlab.warungpintar.co/"

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags="-s -w" \
        -o ./brook ./main.go

# Stage Runtime Applications
FROM alpine:latest

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# add ca-certificates
RUN apk add --no-cache ca-certificates

ENV BUILDDIR /go/src/gitlab.warungpintar.co/sales-platform/brook

# Setting folder workdir
WORKDIR /opt/brook
RUN mkdir config

# Copy Data App
COPY --from=builder $BUILDDIR/brook .

EXPOSE 8009 7070

ENTRYPOINT ["./brook"]