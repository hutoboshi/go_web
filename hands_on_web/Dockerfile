FROM golang:1.23.2-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o -trimpath -ldflags "-w -s" -o app

# -------------------------

  #デプロイ用コンテナ
  FROM debian:bullseye-slim as deploy

  RUN apt-get update

  COPY --from=deploy-builder /app/app .

  CMD ["./app"]

# -------------------------

  #ローカル開発環境で利用するホットリロード環境
FROM golang:1.23.2 as dev
WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest
# RUN go install github.com/air-verse/air@v1.61.1
RUN go install github.com/air-verse/air@latest

CMD ["air"]
