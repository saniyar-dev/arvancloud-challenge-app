.PHONY: build docker-build run docker-run

build:
	go build -o arvancloud-challenge-app cmd/main.go

docker-build:
	docker build -t arvancloud-challenge-app .

run:
	./arvancloud-challenge-app

docker-run:
	docker run -p 8080:8080 arvancloud-challenge-app

