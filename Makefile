.PHONY: day1 day1_1 day1_2
day1_1:
	@mkdir -p day1
	go build -o day1/part1 ./cmd/1_1/main.go
day1_2:
	@mkdir -p day1
	go build -o day1/part2 ./cmd/1_2/main.go
day1: day1_1 day1_2
