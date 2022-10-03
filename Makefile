.PHONY: build

all:
	build

migration:
	@cd $(PWD)/src/migrations/goose/ && read -p "Enter migration name: " migration_name; \
	goose create $${migration_name} sql

migrate:
	@bash $(PWD)/scripts/migration.sh

gen_mock_repo:
	@cd src && mockery --dir=repositories/ --output mocks/repositories --case underscore --all
