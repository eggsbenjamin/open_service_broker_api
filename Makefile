.SILENT:

mocks:
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	go generate ./...

setup_db:
	docker-compose rm -fs postgresql
	docker-compose up -d postgresql

unit_test:
	echo "\nunit tests\n"
	go test ./... -tags=unit

integration_test:
	echo "\nintegration tests\n"
	go test ./... -tags=integration

test: unit_test integration_test
