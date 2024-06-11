ARG GO_VERSION=1.22.2

FROM registry.docker.com/library/golang:$GO_VERSION-alpine as base

# app lives here
WORKDIR /app


# Throw-away build stage to reduce size of final image
FROM base as build

# Install packages needed to build
RUN apk update -qq && \
    apk add --no-cache git

COPY . .

RUN go mod tidy

RUN go build -o api cmd/main.go

ENTRYPOINT ["/app/api"]

CMD ["app/api"]


MAINTAINER RyanB1303 <ryan.brilliant.nirwana@gmail.com>
