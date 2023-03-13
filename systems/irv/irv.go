package irv

import (
	"fmt"
)

type poll[C comparable, V comparable] struct {
	candidates map[C]struct{}
	votes      map[V]Vote[C]
}

type Vote[C comparable] []C

func New[C comparable, V comparable](candidates []C) poll[C, V] {
	poll := poll[C, V]{
		candidates: make(map[C]struct{}),
		votes:      make(map[V]Vote[C]),
	}
	for _, c := range candidates {
		poll.candidates[c] = struct{}{}
	}
	return poll
}

func (p poll[C, V]) CastVote(vote Vote[C], voter V) error {
	for _, c := range vote {
		_, ok := p.candidates[c]
		if !ok {
			return fmt.Errorf("tried to vote for nonexistant candidate: %v", c)
		}
	}
	p.votes[voter] = vote
	return nil
}
