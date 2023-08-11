# Building the binary of the App
FROM golang:1.20-alpine AS build

ARG SSH_PRV_KEY
ARG APP
WORKDIR /go/src

RUN apk add --no-cache git=2.38.5-r0 openssh-client-default=9.1_p1-r4

RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    echo "StrictHostKeyChecking no " > /root/.ssh/config && ls /root/.ssh/config

RUN echo "$SSH_PRV_KEY" > /root/.ssh/id_rsa \
    && chmod 600 /root/.ssh/id_rsa

RUN echo '[url "ssh://git@bitbucket.org/"]' >> ~/.gitconfig \
    && echo '        insteadOf = https://bitbucket.org/' >> ~/.gitconfig \
    && go env -w GOPRIVATE="bitbucket.org/soft-c/*"

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy go.mod and go.sum
COPY go.* ./

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Copy all the Code and stuff to compile everything
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app ./cmd/${APP}/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:3 as prod

WORKDIR /app

ENV GOOSE_DRIVER=mysql
ENV GOOSE_MIGRATION_DIR=/app/migrations

COPY --from=build /go/bin/goose /app/goose

COPY scripts/docker-entrypoint.sh /docker-entrypoint.sh

# `boilerplate` should be replaced here as well
COPY db/migrations /app/migrations
COPY --from=build /go/src/app /app

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

USER guest

ENTRYPOINT ["/docker-entrypoint.sh"]
