package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	Instance int
	Blue     int
	Red      int
	Green    int
}

var outerRegex = `^\w+.(\d+):\s(.*)$`
var outerRe = regexp.MustCompile(outerRegex)
var bagRegex = `(\d+.\w+)`
var bagRe = regexp.MustCompile(bagRegex)

//go:embed input.txt
var file []byte
var targetBag = Bag{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func main() {
	//slog.SetLogLoggerLevel(slog.LevelDebug)
	scanner := bufio.NewScanner(bytes.NewReader(file))

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := extractGame(line)
		slog.Debug("game", game)
		fits := true
		for _, g := range game {
			if !g.fits(targetBag) {
				fits = false
				slog.Debug("Not fit", g.Instance)
			}
		}
		if fits {
			total += game[0].Instance
		}
	}
	slog.Info(strconv.Itoa(total))
}

func (g *Game) fits(bag Bag) bool {
	red := bag.Red >= g.Red
	green := bag.Green >= g.Green
	blue := bag.Blue >= g.Blue

	return red && green && blue
}

func extractGame(line string) []Game {
	outerMatch := outerRe.FindStringSubmatch(line)
	no, _ := strconv.Atoi(outerMatch[1])
	gamePhase := strings.Split(outerMatch[2], ";")
	var games []Game

	for _, phase := range gamePhase {
		innerMatch := bagRe.FindAllString(phase, -1)
		game := Game{
			Instance: no,
		}
		game.sumColors(innerMatch)
		games = append(games, game)
	}
	return games
}

func (g *Game) sumColors(plays []string) {
	slog.Debug("plays", plays)
	for _, play := range plays {
		color := strings.Split(play, " ")
		no, _ := strconv.Atoi(color[0])

		switch v := strings.ToLower(color[1]); v {
		case "red":
			g.Red += no
		case "green":
			g.Green += no
		case "blue":
			g.Blue += no
		default:
			slog.Warn("should have matched a color", v)
		}
	}
}
