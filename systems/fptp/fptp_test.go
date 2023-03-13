package fptp_test

import (
	"testing"

	"github.com/msarkisian/steelpoll/systems/fptp"
)

func TestFTPTPoll(t *testing.T) {
	poll := fptp.New[string, string]([]string{"A", "B", "C"})
	poll.CastVote("A", "Bob")
	poll.CastVote("B", "Joe")
	poll.CastVote("B", "Chris")
	poll.CastVote("A", "Joe") // Change vote

	t.Log(poll)

	res := poll.Tally()
	if res["A"] != 100 {
		t.Errorf("winner of fptp poll should have score 100, got %f", res["A"])
	}
	if res["B"] != 0 {
		t.Errorf("loser of fptp poll should have score 0, got %f", res["B"])
	}
	if res["C"] != 0 {
		t.Errorf("loser of fptp poll should have score 0, got %f", res["C"])
	}
}

func TestFPTPTie(t *testing.T) {
	poll := fptp.New[string, string]([]string{"A", "B", "C"})
	poll.CastVote("A", "Bob")
	poll.CastVote("B", "Joe")

	res := poll.Tally()
	if res["A"] != 50 {
		t.Errorf("tied winner of fptp poll should have score 50, got %f", res["A"])
	}
	if res["B"] != 50 {
		t.Errorf("tied winner of fptp poll should have score 50, got %f", res["B"])
	}
	if res["C"] != 0 {
		t.Errorf("loser of fptp poll should have score 0, got %f", res["C"])
	}

}
