package stv

import (
	"fmt"
	"math"
)

type poll[C comparable, V comparable] struct {
	candidates map[C]struct{}
	votes      map[V]Vote[C]
	winners    int
}

type Vote[C comparable] []C

func New[C comparable, V comparable](candidates []C, winners int) poll[C, V] {
	poll := poll[C, V]{
		candidates: make(map[C]struct{}),
		votes:      make(map[V]Vote[C]),
		winners:    winners,
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
	quota := (len(p.votes) / (p.winners + 1)) + 1
	winners := make([]C, 0, p.winners)
	voteCounts := make(map[C]float64, len(p.candidates))
	eliminatedCandidates := make(map[C]struct{}, len(p.candidates))

Round:
	for len(winners) < p.winners {
		for _, ballot := range p.votes {
			voteWeight := 1.0
			for _, c := range ballot {
				if _, ok := eliminatedCandidates[c]; !ok {
					voteCounts[c] += voteWeight
					break
				}
				voteWeight /= float64(quota)
				// this is wrong
				// this is only true if there's another choice for all voters of the eliminated
			}
		}
		// search for winners (above quota)
		for c, voteCount := range voteCounts {
			if voteCount > float64(quota) {
				winners = append(winners, c)
				eliminatedCandidates[c] = struct{}{}
				continue Round
			}
		}
		// search for losers
		minVoteCount := math.MaxInt
		var minC C

		for c, voteCount := range voteCounts {
			if voteCount < float64(minVoteCount) {
				minVoteCount = int(voteCount)
				minC = c
			}
		}
		eliminatedCandidates[minC] = struct{}{}
	}
	ret := make(map[C]float64)
	for c := range p.candidates {
		ret[c] = 0
	}
	for _, c := range winners {
		ret[c] = 100 / float64(p.winners)
	}
	return ret
}
