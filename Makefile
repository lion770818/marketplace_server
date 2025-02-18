# normal
SERVER_NAME := marketplace_server
# 取得 git remote url 的專案名稱 來當作編譯的 server name
# SERVER_NAME := $(strip $(shell basename -s .git `git config --get remote.origin.url`))
# 輸入參數 手動輸入版本號   例如: make ver=v1.0.0
VERSION     	:= $(ver)
# 取得 git commit id 傳入程式碼中 當作版本查詢的依據
COMMIT_ID 		:= $(shell git rev-parse HEAD)
# 每個服務都有不同的 port
SERVER_PORT 	:= 8888

SRC := ./cmd/${SERVER_NAME}/main.go
BIN_DIR := ./bin
GO_FLAGS := -ldflags="-s -w" # 壓縮二進位檔案大小


# 編譯的主程式
SRC := ./cmd/${SERVER_NAME}/main.go	
# 編譯後的二進位檔案目錄								
BIN_DIR := ./bin
# 壓縮二進位檔案大小 ( -s无法用 gdb 进行调试 -w无法用 dlv（Delve）等调试器调试 )							
GO_FLAGS := -ldflags="-s -w"
# 編譯時間 時分秒 						
#BUILD_TIME=$(shell date +"%F %T")
# 編譯時間 年月日 (太精確每次編譯的md5值都會不一樣, 所以配合版本更新作業,只取到day)						
BUILD_TIME=$(shell date +"%Y-%m-%d")

# 編譯的目標平台 mac GOOS=darwin GOARCH=arm64, linux GOOS=linux GOARCH=amd64
GOOS        := darwin
GOARCH      := amd64
CGO_ENABLED := 0


# dockerhub or harbor
REGISTRY := 192.168.100.69:9003/test
# docker image name
IMAGE_NAME := $(REGISTRY)/$(SERVER_NAME):$(VERSION)
# docker build platform  多重平台編譯用逗點隔開 linux/amd64,linux/arm64
PLATFORM := linux/arm64
# docker build file
DOCKERFILE := ./docker/marketplace_server/Dockerfile 
# docker build source
CONTEXT := ../

# docker compose
COMPOSE_FILE := ./docker/marketplace_server/docker-compose-only-makefile.yml
ENV_FILE := ./.env

.PHONY: all clean build run docker-build docker-push gen-env

all: clean build docker-build docker-push gen-env


# CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -ldflags '-X "main.BUILD_TIME='"$BUILD_TIME"'" -X "main.version='"$Version"'" -X "main.commitId='"$COMMIT_ID"'"' -o ./build/$ServerName ./cmd/main/main.go
build:
	@echo "Building=$(SERVER_NAME) env=$(ENV) Version=$(VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) git commitId= ${COMMIT_ID}  BUILD_TIME=$(BUILD_TIME) "
	@mkdir -p $(BIN_DIR)

	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=" -s -w -X 'main.version=$(VERSION)' -X 'main.commitId=$(COMMIT_ID)' -X 'main.buildTime=$(BUILD_TIME)' " -o $(BIN_DIR)/$(SERVER_NAME) $(SRC)
	@echo "Build successful! Binary located at $(BIN_DIR)/$(SERVER_NAME)"	

run: build
	@$(BIN_DIR)/$(SERVER_NAME)

clean:
	@echo "Cleaning up..."
	@rm -f $(BIN_DIR)/$(SERVER_NAME)
	@echo "Clean complete."

docker-build:
	@echo "Building Docker image $(IMAGE_NAME)..."
	@docker build -t $(IMAGE_NAME) --platform $(PLATFORM)  -f $(DOCKERFILE) $(CONTEXT)
	@echo "Docker image $(IMAGE_NAME) built successfully."

docker-push:
	@echo "Pushing Docker image $(IMAGE_NAME)..."
	@docker push $(IMAGE_NAME)
	@echo "Docker image $(IMAGE_NAME) pushed successfully."

gen-env:
	@echo "Generating .env file Docker image $(IMAGE_NAME)"
	@echo "IMAGE_NAME=$(IMAGE_NAME)" > .env
	@echo "env_flag=dev11" >> .env
	@echo "web_mode=debug" >> .env
	@echo "web_port=$(SERVER_PORT)" >> .env
	@echo "mysql_log_mode=dev" >> .env
	@echo "mysql_host=127.0.0.1" >> .env
	@echo "mysql_port=3306" >> .env
	@echo "mysql_database=test" >> .env
	@echo "mysql_user=leo" >> .env
	@echo "mysql_password=aa1234" >> .env
	@echo "auth_active=jwt" >> .env
	@echo "auth_expireTime=3600" >> .env
	@echo "auth_privateKey=1234" >> .env
	@echo "redis_host=127.0.0.1" >> .env
	@echo "redis_port=6379" >> .env
	@echo "redis_password=" >> .env
	@echo "rabbitmq_host=127.0.0.1" >> .env
	@echo "rabbitmq_port=5672" >> .env
	@echo "rabbitmq_user=admin" >> .env
	@echo "rabbitmq_password=aa1234" >> .env
	@echo "rabbitmq_connectNum=5" >> .env
	@echo "rabbitmq_channelNum=10" >> .env
	@echo "log_env=dev" >> .env
	@echo "log_path=./log/marketplace_server.log" >> .env
	@echo "log_encoding=console" >> .env
	@echo "log_max_size=100" >> .env
	@echo "log_max_age=7" >> .env
	@echo "log_max_backups=5" >> .env
	@echo ".env file generated successfully."

# docker-compose -f docker-compose.yml --env-file .env up
docker-run:
	@echo "Running Docker Compose service $(SERVER_NAME)..."
	@docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) up
	@echo "Docker Compose service $(SERVER_NAME) has been started."		