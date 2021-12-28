PROJECTNAME := $(shell basename "$(PWD)")
STIME 		:= $(shell date +%s)
pkgs  		:= $(shell go list ./... | grep -vE '(vendor|mock)')
## lint: lint program
.PHONY: lint
lint:
	@echo "  >  Linting Program..."
	golangci-lint run --issues-exit-code 0 --timeout 10m

## test: test program
.PHONY: test
test: proto
	@echo "  >  Testing Program..."
	go test -race ./...

## test: test mysql service
.PHONY: test-service
test-service: proto
	@echo "  >  Testing Program..."
	go test -v ./domain/user/repository/mysql

## coverage: coverage program
.PHONY: coverage
coverage: ## Generate global code coverage report
	chmod u+x coverage.sh
	./coverage.sh;

## coverage: coverage program
.PHONY: coverhtml
coverhtml: ## Generate global code coverage report in HTML
	chmod u+x coverage.sh
	./coverage.sh html;

.PHONY: build
build: proto
	@echo ">  Building Program..."
	GOPRIVATE=gitlab.warungpintar.co go build -ldflags="-s -w" -o bin/${PROJECTNAME} main.go; 
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## rest: start without docker
.PHONY: rest
rest: build
	@echo "  >  Starting Program..."
	./bin/${PROJECTNAME} rest
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## grpc: start without docker
.PHONY: grpc
grpc: build
	@echo "  >  Starting Program..."
	./bin/${PROJECTNAME} grpc
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## build-docker 
.PHONY: build-docker
build-docker:
	@echo ">  Building Program..."
	docker build -t sales-platform/brook .
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## docker-rest: start with docker
.PHONY: docker-rest
docker-rest: build-docker
	@echo "  >  Starting Program..."
	docker run -d \
		--name brook \
	    -p 8009:8009 \
		--mount type=bind,source="$(CURDIR)"/config,target=/config \
		--network brook_default \
		--link brook_db_1:brook_db_1 \
		--rm sales-platform/brook rest
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

mock:
	mockgen --source=domain/domain.go --destination=domain/mocks/repository.go --package mocks
	mockgen --source=domain/user/repository/repository.go --destination=domain/user/repository/mocks/repository.go --package mocks
	mockgen --source=domain/user/usecase/user.go --destination=domain/user/usecase/mocks/service.go --package mocks
	mockgen --source=pkg/webservice/server.type.go --destination=pkg/webservice/mocks/server.type.go --package mocks
	mockgen --source=pkg/router/router.go --destination=pkg/router/mocks/router.go --package mocks


.PHONY: docker-grpc
docker-grpc: build-docker
	@echo "  >  Starting Program..."
	docker run -d \
		--name brook \
	    -p 8009:8009 \
		--mount type=bind,source="$(CURDIR)"/config,target=/config \
		--network brook_default \
		--link brook_db_1:brook_db_1 \
		--rm sales-platform/brook grpc
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

.PHONY: start-dev
start-dev: ## Starting environments for development in docker compose
	@echo " > Start Development ENV..."
	docker-compose up --build -d
	@echo " > Done Start Development ENV..."
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

.PHONY: stop-dev
stop-dev: ## Stoping environments for development in docker compose
	@echo " > Stop Development ENV..."
	docker-compose down
	@echo " > Done Stop Development ENV..."
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

.PHONY: check-dev
check-dev: ## Checking our local env is ready
	docker-compose ps

.PHONY: migrate-up
migrate-up: ## Migrate data to development database
	goose --dir=db/migrations mysql "brook:brook@tcp(localhost:3307)/brook?parseTime=true&timeout=5s" up

.PHONY: migrate-down
migrate-down: ## Reset data to development database
	goose --dir=db/migrations mysql "brook:brook@tcp(localhost:3307)/brook?parseTime=true&timeout=5s" reset

.PHONY: seed
seed: build ## Seeding Data
	@echo "  >  Starting Seeding Data..."
	./bin/${PROJECTNAME} seed
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

.PHONY: proto
proto: 
	@echo "  >  Start Generate Proto..."
	protoc brook.proto --proto_path=proto/brook/ --go_out=plugins=grpc:proto

.PHONY: sonarqube-server
sonarqube-server:
	@docker create --name=sonarqube-server -p 9000:9000 sonarqube:8.5-community

.PHONY: run-sonarqube
run-sonarqube:
	@docker start sonarqube-server

.PHONY: unit-test
unit-test: proto
	@rm -f cover.out
	@echo "${STIME} running tests"
	@go test -race -coverprofile=cover.out $(pkgs)

.PHONY: sonar-scan
sonar-scan:
	@echo "${STIME} Scan project"
	@sonar-scanner -Dproject.settings=sonar-project.properties

.PHONY: sonar-coverage
sonar-coverage: unit-test sonar-scan
