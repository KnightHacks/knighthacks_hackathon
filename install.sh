export GOBIN="$PWD"/bin
export PATH="$GOBIN":"$PATH"
go get github.com/99designs/gqlgen
go install github.com/99designs/gqlgen
