GOPACKAGES=$$(go list ./... )

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

swagger:
	swagger generate spec -o ./docs/swagger.json --scan-models

swagger/api: swagger
	swagger serve -F=swagger ./docs/swagger.json
