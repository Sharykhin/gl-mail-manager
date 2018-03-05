ENV_MG_API_KEY="key-3c03f2d382df4184ecd9ee60e58fbff7"
ENV_MG_DOMAIN="sandboxad94a90d88814c648febc87b5b026118.mailgun.org"
ENV_MG_PUBLIC_API_KEY="pubkey-5933f91a99becbdd97ecaec90dda2985"
ENV_MG_URL="https://api.mailgun.net/v3"

serve-dev:
	MG_API_KEY=$(ENV_MG_API_KEY) MG_DOMAIN=$(ENV_MG_DOMAIN) MG_PUBLIC_API_KEY=$(ENV_MG_PUBLIC_API_KEY) MG_URL=$(ENV_MG_URL) AMPQ_ADDRESS="amqp://127.0.0.1:5672" TEST_FAIL=false APP_ENV=dev GRPC_PUBLIC_KEY=server.crt GRPC_SERVER_ADDRESS=localhost:50051 go run main.go

docker-serve: docker-serve-dev

docker-serve-dev:
	MG_API_KEY=$(ENV_MG_API_KEY) MG_DOMAIN=$(ENV_MG_DOMAIN) MG_PUBLIC_API_KEY=$(ENV_MG_PUBLIC_API_KEY) MG_URL=$(ENV_MG_URL) AMPQ_ADDRESS="amqp://gl-mail-manager-rabbitmq:5672" TEST_FAIL=false APP_ENV=dev GRPC_PUBLIC_KEY=server.crt GRPC_SERVER_ADDRESS="gl-mail-grpc-server-golang:50051" go run main.go

docker-serve-test:
	MG_API_KEY=$(ENV_MG_API_KEY) MG_DOMAIN=$(ENV_MG_DOMAIN) MG_PUBLIC_API_KEY=$(ENV_MG_PUBLIC_API_KEY) MG_URL=$(ENV_MG_URL) AMPQ_ADDRESS="amqp://gl-mail-manager-rabbitmq:5672" TEST_FAIL=OK APP_ENV=test GRPC_PUBLIC_KEY=server.crt GRPC_SERVER_ADDRESS="gl-mail-grpc-server-golang:50051" go run main.go

lint:
	gometalinter ./...