package myer

import (
	"errors"
	"fmt"
)

type Instruction struct {
	Data     string
	Decision rune
}

func DisplayInstructions(instructions []Instruction) {
	for i := range instructions {
		fmt.Printf("%c %s\n", instructions[i].Decision, instructions[i].Data)
	}
}

type V struct {
	Data    []int
	start   int
	end     int
	history [][]Instruction
}

func Initialize(start int, end int) V {
	return V{
		Data:    make([]int, end-start+1),
		start:   start,
		end:     end,
		history: make([][]Instruction, end-start+1),
	}
}

func (v *V) Set(i int, value int) {
	v.Data[i-v.start] = value
}

func (v V) Get(i int) int {
	return v.Data[i-v.start]
}

func (v V) GetB(i int) (int, []Instruction) {
	return v.Data[i-v.start], v.history[i-v.start]
}

func (v *V) SetB(i int, value int, history []Instruction) {
	v.Data[i-v.start] = value
	v.history[i-v.start] = history
}

func MyersDiff(a []string, b []string) ([]Instruction, error) {
	M := len(a)
	N := len(b)
	MAX := M + N

	V := Initialize(-MAX, MAX)
	var history []Instruction
	V.Set(1, 0)
	var x, y, old_x int
	var step_down bool
	for D := 0; D <= MAX; D++ {
		for k := -D; k <= D; k += 2 {

			step_down = (k == -D || (k != D && V.Get(k-1) < V.Get(k+1)))
			if step_down {
				old_x, history = V.GetB(k + 1)
				x = old_x
			} else {
				old_x, history = V.GetB(k - 1)
				x = old_x + 1
			}
			y = x - k

			if (1 <= y && y <= N) && step_down {
				history = append(history, Instruction{
					Data:     b[y-1],
					Decision: '+',
				})
			} else if 1 <= x && x <= M {
				history = append(history, Instruction{
					Data:     a[x-1],
					Decision: '-',
				})
			}

			for x < M && y < N && a[x] == b[y] {
				x += 1
				y += 1
				history = append(history, Instruction{
					Data:     a[x-1],
					Decision: '=',
				})
			}
			V.Set(k, x)
			if x >= M && y >= N {
				return history, nil
			} else {
				V.SetB(k, x, history)
			}
		}
	}
	return make([]Instruction, 0), errors.New("Failed to find edit script")
}
