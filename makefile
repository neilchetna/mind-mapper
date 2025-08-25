.PHONEY: dev

# [BACKEND]
start-db:
	docker compose up -d 

start-go:
	cd api && air | sed 's/^/[GO] /'

migrations:
	cd api && go run cmd/migrations/main.go

# [WEBAPP]
start-next:
	cd webapp && pnpm run dev 2>&1 | sed 's/^/[SVELTE] /'

dev: start-db
	@$(MAKE) -j2 start-go start-next