GOPACKAGES=$$(go list ./... )

ifdef file
FILE=$(file)
endif

# =========== Dependencies =============
deps/up:
	docker-compose up -d
	@docker-compose run wait
	@make deps/migrate

deps/migrate:
	docker run --rm --network=host -v "$(PWD)/migrations:/flyway/sql:ro"  \
			boxfuse/flyway:5.2.4-alpine \
			-driver="org.postgresql.Driver" \
			-user="streetfair" \
			-schemas="streetfair" \
			-password="123456" \
			-url="jdbc:postgresql://localhost:5432/streetfair" \
			migrate

deps/down:
	docker-compose down

# =========== Test =============
test:
	go test -failfast -count=1 -v $(GOPACKAGES)

# =========== Log =============
show/logs:
	tail -f fair.log

# =========== Swagger =============
check/swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check/swagger install
	swagger generate spec -o ./docs/swagger.json --scan-models

swagger/api: swagger
	swagger serve -F=swagger ./docs/swagger.json

# =========== Report =============

clean-report:
	rm -f report.json cover.out govet-report.out

coverage: clean-report
	go test -tags="all" -covermode="count" -coverprofile="cover.out" $(GOPACKAGES) -json > report.json

vet: clean-report
	go vet ./... > govet-report.out

coverage-html: coverage install
	go tool cover -html=cover.out

# =========== Docker =============
docker/build:
	docker build -t silvergama/streetfair:test -f Dockerfile.build .

docker/image:
	docker build -t silvergama/streetfair .

docker/test:
	docker run --rm --net=host --entrypoint /bin/sh silvergama/streetfair:test -c "make test"

docker/coverage:
	-docker rm -f coverage
	docker run --net=host --name=coverage --entrypoint /bin/sh silvergama/streetfair:test -c "make coverage"

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
