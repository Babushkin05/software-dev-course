FROM golang:1.24 as builder

RUN apk add --no-cache git make protobuf

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make proto

RUN go build -o analysis-service ./cmd/main.go


FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app


COPY --from=builder /app/analysis-service .
COPY config/local.yaml ./config/local.yaml

EXPOSE 50052

CMD ["./analysis-service", "--config=config/local.yaml"]
