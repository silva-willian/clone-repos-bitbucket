FROM golang:1.13 as builder

COPY . /go/src/github.com/silva-willian/clone-repos-bitbucket

WORKDIR /go/src/github.com/silva-willian/clone-repos-bitbucket

RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure -vendor-only

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o clone

FROM golang:1.13

WORKDIR /app

COPY --from=builder /go/src/github.com/silva-willian/clone-repos-bitbucket .

CMD ./clone 