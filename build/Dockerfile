###
# Build stage #
###############
FROM golang:1.15 AS builder

## SETUP for git
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPRIVATE=code.it4i.cz

# Copy .gitconfig
COPY gitconfig /root/.gitconfig

## SETUP for the building

# Get the source code
WORKDIR /go/src/code.it4i.cz/lexis/wp8/

COPY . ./SERVICE

# Set the workdir
WORKDIR /go/src/code.it4i.cz/lexis/wp8/SERVICE/server

# Build the service
ARG COMPILATION
RUN go mod download
RUN go build -a -ldflags "-extldflags \"-static\" -X main.version=SERVICE-CICD-$(cat ../VERSION)-${COMPILATION}" -o /SERVICE

###
# Deploy stage #
################
FROM alpine:latest AS deploy

# Update and certs install
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Set the workdir
WORKDIR /

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

# Copying the binary
COPY --from=builder /SERVICE .

# Entrypoint
CMD /wait && /SERVICE -conf=./config
