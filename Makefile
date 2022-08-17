generate-mocks:
	mockgen -destination=internal/mocks/message/mock_message_service.go -package mocks github.com/eneskzlcn/catbyte-test-task/internal/message MessageService
	mockgen -destination=internal/mocks/message/mock_rabbitmq_client.go -package mocks github.com/eneskzlcn/catbyte-test-task/internal/message RabbitMQClient
	mockgen -destination=internal/mocks/processor/mock_rabbitmq_client.go -package mocks github.com/eneskzlcn/catbyte-test-task/internal/processor RabbitMQClient
	mockgen -destination=internal/mocks/processor/mock_redis_client.go -package mocks github.com/eneskzlcn/catbyte-test-task/internal/processor RedisClient
	mockgen -destination=internal/mocks/reporting/mock_redis_client.go -package mocks github.com/eneskzlcn/catbyte-test-task/internal/reporting RedisClient