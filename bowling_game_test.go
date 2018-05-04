package bowling

import "testing"

func TestBowlingGame(t *testing.T) {
  noOfPins := 3
  var bowling BowlingGame
  bowling = BowlingGame{}
  var score = bowling.roll(noOfPins)
  if (score != 3) {
    t.Fail()
  }
  //BowlingGame.score()
}

