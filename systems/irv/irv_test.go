package irv_test

import (
	"testing"

	"github.com/msarkisian/steelpoll/systems/irv"
)

func TestIRVPoll(t *testing.T) {
	poll := irv.New[string, string]([]string{"A", "B", "C"})
	poll.CastVote(irv.Vote[string]{"C", "A", "B"}, "Bob")
	poll.CastVote(irv.Vote[string]{"A", "B"}, "Joe")
	poll.CastVote(irv.Vote[string]{"B", "A"}, "Chris")
	poll.CastVote(irv.Vote[string]{"A"}, "Will")
	poll.CastVote(irv.Vote[string]{"B"}, "Richard")

	res := poll.Tally()
	if res["A"] != 100 {
		t.Errorf("winner of irv poll should have score 100, got %f", res["A"])
	}
	if res["B"] != 0 {
		t.Errorf("loser of irv poll should have score 0, got %f", res["B"])
	}
	if res["C"] != 0 {
		t.Errorf("loser of irv poll should have score 0, got %f", res["C"])
	}
}
