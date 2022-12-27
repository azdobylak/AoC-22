package main

import (
	"aoc22/day09/rope"
	"testing"
)

func assertPosition(t *testing.T, k rope.Knot, x_exp, y_exp int) {
	if k.X != x_exp {
		t.Errorf("Knot(%s) x position is incorrect: %d != %d", k.Name, k.X, x_exp)
	}

	if k.Y != y_exp {
		t.Errorf("Knot(%s) y position is incorrect: %d != %d", k.Name, k.Y, y_exp)
	}
}

func TestKnotMovements(t *testing.T) {
	head := rope.NewKnot("H")

	head.Move(rope.UP)
	assertPosition(t, head, 0, 1)

	head.Move(rope.UP)
	assertPosition(t, head, 0, 2)

	head.Move(rope.RIGHT)
	assertPosition(t, head, 1, 2)

	head.Move(rope.DOWN)
	assertPosition(t, head, 1, 1)

	head.Move(rope.LEFT)
	assertPosition(t, head, 0, 1)

	head.Move(rope.RIGHT)
	head.Move(rope.LEFT)
	assertPosition(t, head, 0, 1)
}

func TestTailKnotMovements01(t *testing.T) {
	head := rope.NewKnot("H")
	tail := rope.NewKnot("T")

	// R 4
	assertPosition(t, head, 0, 0)
	assertPosition(t, tail, 0, 0)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 1, 0)
	assertPosition(t, tail, 0, 0)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 2, 0)
	assertPosition(t, tail, 1, 0)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 3, 0)
	assertPosition(t, tail, 2, 0)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 4, 0)
	assertPosition(t, tail, 3, 0)

	// U 4
	move(&head, &tail, rope.UP)
	assertPosition(t, head, 4, 1)
	assertPosition(t, tail, 3, 0)

	move(&head, &tail, rope.UP)
	assertPosition(t, head, 4, 2)
	assertPosition(t, tail, 4, 1)

	move(&head, &tail, rope.UP)
	assertPosition(t, head, 4, 3)
	assertPosition(t, tail, 4, 2)

	move(&head, &tail, rope.UP)
	assertPosition(t, head, 4, 4)
	assertPosition(t, tail, 4, 3)

	// L 3
	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 3, 4)
	assertPosition(t, tail, 4, 3)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 2, 4)
	assertPosition(t, tail, 3, 4)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 1, 4)
	assertPosition(t, tail, 2, 4)

	// D 1
	move(&head, &tail, rope.DOWN)
	assertPosition(t, head, 1, 3)
	assertPosition(t, tail, 2, 4)

	// R 4
	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 2, 3)
	assertPosition(t, tail, 2, 4)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 3, 3)
	assertPosition(t, tail, 2, 4)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 4, 3)
	assertPosition(t, tail, 3, 3)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 5, 3)
	assertPosition(t, tail, 4, 3)

	// D 1
	move(&head, &tail, rope.DOWN)
	assertPosition(t, head, 5, 2)
	assertPosition(t, tail, 4, 3)

	// L 5
	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 4, 2)
	assertPosition(t, tail, 4, 3)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 3, 2)
	assertPosition(t, tail, 4, 3)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 2, 2)
	assertPosition(t, tail, 3, 2)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 1, 2)
	assertPosition(t, tail, 2, 2)

	move(&head, &tail, rope.LEFT)
	assertPosition(t, head, 0, 2)
	assertPosition(t, tail, 1, 2)

	// R2
	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 1, 2)
	assertPosition(t, tail, 1, 2)

	move(&head, &tail, rope.RIGHT)
	assertPosition(t, head, 2, 2)
	assertPosition(t, tail, 1, 2)
}
