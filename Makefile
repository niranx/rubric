test:
	go test ./internal/...

rubric:
	go build -o rubric ./cmd/rubric/main.go

clean:
	rm -f rubric

tidy:
	go mod tidy

.PHONY: clean crypt test install tidy
