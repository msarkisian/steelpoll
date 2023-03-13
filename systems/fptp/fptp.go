package fptp

import (
	"errors"
)

// TODO:
// gracefully handle/deny changing votes
// handle ties

type Poll[C comparable, V comparable] struct {
	Candidates map[C]struct{}
	Votes      map[V]C
}

func New[C comparable, V comparable](candidates []C) Poll[C, V] {
	poll := Poll[C, V]{
		Candidates: make(map[C]struct{}),
		Votes:      make(map[V]C),
	}
	for _, c := range candidates {
		poll.Candidates[c] = struct{}{}
	}
	return poll
}

func (p Poll[C, V]) CastVote(candidate C, voter V) error {
	_, ok := p.Candidates[candidate]
	if !ok {
		return errors.New("tried to vote for nonexistant candidate")
	}
	p.Votes[voter] = candidate
	return nil
}

func (p Poll[C, V]) Tally() map[C]float64 {
	maxVoteCount := 0
	var leader *C = nil
	voteCounts := make(map[C]int, len(p.Candidates))
	for _, c := range p.Votes {
		voteCounts[c] += 1
		if voteCounts[c] > maxVoteCount {
			maxVoteCount = voteCounts[c]
			leader = &c
		}
	}
	ret := make(map[C]float64, len(voteCounts))
	for c := range voteCounts {
		if c == *leader {
			ret[c] = 100
		} else {
			ret[c] = 0
		}
	}
	return ret
}
