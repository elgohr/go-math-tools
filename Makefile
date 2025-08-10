all: # no default build

test:
	go test -race ./...

update-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

lint: update-tools
	golangci-lint run -v ./...
	govulncheck ./...

install-release-tool:
	go install github.com/elgohr/semv@latest

porcelain:
	./scripts/porcelain.sh

new-patch-release: porcelain
	./scripts/release.sh --patch

new-minor-release: porcelain
	./scripts/release.sh --minor

new-major-release: porcelain
	./scripts/release.sh --major
