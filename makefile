server:
	go run main.go

gen.pem:
	go run cmd/key-pair/main.go

gen.model:
	go run cmd/gorm-gen/main.go

migration:
	go run cmd/migration/main.go

git.push:
	git add .
	git commit -m "$m"
	git push