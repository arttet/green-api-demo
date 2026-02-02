BIN := "bin"
EXT := `go env GOEXE`

SERVER_APP := "green-api-proxy"
MODULE_APP := "cmd/green-api-proxy/main.go"

[doc('Show available recipes and their descriptions')]
help:
    @just --list --unsorted

# ------------------------------------------------------------------------------
# Development
# ------------------------------------------------------------------------------

[group('Development')]
[doc('Format code')]
fmt:
    gofmt -w -s .

[group('Development')]
[doc('Lint code')]
lint:
    golangci-lint run

[group('Development')]
[doc('Build app')]
build:
    go build -o {{ BIN }}/{{ SERVER_APP }}{{ EXT }} {{ MODULE_APP }}

[group('Development')]
[doc('Run app in development mode')]
run:
    go run {{ MODULE_APP }}

# ------------------------------------------------------------------------------
# Web
# ------------------------------------------------------------------------------

[group('Web')]
[doc('Build the site for production')]
[working-directory: 'web']
bundle:
    bun run build

[group('Web')]
[doc('Start development server')]
[working-directory: 'web']
dev:
    bun run dev
