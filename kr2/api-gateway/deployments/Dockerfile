FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make proto

RUN go build -o /api-gateway ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /api-gateway /app/api-gateway
COPY --from=builder /app/config/local.yaml /app/config/local.yaml

ENTRYPOINT ["/app/api-gateway"]
CMD ["--config=config/local.yaml"]
