NAME=EIMS


build:
	go build -v -trimpath -ldflags "-buildid=" ./cmd/eims