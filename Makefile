.PHONY: api-gateway

# 启动单个服务
start-user-api:
	@echo "启动 user-api..."
	@cd api-gateway && go run main.go &


# 启动所有服务
start-all: api-gateway
	@echo "所有服务已启动！"



