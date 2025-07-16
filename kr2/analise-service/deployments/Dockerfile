FROM golang:1.24-alpine AS builder

RUN apk add --no-cache \
    protobuf \
    make \
    bash \
    git

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make proto

RUN go build -o /analise-service ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /analise-service /analise-service
COPY config /config  

EXPOSE 50052

ENTRYPOINT ["/analise-service", "--config=config/local.yaml"]
