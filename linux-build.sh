GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=1.0" && mkdir -p build/linux64 && mv trec build/linux64
