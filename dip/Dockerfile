FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# ビルド（Cloud Runに最適化された静的バイナリを生成）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/server /server

COPY .env /.

EXPOSE 8080

CMD ["/server"]
