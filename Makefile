dev:
	make build
	make run
build:
	rm ./bin/advent-of-code
	go build -o ./bin/advent-of-code ./cmd/main.go 
run:
	./bin/advent-of-code
