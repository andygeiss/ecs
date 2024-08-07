all : test

test :
	@go test -v github.com/andygeiss/ecs/core/...

coverprofile :
	@go test -coverprofile=unit.coverage.out github.com/andygeiss/ecs/core/...
	@curl -o get.sh -Ls https://coverage.codacy.com/get.sh
	@CODACY_PROJECT_TOKEN=a1929d8d601b4b97a8cef7233ff368ff bash get.sh report \
    --force-coverage-parser go -r unit.coverage.out
