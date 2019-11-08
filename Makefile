# Build
producer: 
		@go run cmd/producer/main.go

consumer:
		@go run cmd/consumer/main.go

.PHONY: producer consumer