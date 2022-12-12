

# migrateion tool

copy .env, change db_Host to localhost

go run main.go -e .env.migrate -m down


go run main.go -u 