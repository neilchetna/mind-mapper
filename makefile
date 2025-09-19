.PHONEY: dev

# [BACKEND]
start-db:
	docker compose up -d 

start-go:
	cd api && air -c .air.server.toml | sed 's/^/[GO] /'

start-worker:
	cd api && go run cmd/worker/main.go | sed 's/^/[WORKER] /'

migrations:
	cd api && go run cmd/migrations/main.go

# [WEBAPP]
start-webapp:
	cd webapp && pnpm run dev 2>&1 | sed 's/^/[SVELTE] /'

dev: start-db
	@$(MAKE) -j2 start-go start-webapp