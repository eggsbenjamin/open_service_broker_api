.SILENT:

clean:
	rm -rf ./bin

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
	go test ./... -count=1 -tags=integration

service_test:
	echo "\nsystem tests\n"
	DB_HOST=localhost \
	DB_PORT=32768 \
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
