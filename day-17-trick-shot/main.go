package main

import (
	"fmt"
	"strings"
)

func main() {
	file := openFileFromArgs()
	defer file.Close()

	var t target
	scanner := func(line string) error {
		if strings.Index(line, "target area: ") != 0 {
			return fmt.Errorf("invalid input1: %s", line)
		}
		ss := strings.Split(line[len("target area: "):], ", ")
		if len(ss) != 2 {
			return fmt.Errorf("invalid input2: %s", line)
		}
		if strings.Index(ss[0], "x=") != 0 || strings.Index(ss[1], "y=") != 0 {
			return fmt.Errorf("invalid input3: %s", line)
		}
		xs := strings.Split(ss[0][len("x="):], "..")
		if len(xs) != 2 {
			return fmt.Errorf("invalid x coordinates: %s", ss[0])
		}
		ys := strings.Split(ss[1][len("y="):], "..")
		if len(ys) != 2 {
			return fmt.Errorf("invalid y coordinates: %s", ss[1])
		}
		xx := []int{parseInt(xs[0]), parseInt(xs[1])}
		yy := []int{parseInt(ys[0]), parseInt(ys[1])}
		t.ul = coords{
			x: min(xx[0], xx[1]),
			y: max(yy[0], yy[1]),
		}
		t.lr = coords{
			x: max(xx[0], xx[1]),
			y: min(yy[0], yy[1]),
		}
		return nil
	}
	scanLines(file, scanner)

	fmt.Println("INPUT:")
	fmt.Println(t)

	result1 := puzzle1(t)
	result2 := puzzle2(t)

	fmt.Println("RESULT1:", result1)
	fmt.Println("RESULT2:", result2)
}

func puzzle1(t target) string {
	// 1) The 'x' coordinate does not matter
	// 2) Assuming that 'y' < 0:
	// 2.1) The 'y' coordinate moves up and down symmetrically, i.e. at some moment shot[n].y == shot[0].y
	// 2.2) The max 'y' is the highest if the fall after step 'n' is the largest possible to hit target
	// 2.3) The fall after step 'n' is the largest possible if it is 'target.lr.y'
	// 2.4) The velocity after step 'n' is 1 greater than initial velocity
	sum := 0
	for i := 0; i < abs(t.lr.y); i++ {
		sum += i
	}
	result := fmt.Sprintf("%d", sum)

	return result
}

func puzzle2(t target) string {
	// 1) Find the range for x velocity:
	// 1.1) Min is the lower number that would reach t.x range
	// 1.2) Max is t.lr.x to include single step shots
	// 2) Find the range for y velocity:
	// 2.1) Looking at the test data solution, assuming that it's between t.lr.y (negative) and abs(t.lr.y)-1
	minVelX := findMinVelX(t)
	maxVelX := t.lr.x
	minVelY := t.lr.y
	maxVelY := abs(t.lr.y) - 1
	fmt.Printf("vel.x: %d..%d\n", minVelX, maxVelX)
	fmt.Printf("vel.y: %d..%d\n", minVelY, maxVelY)

	count := 0
	for x := minVelX; x <= maxVelX; x++ {
		for y := minVelY; y <= maxVelY; y++ {
			s := shot{vel: coords{x: x, y: y}}
			for {
				fmt.Println("  ", x, y, s)
				if t.hit(s.c) {
					count++
					break
				}
				if s.missed(t) {
					break
				}
				s.next()
			}
		}
	}
	result := fmt.Sprintf("%d", count)

	return result
}

func findMinVelX(t target) int {
	for x := 1; ; x++ {
		// Simulate a horizontal shot
		c := coords{x: 0, y: t.ul.y}
		velX := x
		for {
			fmt.Println("  ", x, velX, c)
			if t.hit(c) {
				return x
			}
			if t.missed(c) {
				break
			}
			c.x += velX
			if velX > 0 {
				velX--
			} else if velX < 0 {
				velX++
			} else {
				break
			}
		}
	}
}

type coords struct {
	x int
	y int
}

type shot struct {
	// c is the current coordinates of the bullet
	c coords
	// vel is the velocity, represented as coordinates relative to (0, 0)
	vel coords
}

func (s *shot) next() coords {
	s.c.x += s.vel.x
	s.c.y += s.vel.y

	if s.vel.x > 0 {
		s.vel.x--
	} else if s.vel.x < 0 {
		s.vel.x++
	}
	s.vel.y--

	return s.c
}

func (s shot) missed(t target) bool {
	if t.missed(s.c) {
		return true
	}
	if s.c.x < t.ul.x && s.vel.x == 0 {
		return true
	}
	if s.vel.y < 0 && s.c.y < t.lr.y {
		return true
	}
	return false
}

type target struct {
	// ul is upper left corner
	ul coords
	// lr is lower right corner
	lr coords
}

// hit returns whether the coordinates are inside the target
func (t target) hit(c coords) bool {
	return c.x >= t.ul.x && c.y <= t.ul.y &&
		c.x <= t.lr.x && c.y >= t.lr.y
}

// missed returns whether the coordinates are passet the target, i.e will not hit in future steps
func (t target) missed(c coords) bool {
	return c.x > t.lr.x || c.y < t.lr.y
}
