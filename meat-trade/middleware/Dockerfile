FROM golang:latest

RUN go get github.com/gorilla/mux && \
    go get github.com/hyperledger/fabric-sdk-go/ && \
    go get github.com/pkg/errors && \
    go get github.com/cloudflare/cfssl/api && \
    go get github.com/cloudflare/cfssl/csr && \
    go get github.com/cloudflare/cfssl/helpers && \
    go get github.com/cloudflare/cfssl/log && \
    go get github.com/cloudflare/cfssl/signer && \
    go get github.com/golang/mock/gomock && \
    go get github.com/golang/protobuf/proto && \
    go get github.com/golang/protobuf/ptypes && \
    go get github.com/golang/protobuf/ptypes/timestamp && \
    go get github.com/mitchellh/mapstructure && \
    go get github.com/spf13/cast && \
    go get github.com/spf13/viper && \
    go get github.com/stretchr/testify/assert && \
    go get golang.org/x/crypto/ocsp && \
    go get golang.org/x/crypto/sha3 && \
    go get golang.org/x/net/context && \
    go get google.golang.org/grpc && \
    go get google.golang.org/grpc/codes && \
    go get google.golang.org/grpc/connectivity && \
    go get google.golang.org/grpc/credentials && \
    go get google.golang.org/grpc/keepalive && \
    go get google.golang.org/grpc/peer && \
    go get google.golang.org/grpc/status

WORKDIR /src
