rest-up:
	docker compose up -d
rest-stop:
	docker compose stop
rest-down:
	docker compose down
rest-init:
	go mod vendor
	echo "Project initiated"