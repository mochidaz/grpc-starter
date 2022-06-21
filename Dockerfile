FROM golang:1.18 AS base

# define timezone
ENV TZ Asia/Jakarta

# define work directory
WORKDIR /app

# copy the sourcecode
COPY . .

# installing grpc healthcheck and wait-for
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
RUN WAIT_FOR_VERSION=v2.1.2 && \
    wget -qO/bin/wait-for https://github.com/eficode/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for && \
    chmod +x /bin/wait-for

# install protoc
RUN apt-get update && apt-get install -y protobuf-compiler

# install migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# install all go dependencies
RUN cd ~ && export GO111MODULE=auto && go install github.com/envoyproxy/protoc-gen-validate@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

# generate protocol buffers
RUN make generate-pb

# build beedoor exec
RUN cd /app/cmd/server && go mod vendor && go build -o grpc-starter

# EXPOSE 8080 is the port that the REST API will be exposed on
EXPOSE 8080
# EXPOSE 8080 is the port that the GRPC will be exposed on. But if deployed in cloud run just use the 8080 port
EXPOSE 8081

CMD [ "./cmd/server/grpc-starter" ]