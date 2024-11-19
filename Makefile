
build:
	go mod tidy && go build -o ./bin/mrt2mmdb main.go

run: build
	./bin/mrt2mmdb /data/data.mrt out.mmdb