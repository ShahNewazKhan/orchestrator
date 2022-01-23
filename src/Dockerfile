############################
# STEP 1 build executable binary
############################
# Pinnin to alpine3.12 to resolve docker build DNS error 
# https://github.com/gliderlabs/docker-alpine/issues/386
FROM golang:alpine3.12 AS builder

# Install git for fetching go dependencies.
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/orc/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/orchestrator

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/orchestrator /go/bin/orchestrator

EXPOSE 3000
# Run the orc binary.
ENTRYPOINT ["/go/bin/orchestrator"]
