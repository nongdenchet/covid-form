FROM golang:1.14-alpine3.11

# Copy source
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ADD . /usr/src/app
WORKDIR /usr/src/app

# Environment
ENV ROOT_DIR /usr/src/app

# Build and run
RUN go build -o main .
CMD ["/app/main"]
