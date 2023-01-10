build:
	docker build -t thiagozs/octopus-web .

run:
	docker run --name octopus-web -d -p 8080:8080 thiagozs/octopus-web:latest

clean:
	docker stop octopus-web
	docker rm octopus-web