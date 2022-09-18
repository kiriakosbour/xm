FROM golang:1.17-alpine AS builder
# Metadata for the build, compile only for linux target, smaller binary size

ENV GOOS linux
ENV GOARCH amd64
# Turn off CGO since that can result in dynamic links to libc/libmusl.
ENV CGO_ENABLED 0
WORKDIR /xm
#add git needed bellow
RUN apk add git

# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# Copy and build the app
COPY . .
RUN go build -ldflags="-w -s" -o xm-app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /xm/xm-app .
EXPOSE 8080
CMD ["./xm-app"]