GOPACKAGES=$$(go list ./... )

ifdef file
FILE=$(file)
endif

test/deps/up:
	docker-compose up -d
	@docker-compose run wait
	@make test/deps/migrate

test/deps/down:
	docker-compose down

test/deps/migrate:
	docker run --rm --network=host -v "$(PWD)/migrations:/flyway/sql:ro"  \
			boxfuse/flyway:5.2.4-alpine \
			-driver="org.postgresql.Driver" \
			-user="unico" \
			-schemas="unico" \
			-password="123456" \
			-url="jdbc:postgresql://localhost:5432/unico" \
			migrate

test-local:
	@make test/deps/down
	@make test/deps/up
	go test -failfast -count=1 -v $(GOPACKAGES)

check/swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

clean:
	rm -rf vendor/

install: clean
	go mod vendor && go mod tidy

run/api:
	go run main.go api

run/import:
	go run main.go import $(FILE)

swagger:
	swagger generate spec -o ./docs/swagger.json --scan-models

swagger/api: swagger
	swagger serve -F=swagger ./docs/swagger.json

clean-coverage:
	mkdir -p .cover
	rm -rf .cover/*

coverage: clean-coverage
	go test -tags="all" -covermode="count" -coverprofile=".cover/cover.out" $(GOPACKAGES)

coverage-html: coverage
	go tool cover -html=.cover/cover.out

docker/coverage:
	-docker rm -f coverage
	docker run --net=host --name=coverage --entrypoint /bin/sh silvergama/unico:test -c "make coverage"

docker/build:
	docker build -t silvergama/unico:test .