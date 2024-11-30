package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log/slog"
	"regexp"
	"strconv"
)

//go:embed input.txt
var file []byte

func main() {
	scanner := bufio.NewScanner(bytes.NewReader(file))

	var total = 0
	for scanner.Scan() {
		line := scanner.Text()
		re, _ := regexp.Compile(`([^\D+])`)
		if ok := re.MatchString(line); !ok {
			slog.Error("must match all lines", line)
		}
		matches := re.FindAllString(line, -1)
		num := ""
		if len(matches) >= 2 {
			num += matches[0]
			num += matches[len(matches)-1]
		} else {
			num += matches[0]
			num += matches[0]
		}
		integer, _ := strconv.Atoi(num)
		total += integer
		slog.Debug(strconv.Itoa(integer))
	}

	slog.Info(strconv.Itoa(total))
}
