package fptp

import (
	"fmt"
)

type poll[C comparable, V comparable] struct {
	candidates map[C]struct{}
	votes      map[V]C
}

func New[C comparable, V comparable](candidates []C) poll[C, V] {
	poll := poll[C, V]{
		candidates: make(map[C]struct{}),
		votes:      make(map[V]C),
	}
	for _, c := range candidates {
		poll.candidates[c] = struct{}{}
	}
	return poll
}

func (p poll[C, V]) CastVote(candidate C, voter V) error {
	if _, ok := p.candidates[candidate]; !ok {
		return fmt.Errorf("tried to vote for nonexistant candidate: %v", candidate)
	}
	p.votes[voter] = candidate
	return nil
}

func (p poll[C, V]) Tally() map[C]float64 {
	voteCounts := make(map[C]int, len(p.candidates))
	for _, c := range p.votes {
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
