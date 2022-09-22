postgres:
	docker run --name local-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432  -d postgres

redis:
	docker run --name local-redis -p 6379:6379 -d redis

create_db:
	docker exec -it local-postgres createdb --username=root --owner=root portto

migrate_up:
	migrate -path deployment/migration -database "postgresql://root:secret@localhost:5432/portto?sslmode=disable" -verbose up

migrate_down:
	migrate -path deployment/migration -database "postgresql://root:secret@localhost:5432/portto?sslmode=disable" -verbose down

run_server:
	go run ./main.go

get_block_by_number:
	curl https://data-seed-prebsc-2-s3.binance.org:8545/ \
		-X POST \
		-H 'Content-type: application/json' \
		-d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1", true], "id": 1}'

get_transaction_by_hash:
	curl https://data-seed-prebsc-2-s3.binance.org:8545/ \
		-X POST \
		-H 'Content-type: application/json' \
		-d '{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0xa32cf2bb6ec953f2f1d912f23d7be1124fe786e0c0b1b03c6713174a9dba0404"],"id":1}'

get_transaction_receipt:
	curl https://data-seed-prebsc-2-s3.binance.org:8545/ \
		-X POST \
		-H 'Content-type: application/json' \
		-d '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":	["0xa32cf2bb6ec953f2f1d912f23d7be1124fe786e0c0b1b03c6713174a9dba0404"],"id":1}'

.PHONY: postgres create_db migrate_up migrate_down run_server get_block_by_number get_transaction_by_hash get_transaction_receipt