.PHONY: day1 day2 day1_1 day1_2 day2_1 day2_2
day1_1:
	@mkdir -p day1
	go build -o day1/part1 ./cmd/1_1/main.go
day1_2:
	@mkdir -p day1
	go build -o day1/part2 ./cmd/1_2/main.go
day1: day1_1 day1_2
day2_1:
	@mkdir -p day2
	go build -o day1/part1 ./cmd/1_1/main.go
day2_2:
	@mkdir -p day2
	go build -o day1/part2 ./cmd/1_2/main.go
day2: day2_1 day2_2