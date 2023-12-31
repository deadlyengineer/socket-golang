docsurl := "http://localhost:8080/pkg/github.com/deadlyengineer/socket-golang"

_default:
  @just --list --list-prefix '  > '

# host docs with godoc and open documentation page
docs:
  sleep 2 && if [ {{os()}} == "macos" ]; then open {{docsurl}}; else xdg-open {{docsurl}}; fi &
  godoc -http=localhost:8080

# test all packages
test:
  go test -v ./...

# test all packages, generate coverfile, and open coverfile in browser
test-cover:
  go test -v ./... -coverprofile=c.out; go tool cover -html=c.out

# remove residual test files
clean:
  rm -f c.out
