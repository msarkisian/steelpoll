package fptp

import (
	"errors"
)

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
	voteCounts := make(map[C]int, len(p.Candidates))
	for _, c := range p.Votes {
		voteCounts[c] += 1
	}

	max := 0
	winners := []C{}

	for c, s := range voteCounts {
		switch {
		case s > max:
			max = s
			winners = []C{c}
		case s == max:
			winners = append(winners, c)
		}
	}

	ret := make(map[C]float64, len(voteCounts))
	for c := range voteCounts {
		ret[c] = 0
	}
	for _, w := range winners {
		ret[w] = 100 / float64((len(winners)))
	}
	return ret
}
