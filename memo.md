# export NODE_OPTIONS=--openssl-legacy-provider

npm cache clean --force
docker builder prune
docker compose up -d

go clean -cache
go mod tidy
