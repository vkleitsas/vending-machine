run:
	docker-compose up --build

test:
	cd ./vending-machine; go test ./...