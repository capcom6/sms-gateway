# Building the binary of the App
FROM golang:1.21-alpine AS build

ARG APP
WORKDIR /go/src

# Copy go.mod and go.sum
COPY go.* ./

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Copy all the Code and stuff to compile everything
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app ./cmd/${APP}/main.go

# Build MkDocs
FROM squidfunk/mkdocs-material:9.5.15 AS mkdocs

WORKDIR /docs

RUN pip install --no-cache-dir mkdocs-render-swagger-plugin

COPY web/mkdocs/fonts ./.cache/plugin/social/fonts/Roboto
COPY web/mkdocs .

RUN mkdocs build


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:3 as prod

WORKDIR /app

RUN apk add --no-cache tzdata

COPY scripts/docker-entrypoint.sh /docker-entrypoint.sh

COPY --from=mkdocs /docs/site /app/static
COPY --from=build /go/src/app /app

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

USER guest

ENTRYPOINT ["/docker-entrypoint.sh"]
