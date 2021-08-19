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

vendor:
	go mod vendor

install: vendor
