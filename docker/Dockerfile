# -- multistage docker build: stage #1: build stage
FROM golang:1.18-alpine AS build

RUN mkdir -p /go/src/github.com/karlsen-network/dnsseeder

WORKDIR /go/src/github.com/karlsen-network/dnsseeder

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

# For development it's useful to have karlsend from filesystem, but for deployment
# we should use the one in github
RUN go mod edit -dropreplace github.com/karlsen-network/karlsend
ARG KARLSEND_VERSION
ARG KARLSEND_REPOSITORY='github.com/karlsen-network/karlsend'
RUN if [ -n "${KARLSEND_VERSION}" ]; then \
        # use replace instead of require - to propagate into nested dependancies \
        go mod edit -replace "github.com/karlsen-network/karlsend=$KARLSEND_REPOSITORY@$KARLSEND_VERSION"; \
    fi

RUN go mod download

RUN cp go.mod go.mod.bu
RUN cp go.sum go.sum.bu

COPY . .

# Restore go.mod and go.sum because `COPY . .` overwrote them
RUN mv go.mod.bu go.mod
RUN mv go.sum.bu go.sum

RUN go mod tidy

RUN GOOS=linux go build -a -installsuffix cgo -o dnsseeder .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache tini

COPY --from=build /go/src/github.com/karlsen-network/dnsseeder/ /app/

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/dnsseeder"]
