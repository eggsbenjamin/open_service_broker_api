.SILENT:

all: lint test build

setup_dev: setup_db create_schema seed_test_data deps run_dev

build: clean
	GOOS=linux $GOARCH=amd64 CGO_ENABLED=0 go build -o bin/main main.go

docker_build: build
	docker build .

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure -v

clean:
	rm -rf ./bin

run_dev:
	echo "\nrunning service with dev config\n"
	DB_HOST=localhost \
	DB_PORT=1234 \
	DB_USER=postgres \
	DB_PWD=postgres \
	DB_NAME=service_catalog \
	go run main.go

mocks:
	echo "\ncreating mocks\n"
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
	go generate ./...

setup_db:
	echo "\nsetting up db\n"
	docker-compose rm -fs postgresql
	docker-compose up -d postgresql
	sleep 5

create_schema:
	echo "\ncreating schema\n"
	psql -h localhost -p 1234 -U postgres -d service_catalog -f ./sql/1_service.sql
	psql -h localhost -p 1234 -U postgres -d service_catalog -f ./sql/2_service_plan.sql
	psql -h localhost -p 1234 -U postgres -d service_catalog -f ./sql/3_service_instance.sql

seed_test_data:
	echo "\nseeding test data\n"
	psql -h localhost -p 1234 -U postgres -d service_catalog -f ./servicetest/testdata/catalog.sql

unit_test:
	echo "\nunit tests\n"
	go test ./... -tags=unit

integration_test: setup_db create_schema
	echo "\nintegration tests\n"
	go test ./... -count=1 -tags=integration

service_test: setup_db create_schema
	echo "\nsystem tests\n"
	DB_HOST=localhost \
	DB_PORT=1234 \
	DB_USER=postgres \
	DB_PWD=postgres \
	DB_NAME=service_catalog \
	go test ./... -count=1 -tags=service

test: unit_test integration_test service_test

lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... --vendor --skip=vendor --exclude=\.*_mock\.*\.go --exclude=vendor\.* --cyclo-over=20 --deadline=10m --disable-all \
	--enable=errcheck \
	--enable=vet \
	--enable=deadcode \
	--enable=gocyclo \
	--enable=varcheck \
	--enable=structcheck \
	--enable=maligned \
	--enable=vetshadow \
	--enable=ineffassign \
	--enable=interfacer \
	--enable=unconvert \
	--enable=goconst \
	--enable=gosimple \
	--enable=staticcheck \
	--enable=gosec \
	--enable=safesql
