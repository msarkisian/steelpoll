package irv

import (
	"fmt"
	"math"
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
		if _, ok := p.candidates[c]; !ok {
			return fmt.Errorf("tried to vote for nonexistant candidate: %v", c)
		}
	}
	p.votes[voter] = vote
	return nil
}

func (p poll[C, V]) Tally() map[C]float64 {
	eliminatedCandidates := make(map[C]struct{}, len(p.candidates))
	for {
		voteCounts := make(map[C]int, len(p.candidates))
		threshold := 0
		for _, v := range p.votes {
			if len(v) > 0 {
				for _, c := range v {
					if _, ok := eliminatedCandidates[c]; !ok {
						threshold += 1
						voteCounts[c] += 1
						break
					}
				}
			}
		}
		min := math.MaxInt
		for c, voteCount := range voteCounts {
			if float64(voteCount) > float64(threshold)/2.0 {
				// found a winner
				ret := make(map[C]float64, len(p.candidates))
				for c := range p.candidates {
					ret[c] = 0
				}
				ret[c] = 100
				return ret
			}
			if voteCount < min {
				min = voteCount
			}
		}
		for c := range p.candidates {
			if voteCounts[c] == min {
				eliminatedCandidates[c] = struct{}{}
			}
		}
	}
}
