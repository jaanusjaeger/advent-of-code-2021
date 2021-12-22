package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	lc := 0
	players := [2]player{}
	scanner := func(line string) error {
		ss := strings.Split(line, ":")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input: %s", line)
		}
		pos := parseInt(strings.TrimSpace(ss[1]))
		players[lc] = player{pos: pos - 1}
		lc++
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(players)

	result1 := puzzle1(players)
	result2 := puzzle2(players)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(players [2]player) string {
	d := dice{}
	playerIndex := 0
	var loser player

	for {
		// fmt.Println(players)
		player := &players[playerIndex]
		playerIndex = (playerIndex + 1) % len(players)
		m := d.roll() + d.roll() + d.roll()
		// fmt.Println("  --> move:", m)
		player.move(m)
		if player.score >= 1000 {
			loser = players[playerIndex]
			break
		}
	}

	result := fmt.Sprint(loser.score * d.count)

	return result
}

func puzzle2(players [2]player) string {
	universes := map[[2]player]int{}
	wins := [2]int{}

	universes[players] = 1

	playerIndex := 0
	for {
		fmt.Println("puzzle2:", universes)
		universes = copyUniversesForPlayer(universes, playerIndex)

		for ps, count := range universes {
			for i, p := range ps {
				if p.score > 20 {
					wins[i] += count
					delete(universes, ps)
					break
				}
			}
		}
		if len(universes) == 0 {
			break
		}
		playerIndex = (playerIndex + 1) % len(players)
	}

	fmt.Println("WINS:", wins)
	result := fmt.Sprint(max(int(wins[0]), int(wins[1])))

	return result
}

type dice struct {
	count int
	value int
}

func (d *dice) roll() int {
	d.count++
	d.value += 1
	if d.value > 100 {
		d.value = 1
	}
	return d.value
}

type player struct {
	pos   int
	score int
}

func (p *player) move(m int) {
	p.pos = (p.pos + m) % 10
	p.score += p.pos + 1
}

func (p player) String() string {
	return fmt.Sprintf("[pos: %d, score: %d]", p.pos, p.score)
}

func copyUniversesForPlayer(universes map[[2]player]int, playerIndex int) map[[2]player]int {
	result := map[[2]player]int{}

	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				move := i + j + k
				for players, count := range universes {
					players[playerIndex].move(move)
					result[players] += count
				}
			}
		}
	}

	return result
}
