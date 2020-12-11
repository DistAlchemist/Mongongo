// Copyright (c) 2020 DistAlchemist
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package utils

import (
	"math"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

// BoundedStatsDeque ...
type BoundedStatsDeque struct {
	size int
	d    *deque.Deque
}

// NewBoundedStatsDeque ...
func NewBoundedStatsDeque(size int) *BoundedStatsDeque {
	b := &BoundedStatsDeque{}
	b.size = size
	b.d = deque.New()
	return b
}

// Size ...
func (p *BoundedStatsDeque) Size() int {
	return p.d.Size()
}

// Clear ...
func (p *BoundedStatsDeque) Clear() {
	p.d.Reset()
}

// Add ...
func (p *BoundedStatsDeque) Add(o float64) {
	if p.size == p.d.Size() {
		p.d.PopLeft()
	}
	p.d.PushRight(o)
}

// Sum ...
func (p *BoundedStatsDeque) Sum() float64 {
	sum := float64(0)
	r := deque.New()
	for p.d.Empty() == false {
		interval := p.d.PopLeft()
		sum += interval.(float64)
		r.PushRight(interval)
	}
	p.d = r
	return sum
}

// SumOfDeviations ...
func (p *BoundedStatsDeque) SumOfDeviations() float64 {
	res := float64(0)
	mean := p.Mean()
	r := deque.New()
	for p.d.Empty() == false {
		interval := p.d.PopLeft()
		v := interval.(float64) - mean
		res += v * v
		r.PushRight(interval)
	}
	p.d = r
	return res
}

// Mean ...
func (p *BoundedStatsDeque) Mean() float64 {
	return p.Sum() / float64(p.Size())
}

// Variance ...
func (p *BoundedStatsDeque) Variance() float64 {
	return p.SumOfDeviations() / float64(p.Size())
}

// Stdev ...
func (p *BoundedStatsDeque) Stdev() float64 {
	return math.Sqrt(p.Variance())
}
