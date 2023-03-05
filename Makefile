hellserver: main.go
	go build -o $@ $<

test: 
	go test ./...
