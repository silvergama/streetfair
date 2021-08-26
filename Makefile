GOPACKAGES=$$(go list ./... )

ifdef file
FILE=$(file)
endif

# =========== Dependencies =============
test/deps/up:
	docker-compose up -d
	@docker-compose run wait
	@make test/deps/migrate

test/deps/migrate:
	docker run --rm --network=host -v "$(PWD)/migrations:/flyway/sql:ro"  \
			boxfuse/flyway:5.2.4-alpine \
			-driver="org.postgresql.Driver" \
			-user="street_fair" \
			-schemas="street_fair" \
			-password="123456" \
			-url="jdbc:postgresql://localhost:5432/street_fair" \
			migrate

test/deps/down:
	docker-compose down

# =========== Test =============
test-local:
	@make test/deps/down
	@make test/deps/up
	go test -failfast -count=1 -v $(GOPACKAGES)

# =========== Swagger =============
swagger:
	swagger generate spec -o ./docs/swagger.json --scan-models

swagger/api: swagger
	swagger serve -F=swagger ./docs/swagger.json

# =========== Coverage =============
check/swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

clean-coverage:
	mkdir -p .cover
	rm -rf .cover/*

coverage: clean-coverage
	go test -tags="all" -covermode="count" -coverprofile=".cover/cover.out" $(GOPACKAGES)

coverage-html: coverage
	go tool cover -html=.cover/cover.out

docker/coverage:
	-docker rm -f coverage
	docker run --net=host --name=coverage --entrypoint /bin/sh silvergama/street_fair:test -c "make coverage"

# =========== Docker =============
docker/build:
	docker build -t silvergama/street_fair:test -f Dockerfile.build .

docker/image:
	docker build -t silvergama/street_fair .

# =========== App =============
clean:
	rm -rf vendor/

install: clean
	go mod vendor && go mod tidy

help:
	go run main.go help

run/api:
	go run main.go api

run/import:
	go run main.go import $(FILE)

build:
	go build -v --ldflags="-s"