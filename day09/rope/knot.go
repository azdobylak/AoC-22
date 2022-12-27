package rope

type Knot struct {
	Name string
	X    int
	Y    int
}

type Direction string

const (
	UP    Direction = "U"
	DOWN  Direction = "D"
	RIGHT Direction = "R"
	LEFT  Direction = "L"
)

func (this *Knot) Move(d Direction) {
	switch d {
	case UP:
		this.Y += 1
	case DOWN:
		this.Y -= 1
	case RIGHT:
		this.X += 1
	case LEFT:
		this.X -= 1
	}
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func orientation(head, tail int) int {
	if head > tail {
		return 1
	} else if head < tail {
		return -1
	} else {
		panic("Head an tail are the same")
	}
}

func (this *Knot) Touches(other Knot) bool {
	return Abs(this.X, other.X) <= 1 && Abs(this.Y, other.Y) <= 1
}

func (this *Knot) Follow(h Knot) {
	if this.Touches(h) {
		return
	}

	if this.X == h.X {
		this.Y += orientation(h.Y, this.Y)
	} else {
		this.X += orientation(h.X, this.X)
		if this.Y != h.Y {
			this.Y += orientation(h.Y, this.Y)
		}
	}
}

var charToDirection = map[string]Direction{
	"U": UP,
	"D": DOWN,
	"R": RIGHT,
	"L": LEFT,
}

func NewKnot(name string) Knot {
	return Knot{X: 0, Y: 0, Name: name}
}

func NewDirection(char string) Direction {
	return charToDirection[char]
}
