mocks:
	go generate ./...

run:
	go run ./cmd/api

run-dev:
	go run ./cmd/api -env=DEV

run-uat:
	go run ./cmd/api -env=UAT

run-prod:
	go run ./cmd/api -env=PROD

