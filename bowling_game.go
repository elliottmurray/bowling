package bowling

import "fmt"

func main() {
	fmt.Printf("Hello world")
}

type BowlingGame struct {
  Score int
}

func (b BowlingGame) roll(numberOfPins int) int {
  fmt.Printf("Number of pins %d", numberOfPins)

  b.Score = b.Score + numberOfPins
  return b.Score
  //b.Score := b.Score + numberOfPins
}
