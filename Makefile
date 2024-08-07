all : test

test :
	@go test -v ./...

cover :
	@go test -coverprofile=unit.coverage.out ./...
	@curl -o get.sh -Ls https://coverage.codacy.com/get.sh
	@bash get.sh report \
    --force-coverage-parser go -r unit.coverage.out
