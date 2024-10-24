# ============================================================================
# 👽 vars {{{
-include config.env
export
# }}}
# ============================================================================

# 🛳️ docker {{{
db-pull:
	@docker pull postgres:$(PG_VERSION)

db-run:
	@docker run --name $(DB_DRIVER)-$(DB_NAME) \
			    -e POSTGRES_PASSWORD=$(DB_PASSWORD) \
				-e POSTGRES_USER=$(DB_USER) \
				-p $(DB_PORT):5432 -d postgres:$(DB_VERSION)
db-start:
	@docker start $(DB_DRIVER)-$(DB_NAME)

db-create:
	@docker exec -it $(DB_DRIVER)-$(DB_NAME) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

db-enter:
	@docker exec -it $(DB_DRIVER)-$(DB_NAME) psql -U $(DB_USER)

db-stop:
	@docker stop $(DB_DRIVER)-$(DB_NAME)

db-rm:
	@docker rm $(DB_DRIVER)-$(DB_NAME)

db-drop:
	@docker exec -it $(DB_DRIVER)-$(DB_NAME) dropdb $(DB_NAME)

db-dump:
	pg_dump -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -p $(DB_PORT) -F t -v -x -O -f "dump.tar"

db-restore:
	pg_restore -U $(DB_USER) -d $(DB_NAME) -h $(DB_HOST) -p $(DB_PORT) -F t -v "dump.tar"

db-creds:
	@echo "DB_DRIVER: $(DB_DRIVER)"
	@echo "DB_NAME: $(DB_NAME)"
	@echo "DB_USER: $(DB_USER)"
	@echo "DB_PASSWORD: $(DB_PASSWORD)"
	@echo "DB_HOST: $(DB_HOST)"
	@echo "DB_URL: $(DB_URL)"
	@echo "DB_PORT: $(DB_PORT)"

.PHONY: db-pull db-run db-start db-create db-enter db-stop db-rm db-drop db-dump db-restore db-creds
# }}}

# 🚀 go {{{
r:
	@air -c $(PATH_AIR)

s:
	lightningcss --bundle --targets '>= 0.5%' $(CSS_INPUT) -o $(CSS_OUTPUT)

go-gen: s
	@./bin/env.sh
	@rm -f $(BIN_LOC)$(BIN_NAME)
	@templ generate

go-build: go-gen
	@go build -o=$(BIN_LOC)$(BIN_NAME) $(PATH_GO)

go-run: go-build
	@echo ""
	@$(BIN_LOC)$(BIN_NAME)

.PHONY: go-gen go-build go-run go-tidy r w
# }}}

# 💾 database {{{

# regenerates the sqlc using current schema file
sqlc:
	@sqlc generate -f $(CONFIG_SQLC)

# dumps the schema from database and generates sqlc files from it
dsqlc:
	@export PGPASSWORD=$(DB_PASSWORD); pg_dump -U $(DB_USER) -d $(DB_NAME) -p $(DB_PORT) -h $(DB_HOST) --schema-only --no-owner --no-privileges > $(CONFIG_SQLC)
	@sqlc generate -f $(CONFIG_SQLC)

# go goose---migration commands
CONFIG_GOOSE=GOOSE_MIGRATION_DIR=$(PATH_MIGRATION) \
			 GOOSE_DRIVER=$(DB_DRIVER) \
			 GOOSE_DBSTRING=$(DB_URL)
gi:
	@$(CONFIG_GOOSE) goose create init sql
gs:
	@$(CONFIG_GOOSE) goose status
gU:
	@$(CONFIG_GOOSE) goose up
gu:
	@$(CONFIG_GOOSE) goose up-by-one
gd:
	@$(CONFIG_GOOSE) goose down
gr:
	@$(CONFIG_GOOSE) goose redo
gc:
	@read -p "Enter migration name: " NAME; \
    $(CONFIG_GOOSE) goose create "$$NAME" sql
gv:
	@$(CONFIG_GOOSE) goose validate

.PHONY: sqlc dsqlc gi gs gU gu gd gr gc gv
# }}}
