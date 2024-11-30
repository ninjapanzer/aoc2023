package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log"
	"log/slog"
	"regexp"
	"strconv"
)

//go:embed input.txt
var file []byte

var numberRegex = `([1-9]|nine|eight|seven|six|five|four|three|two|one)`
var firstRe = regexp.MustCompile(numberRegex)
var reversedNumberRegex = `([1-9]|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno)`
var lastRe = regexp.MustCompile(reversedNumberRegex)

// Since there was some generally undefined rules about how word combinations should be addressed
// we will follow a last first solution here by reversing the lookup process.
// While in Part 1 we were able to trust the numbers once we have to defend against words that can overlap
// like eightwo or twone
// It's hard to interpret if this should be 8 2 or 2 or 1 and golang doesn't support regex positive lookeahead
// Instead we examine each string forwards and backwards selecting the first match each direction.
// This will be the reversed and uses a reversed regex
func main() {
	scanner := bufio.NewScanner(bytes.NewReader(file))

	var total = 0
	for scanner.Scan() {
		num := ""
		line := scanner.Text()
		num += convertWordtoInt(
			firstDigit(line, firstRe))
		num += convertWordtoInt(
			reverse(
				firstDigit(
					reverse(line), lastRe)))

		integer, err := strconv.Atoi(num)
		if err != nil {
			slog.Error("couldn't convert", err)
		}
		total += integer
		slog.Debug(strconv.Itoa(integer))
	}

	log.Println(total)
}

func firstDigit(word string, re *regexp.Regexp) string {
	return re.FindString(word)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func convertWordtoInt(word string) string {
	if len(word) > 1 {
		switch word {
		case "one":
			return "1"
		case "two":
			return "2"
		case "three":
			return "3"
		case "four":
			return "4"
		case "five":
			return "5"
		case "six":
			return "6"
		case "seven":
			return "7"
		case "eight":
			return "8"
		case "nine":
			return "9"
		}
	} else {
		return word
	}
	return ""
}
