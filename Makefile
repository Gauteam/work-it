build:
	@docker build . -t work-it:latest

docker-run:
	@docker run work-it:latest

run:
	@go run .